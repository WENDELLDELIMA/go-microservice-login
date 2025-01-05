package routes

import (
	"github.com/WENDELLDELIMA/go-microservice-login/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, authHandler *handler.AuthHandler) {
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
}
