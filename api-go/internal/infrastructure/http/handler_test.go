package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"api-go/internal/application"
	infraHttp "api-go/internal/infrastructure/http"

	"github.com/gofiber/fiber/v3"
)

func TestProcessMatrix_Integration(t *testing.T) {
	os.Setenv("API_USER", "root")
	os.Setenv("API_PASS", "root")
	os.Setenv("JWT_SECRET", "SECRET_KEY_INTERSEGURO")

	defer os.Clearenv()

	app := fiber.New()

	matrixService := application.NewMatrixService(nil)
	matrixHandler := infraHttp.NewMatrixHandler(matrixService)
	authService := application.NewAuthService()
	authHandler := infraHttp.NewAuthHandler(authService)

	infraHttp.SetupRoutes(app, authHandler, matrixHandler)

	token, err := authService.Login("root", "root")
	if err != nil {
		t.Fatalf("Falló el login de prueba: %v", err)
	}
	body := map[string]interface{}{
		"data": [][]float64{{1, 2}, {3, 4}},
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/matrix", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error al ejecutar la petición: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba status 200, se obtuvo %d", resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	if !bytes.Contains(respBody, []byte("factorization_qr")) {
		t.Errorf("La respuesta no contiene los resultados matemáticos: %s", string(respBody))
	}
}
