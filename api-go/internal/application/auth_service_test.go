package application

import (
	"os"
	"testing"
)

func TestAuthService_Login(t *testing.T) {
	os.Setenv("API_USER", "admin")
	os.Setenv("API_PASS", "secreto123")
	os.Setenv("JWT_SECRET", "superclave")
	defer os.Clearenv()

	authService := NewAuthService()

	t.Run("Login Exitoso", func(t *testing.T) {
		token, err := authService.Login("admin", "secreto123")

		if err != nil {
			t.Errorf("No se esperaba error, se obtuvo: %v", err)
		}
		if token == "" {
			t.Error("Se esperaba un token JWT, se obtuvo un string vacío")
		}
	})

	t.Run("Login Fallido - Credenciales Inválidas", func(t *testing.T) {
		token, err := authService.Login("admin", "clave_falsa")

		if err == nil {
			t.Error("Se esperaba un error por credenciales inválidas, pero no ocurrió")
		}
		if token != "" {
			t.Errorf("Se esperaba token vacío tras un error, se obtuvo: %s", token)
		}
	})
}
