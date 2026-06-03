package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestJWTMiddleware(t *testing.T) {
	app := fiber.New()

	app.Get("/protegido", NewAuthMiddleware(), func(c fiber.Ctx) error {
		return c.SendString("Acceso Concedido")
	})

	t.Run("Acceso Denegado - Sin Token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("Se esperaba status 401 Unauthorized, se obtuvo %d", resp.StatusCode)
		}
	})

	t.Run("Acceso Denegado - Token Inválido", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/protegido", nil)
		req.Header.Set("Authorization", "Bearer token_falso_inventado")
		resp, _ := app.Test(req)

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Errorf("Se esperaba status 401 Unauthorized, se obtuvo %d", resp.StatusCode)
		}
	})
}
