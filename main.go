package main

import (
	"FIBER/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New()

	var engine = html.New("./views", ".html")

	app = fiber.New(fiber.Config{Views: engine})
	app.Static("/static", "./static")

	app.Use(loginMiddleware)

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func loginMiddleware(c *fiber.Ctx) error {
	log.Printf("Solicitud recibida: %s %s", c.Method(), c.Path())
	return c.Next()
}
