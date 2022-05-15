package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

const (
	crimeAuthURLRegister = "http://localhost:8080/users/register"
	crimeAuthURLLogin    = "http://localhost:8080/users/login"
)

type AuthHandler struct {
	log *logrus.Logger
}

func NewAuthHandler(log *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		log: log,
	}
}

func (h AuthHandler) InitRoutes(e *echo.Echo, timeout time.Duration) {
	users := e.Group(
		"users",
		middleware.RequestID(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}),
		middleware.Recover(),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Skipper:      middleware.DefaultSkipper,
			ErrorMessage: ErrTimout.Error(),
			OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
				c.Error(err)
			},
			Timeout: timeout * time.Second,
		}),
	)
	users.POST("/login", h.SignInHandler)
	users.POST("/registry", h.SignUpHandler)
}

// SignUpHandler godoc
// @Summary      provides signing up operation
// @Description  registers user in crime-app microservices ecosystem
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   body      domain.User  true  "User's consisted of login and password"
// @Success      201  {object}  domain.User
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Failure default {object} echo.HTTPError
// @Router       /users/register [post]
func (h AuthHandler) SignUpHandler(c echo.Context) error {
	return c.Redirect(http.StatusPermanentRedirect, crimeAuthURLRegister)
}

// SignInHandler godoc
// @Summary      provides signing in operation
// @Description  authorize user in crime-app microservices ecosystem
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   body      domain.User  true  "User's consisted of login and password"
// @Success      200  {object}  string
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Failure default {object} echo.HTTPError
// @Router       /users/login [post]
func (h AuthHandler) SignInHandler(c echo.Context) error {
	return c.Redirect(http.StatusPermanentRedirect, crimeAuthURLLogin)
}
