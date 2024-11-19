package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	user "github.com/stelgkio/otoo/internal/core/domain"
)

const (
	accessTokenCookieName  = "accesstoken"
	refreshTokenCookieName = "refreshtoken"
	// Just for the demo purpose, I declared secrets here. In the real-world application, you might need
	// to get it from the env variables.
	jwtSecretKey        = "some-secret-key"
	jwtRefreshSecretKey = "some-refresh-secret-key"
)

func GetJWTSecret() string {
	return jwtSecretKey
}

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func GenerateTokensAndSetCookies(user *user.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(user, exp, c)
	refreshToken, exp, err := generateRefreshToken(user)
	if err != nil {
		return err
	}
	setTokenCookie(refreshTokenCookieName, refreshToken, exp, c)

	return nil
}

func RemoveTokensAndDeleteCookies(user *user.User, c echo.Context) error {

	removeTokenCookie(accessTokenCookieName, c)
	removeUserCookie(c)
	removeTokenCookie(refreshTokenCookieName, c)

	return nil
}

func generateAccessToken(user *user.User) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().UTC().Add(24 * 30 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateRefreshToken(user *user.User) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().UTC().Add(24 * 30 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetRefreshJWTSecret()))
}

func generateToken(user *user.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &JwtCustomClaims{
		Name:  user.Email,
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 72)),
			ID:        user.Id.String(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now().UTC(), err
	}

	return tokenString, expirationTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false

	c.SetCookie(cookie)
}
func removeTokenCookie(name string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false

	c.SetCookie(cookie)
}

func setUserCookie(user *user.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Email
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false
	c.SetCookie(cookie)
}
func removeUserCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"

	c.SetCookie(cookie)
}

// JWTErrorChecker will be executed when user try to access a protected path.
// func JWTErrorChecker(c echo.Context, err error) error {
// 	slog.Error("JWTErrorChecker", "error", err)

//		return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("SignInForm"))
//	}
func JWTErrorChecker(c echo.Context, err error) error {
	if errors.Is(err, jwt.ErrTokenExpired) {
		slog.Warn("JWT token expired", "error", err)
		return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("SignInForm"))
	}
	slog.Error("JWT authentication error", "error", err)
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("SignInForm"))
}

// TokenRefresherMiddleware middleware, which refreshes JWT tokens if the access token is about to expire.
func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check if the user is authenticated
		userToken, ok := c.Get("user").(*jwt.Token)
		if !ok || userToken == nil {
			//	slog.Error("No valid user token in context")
			return next(c)
		}

		// Retrieve and verify claims from user token
		claims, ok := userToken.Claims.(*JwtCustomClaims)
		if !ok || claims == nil {
			//	slog.Error("Failed to retrieve claims from user token")
			return next(c)
		}

		// Debug: log the expiration time of the access token
		slog.Info("Access token expiration time", "expiresAt", claims.ExpiresAt.Time)

		// Check if access token is expiring soon (within 15 minutes)
		if time.Until(claims.ExpiresAt.Time) < 24*time.Hour {
			// Log that we're attempting to refresh the token
			slog.Info("Access token is expiring soon, attempting refresh")

			// Retrieve the refresh token from cookies
			refreshCookie, err := c.Cookie(refreshTokenCookieName)
			if err != nil || refreshCookie == nil {
				slog.Error("Missing refresh token cookie", "error", err)
				return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("SignInForm")) // or handle the error differently
			}

			// Parse and validate the refresh token
			refreshClaims := &JwtCustomClaims{}
			refreshToken, err := jwt.ParseWithClaims(refreshCookie.Value, refreshClaims, func(token *jwt.Token) (interface{}, error) {
				return []byte(GetRefreshJWTSecret()), nil
			})

			if err != nil || !refreshToken.Valid {
				slog.Error("Invalid or expired refresh token", "error", err)
				return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("SignInForm"))
			}

			// Log that refresh token is valid and a new access token will be generated
			slog.Info("Refresh token is valid; generating new tokens")

			// Generate new tokens and set cookies
			err = GenerateTokensAndSetCookies(&user.User{
				Name: refreshClaims.Name,
			}, c)

			if err != nil {
				slog.Error("Failed to set new tokens", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to refresh token")
			}

			slog.Info("Tokens refreshed successfully")
		}

		// Continue to the next handler if token is valid or has been refreshed
		return next(c)
	}
}

func GetUserID(c echo.Context) (uuid.UUID, error) {
	u := c.Get("user").(*jwt.Token)

	claims := u.Claims.(*JwtCustomClaims)
	uuidValue, err := uuid.Parse(claims.ID)

	if err != nil {
		return uuidValue, errors.New("error parsing uuid for user")
	}

	return uuidValue, nil
}

func GenerateWebHookAccessToken(projectId string) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().UTC().Add(5 * time.Hour)

	return generateWebHookToken(projectId, expirationTime, []byte(GetJWTSecret()))
}
func generateWebHookToken(projectId string, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &JwtCustomClaims{
		Name:  projectId,
		Admin: true,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 72)),
			ID:        projectId,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now().UTC(), err
	}

	return tokenString, expirationTime, nil
}
