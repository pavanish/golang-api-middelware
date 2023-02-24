package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/pavanish/go-react-jwt/database"
	"github.com/pavanish/go-react-jwt/routes"
)

func main() {

	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	//app.Use(middleware.New())

	routes.Setup(app)

	app.Listen(":3000")
}
