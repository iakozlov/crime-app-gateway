package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github/iakozlov/crime-app-gateway/docs" // register my own docs
)

// InitCommonRoutes registers handlers and uses middlewares.
func InitCommonRoutes(e *echo.Echo) http.Handler {
	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	return e
}
