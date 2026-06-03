package main

import (
	_ "api-go/docs"
	"api-go/internal/application"
	"api-go/internal/infrastructure/client"
	infraHttp "api-go/internal/infrastructure/http"
	"log"
	"os"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found: %v, reading from environment system", err)
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	grpcAddress := os.Getenv("NODE_GRPC_ADDRESS")

	nodeClient, err := client.NewNodeGRPCClient(grpcAddress)

	authService := application.NewAuthService()

	matrixService := application.NewMatrixService(nodeClient)

	authHandler := infraHttp.NewAuthHandler(authService)
	matrixHandler := infraHttp.NewMatrixHandler(matrixService)

	infraHttp.SetupRoutes(app, authHandler, matrixHandler)

	app.Get("/swagger/*", swaggo.New())

	log.Printf("Starting API on Go on port %s...", os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
