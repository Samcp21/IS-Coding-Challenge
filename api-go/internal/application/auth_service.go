package application

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(username, password string) (string, error) {
	expectedUser := os.Getenv("API_USER")
	expectedPass := os.Getenv("API_PASS")

	if username != expectedUser || password != expectedPass {
		return "", errors.New("credenciales inválidas")
	}
	claims := jwt.MapClaims{
		"user": username,
		"role": "admin",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
