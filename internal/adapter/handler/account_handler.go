package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	f "github.com/stelgkio/otoo/internal/adapter/web/view/account/forgot_password"
	l "github.com/stelgkio/otoo/internal/adapter/web/view/account/login"
	reg "github.com/stelgkio/otoo/internal/adapter/web/view/account/register"
	re "github.com/stelgkio/otoo/internal/adapter/web/view/account/reset_password"
	us "github.com/stelgkio/otoo/internal/adapter/web/view/account/user"
	tm "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/team"
	st "github.com/stelgkio/otoo/internal/adapter/web/view/project/settings/template"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
	r "github.com/stelgkio/otoo/internal/core/util"
)

// AuthHandler represents the HTTP handler for authentication-related requests
type AuthHandler struct {
	svc          port.AuthService
	urs          port.UserService
	projectSvc   port.ProjectService
	extensionSvc port.ExtensionService
	protouserSvc port.UserProjectService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(
	svc port.AuthService,
	urs port.UserService,
	projectSvc port.ProjectService,
	extensionSvc port.ExtensionService,
	protouserSvc port.UserProjectService) *AuthHandler {
	return &AuthHandler{
		svc,
		urs,
		projectSvc,
		extensionSvc,
		protouserSvc,
	}
}

// authResponse represents an authentication response body
type authResponse struct {
	AccessToken string `json:"token"`
}

// AuthResponse creates an authentication response
func AuthResponse(token string) authResponse {
	return authResponse{
		AccessToken: token,
	}
}

// loginRequest represents the request body for logging in a user
type loginRequest struct {
	Email    string `form:"email" validate:"required,email" example:"test@example.com"`
	Password string `form:"password" validate:"required,min=8" example:"12345678" minLength:"8"`
}

// loginRequest represents the request body for logging in a user
type forgotPosswordRequest struct {
	Email string `form:"email" validate:"required,email" example:"test@example.com"`
}

// loginRequest represents the request body for logging in a user
type resetPosswordRequest struct {
	Password             string `form:"password" validate:"required,min=8"`
	ConfirmationPassword string `form:"confirmationpassword" validate:"required,min=8"`
}

// registerRequest represents the request body for creating a user
type registerRequest struct {
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"required,min=8"`
	ConfirmationPassword string `form:"confirmationpassword" validate:"required,min=8"`
	Name                 string `form:"name" validate:"required"`
	LastName             string `form:"last_name" validate:"required"`
}

type addmemberRequest struct {
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"required,min=8"`
	ConfirmationPassword string `form:"confirmationpassword" validate:"required,min=8"`
	Name                 string `form:"name" validate:"required"`
	LastName             string `form:"last_name" validate:"required"`
	ReceiveNotification  bool   `form:"receive_notification"`
	UserExist            bool   `form:"user_exist"`
	Role                 string `form:"role"`
}

// Validate validates the request body
func (p *addmemberRequest) Validate() map[string](string) {

	errors := make(map[string]string)

	if p.Email == "" {
		errors["email"] = "Email is required"
	}
	if p.UserExist == false {
		if p.Name == "" {
			errors["name"] = "Name is required"
		}

		if p.LastName == "" {
			errors["lastname"] = "LastName is required"
		}
		if p.Password == "" {
			errors["password"] = "Password key is required"
		}
		if len(p.Password) < 8 {
			errors["password"] = "Password must be at least 8 characters long"
		}
		if p.Role == "" {
			errors["role"] = "Role is required"
		}

		if p.ConfirmationPassword == "" {
			errors["confirmation_password"] = "Confirmation Password is required"
		}
		// Check if Password and Confirmation Password match
		if p.Password != p.ConfirmationPassword {
			errors["confirmation_password"] = "Passwords do not match"
		}
	}

	return errors
}

// Login @Router			/login [post]
func (ah *AuthHandler) Login(ctx echo.Context) (err error) {

	req := new(loginRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	_, err = ah.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return r.Render(ctx, l.Login(err))
	}

	//AuthResponse(token)

	return ctx.Redirect(http.StatusFound, "/dashboard")
}

// LoginForm GET /login
func (ah *AuthHandler) LoginForm(c echo.Context) error {
	c.Response().Header().Set("HX-Redirect", "/login")
	return r.Render(c, l.Login(nil))

}

// Logout GET /logout
func (ah *AuthHandler) Logout(ctx echo.Context) (err error) {
	err = ah.svc.Logout(ctx)
	if err != nil {
		ctx.Response().Header().Set("HX-Redirect", "/")
		return ctx.Redirect(http.StatusAccepted, "/index")
	}
	ctx.Response().Header().Set("HX-Redirect", "/")
	return ctx.Redirect(http.StatusAccepted, "/index")
}

// RegisterForm @Router			/register [get]
func (ah *AuthHandler) RegisterForm(ctx echo.Context) error {
	return r.Render(ctx, reg.Register(0, nil, nil))
}

// Register @Router			/register [post]
func (ah *AuthHandler) Register(ctx echo.Context) error {

	req := new(registerRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	ctx.Validate(req)
	// validate email is not taken
	validate := validator.New()
	// validate password is the same as confirm password
	if req.Password != req.ConfirmationPassword {
		return r.Render(ctx, reg.Register(0, nil, fmt.Errorf("invalid confirmation password")))
	}
	// Validate the User struct
	err := validate.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		return r.Render(ctx, reg.Register(0, errors, nil))
	}
	chechUser, err := ah.urs.GetUserByEmail(ctx, req.Email)
	if err != nil {
		slog.Error("error get user by email:", "StatusBadRequest", err)
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, fmt.Errorf("Somethin when wrong try again later")))
	}
	if chechUser != nil {
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
	}

	user, err := domain.NewUser(req.Email, req.Password, req.Name, req.LastName)
	if err != nil {
		slog.Error("error new user:", "StatusBadRequest", err)
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
	}

	_, err = ah.urs.CreateUser(ctx, user)
	if err != nil {
		slog.Error("error create new user:", "StatusBadRequest", err)
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
	}

	_, err = ah.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		slog.Error("error new user create token:", "StatusInternalServerError", err)
		return r.Render(ctx, reg.Register(http.StatusInternalServerError, nil, nil))
	}
	return ctx.Redirect(http.StatusMovedPermanently, ctx.Echo().Reverse("dashboard"))

}

// ForgotPasswordForm @Router			/ForgotPassword [get]
func (ah *AuthHandler) ForgotPasswordForm(c echo.Context) error {
	return r.Render(c, f.ForgotPassword())

}

// ForgotPassword @Router			/ForgotPassword [post]
func (ah *AuthHandler) ForgotPassword(ctx echo.Context) error {
	req := new(forgotPosswordRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	go ah.svc.ForgotPassword(ctx, req.Email)

	return r.Render(ctx, f.ForgotPasswordSuccess())
}

// ResetPasswordForm @Router			/ForgotPassword [get]
func (ah *AuthHandler) ResetPasswordForm(ctx echo.Context) error {
	token := ctx.Param("token")
	email, _ := r.Decrypt(token)
	user, err := ah.urs.GetUserByEmail(ctx, email)

	if user == nil {
		slog.Error("error get existing user:", "StatusBadRequest", err)
		return r.Render(ctx, re.ResetPasswordError())
	}
	return r.Render(ctx, re.ResetPasswordForm(0, email, nil, nil))

}

// ResetPassword @Router			/ForgotPassword [post]
func (ah *AuthHandler) ResetPassword(ctx echo.Context) error {
	req := new(resetPosswordRequest)
	email := ctx.Param("email")
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	ctx.Validate(req)
	// validate email is not taken
	validate := validator.New()
	// validate password is the same as confirm password
	if req.Password != req.ConfirmationPassword {
		return r.Render(ctx, re.ResetPasswordForm(0, "", nil, fmt.Errorf("invalid confirmation password")))
	}
	// Validate the User struct
	err := validate.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		return r.Render(ctx, re.ResetPasswordForm(0, "", errors, nil))
	}

	fromExistingUser, err := ah.urs.GetUserByEmail(ctx, email)
	if err != nil {
		slog.Error("error get existing user:", "StatusBadRequest", err)
		return r.Render(ctx, re.ResetPasswordForm(http.StatusBadRequest, "", nil, nil))
	}
	user, err := domain.NewUser(fromExistingUser.Email, req.Password, fromExistingUser.Name, fromExistingUser.LastName)
	if err != nil {
		slog.Error("error new user:", "StatusBadRequest", err)
		return r.Render(ctx, re.ResetPasswordForm(http.StatusBadRequest, "", nil, nil))
	}

	_, err = ah.urs.UpdateUser(ctx, user)
	if err != nil {
		slog.Error("error update user:", "StatusBadRequest", err)
		return r.Render(ctx, re.ResetPasswordForm(http.StatusBadRequest, "", nil, nil))
	}

	return ctx.Redirect(http.StatusMovedPermanently, "/login")
}

func (ah *AuthHandler) UserList(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	id, err := uuid.Parse(projectID)
	if err != nil {
		fmt.Println("Invalid UUID format:", err)
		return err
	}
	users, err := ah.urs.FindUsersByProjectId(ctx, id)

	return r.Render(ctx, us.UserList(users, projectID))
}

func (ah *AuthHandler) CheckEmail(ctx echo.Context) error {
	req := new(addmemberRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("Oups something when wrong").Error(),
		})

	}
	users, err := ah.urs.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("Oups something when wrong").Error(),
		})
	}
	if users != nil {
		return r.Render(ctx, us.UserExist(true))
	}
	return r.Render(ctx, us.UserExist(false))
}

func (ah *AuthHandler) CreateMemberModal(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	return r.Render(ctx, us.CreateMeember(projectID, nil))
}

// AddMember @Router			/projects/{projectId}/members [post]
func (ah *AuthHandler) AddMember(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	req := new(addmemberRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("Oups something when wrong").Error(),
		})

	}

	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": util.ConcatenateErrors(validationErrors),
		})

	}

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}

	user, err := ah.urs.GetUserById(ctx, userID)
	if err != nil {
		return err
	}
	project, err := ah.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := ah.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}

	newUser, err := domain.NewClientUser(req.Email, req.Password, req.Name, req.LastName, req.Role, req.ReceiveNotification)
	if err != nil {
		slog.Error("error new user:", "StatusBadRequest", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "error creating new user",
		})
	}

	userExist, err := ah.urs.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if userExist == nil {
		usr, err := ah.urs.CreateUser(ctx, newUser)
		err = ah.protouserSvc.AddUserToProject(ctx, usr.Id, project.Id)
		if err != nil {
			slog.Error("error create new user:", "StatusBadRequest", err)
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": errors.New("error creating new user").Error(),
			})
		}
	} else {
		users, _ := ah.urs.FindUsersByProjectId(ctx, project.Id)
		if !domain.ContainsUserID(users, userExist.Id) {
			err = ah.protouserSvc.AddUserToProject(ctx, userExist.Id, project.Id)
		} else {

			slog.Error("error create new user:", "StatusBadRequest", err)
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": errors.New("user already exist to this project").Error(),
			})
		}

	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, tm.Team(project, projectExtensions, user))
	}
	return r.Render(ctx, st.TeamTemplate(user, project.Name, projectID, project, projectExtensions))
}

// RemoveMember @Router
func (ah *AuthHandler) RemoveMember(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	userID := ctx.Param("userId")
	useruuID, err := uuid.Parse(userID)
	logedinUserID, err := auth.GetUserID(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("invalid user id").Error(),
		})

	}

	user, err := ah.urs.GetUserById(ctx, useruuID)
	if err != nil {
		return err
	}
	project, err := ah.projectSvc.GetProjectByID(ctx, projectID)
	if err != nil {
		return err
	}
	projectExtensions, err := ah.extensionSvc.GetAllProjectExtensions(ctx, projectID)
	if err != nil {
		return err
	}
	if user.Role == domain.Client {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("Cannot remove admin user. Contact our support for more information").Error(),
		})
	}
	err = ah.protouserSvc.RemoveUserFromProject(ctx, useruuID, project.Id)

	err = ah.urs.DeleteUser(ctx, useruuID)
	if useruuID == logedinUserID {
		ah.svc.Logout(ctx)
		ctx.Response().Header().Set("HX-Redirect", "/index")
		return ctx.Redirect(http.StatusAccepted, "/index")
	}
	if err != nil {
		slog.Error("error create new user:", "StatusBadRequest", err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": errors.New("error removing user").Error(),
		})
	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		return r.Render(ctx, tm.Team(project, projectExtensions, user))
	}
	return r.Render(ctx, st.TeamTemplate(user, project.Name, projectID, project, projectExtensions))
}
