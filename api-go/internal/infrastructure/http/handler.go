package http

import (
	"api-go/internal/application"

	"github.com/gofiber/fiber/v3"
)

type MatrixRequest struct {
	Data [][]float64 `json:"data"`
}

type MatrixHandler struct {
	service *application.MatrixService
}

func NewMatrixHandler(service *application.MatrixService) *MatrixHandler {
	return &MatrixHandler{service: service}
}

func (h *MatrixHandler) ProcessMatrix(c fiber.Ctx) error {
	var req MatrixRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato de matriz inválido",
		})
	}

	result, err := h.service.ProcessMatrix(req.Data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error procesando la matriz",
		})
	}

	return c.JSON(result)
}
