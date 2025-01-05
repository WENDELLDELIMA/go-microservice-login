package handler

import (
	"net/http"

	"github.com/WENDELLDELIMA/go-microservice-login/internal/db" // Add this import statement
	"github.com/WENDELLDELIMA/go-microservice-login/internal/service"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service *service.AuthService
	Queries *db.Queries
}

func NewAuthHandler(service *service.AuthService, queries *db.Queries) *AuthHandler {
	return &AuthHandler{Service: service, Queries: queries}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	hashedPassword, err := h.Service.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar senha"})
	}

	user, err := h.Queries.CreateUser(c.Request().Context(), db.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	user, err := h.Queries.GetUserByUsername(c.Request().Context(), req.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuário não encontrado"})
	}

	err = h.Service.VerifyPassword(user.PasswordHash, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Senha inválida"})
	}

	token, err := h.Service.GenerateToken(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao gerar token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
