package routes

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SetupRoutes(app *fiber.App) {

	var renderer = Renderer()

	app.Get("/", handlerInicio)
	app.Get("/about", handlerAbout)
	app.Get("/saludo/:nombre", handlerSaludar)

	app.Get("/:page", dynamicPagesHandler(renderer))

	app.Get("/saludo/:nombre", func(c *fiber.Ctx) error {
		var nombre = c.Params("nombre")
		return c.SendString("Hola " + nombre + "!")
	})

	app.Post("/api/usuarios", func(c *fiber.Ctx) error {
		var usuario User
		if err := c.BodyParser(&usuario); err != nil {
			return err
		}
		return c.JSON(usuario)
	})
}

func Renderer() *template.Template {
	return template.Must(template.ParseGlob("views/*html"))
}

func dynamicPagesHandler(_ *template.Template) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var page = c.Params("page")
		if strings.HasSuffix(page, ".html") {
			page = strings.TrimSuffix(page, ".html")
		}
		if _, err := os.Stat("views/" + page + ".html"); err == nil {
			return c.Render(page, nil)
		}
		return c.Status(http.StatusNotFound).SendString("Pagina no encontrada")
	}
}

func handlerInicio(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":   "Mi aplicacion",
		"Headign": "Hola mundo",
		"Message": "Bienvenido a mi aplicacion web con Fiber y Plantillas HTML",
	})
	// return c.SendString(("Bienvenido a mi Aplicacion"))
}

func handlerAbout(c *fiber.Ctx) error {
	return c.SendString("Paigna de informacion sobre nuestra aplicacion")
}

func handlerSaludar(c *fiber.Ctx) error {
	var nombre = c.Params(("nombre"))
	return c.SendString("Hola " + nombre + "!")
}
