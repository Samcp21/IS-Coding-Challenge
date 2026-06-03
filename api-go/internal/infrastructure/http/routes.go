package http

import (
	"api-go/internal/infrastructure/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, authHandler *AuthHandler, matrixHandler *MatrixHandler) {

	api := app.Group("/api")

	api.Post("/login", authHandler.Login)

	protected := api.Group("/", middleware.NewAuthMiddleware())
	protected.Post("/matrix", matrixHandler.ProcessMatrix)
}
