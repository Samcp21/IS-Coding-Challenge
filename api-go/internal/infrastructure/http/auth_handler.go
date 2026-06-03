package http

import (
	"api-go/internal/application"

	"github.com/gofiber/fiber/v3"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct {
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato de credenciales inválido",
		})
	}

	token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Usuario o contraseña incorrectos",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
