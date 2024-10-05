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
	expirationTime := time.Now().UTC().Add(24 * time.Hour)

	return generateToken(user, expirationTime, []byte(GetJWTSecret()))
}

func generateRefreshToken(user *user.User) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().UTC().Add(24 * time.Hour)

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

	c.SetCookie(cookie)
}
func removeTokenCookie(name string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func setUserCookie(user *user.User, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = user.Email
	cookie.Expires = expiration
	cookie.Path = "/"

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
func JWTErrorChecker(c echo.Context, err error) error {
	slog.Error("JWTErrorChecker", "error", err)

	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("SignInForm"))
}

// TokenRefresherMiddleware middleware, which refreshes JWT tokens if the access token is about to expire.
func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// If the user is not authenticated (no user token data in the context), don't do anything.
		if c.Get("user") == nil {
			return next(c)
		}
		// Gets user token from the context.
		u := c.Get("user").(*jwt.Token)

		claims := u.Claims.(*JwtCustomClaims)

		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 15 mins of expiry.
		if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now().UTC()) < 15*time.Minute {
			// Gets the refresh token from the cookie.
			rc, err := c.Cookie(refreshTokenCookieName)
			if err == nil && rc != nil {
				// Parses token and checks if it valid.
				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(GetRefreshJWTSecret()), nil
				})
				if err != nil {
					if err == jwt.ErrSignatureInvalid {

						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					// If everything is good, update tokens.
					_ = GenerateTokensAndSetCookies(&user.User{
						Name: claims.Name,
					}, c)
				}
			}
		}

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
