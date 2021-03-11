package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

//go:embed views/pages/*.html views/layouts/*.html
var pagesFs embed.FS

func main() {
	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: django.NewFileSystem(http.FS(pagesFs), ".html"),
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("views/pages/login", fiber.Map{
			"message": "Login data injection works!",
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("views/pages/dashboard", fiber.Map{
			"message": "Dashboard data injection works!",
		})
	})

	log.Fatal(app.Listen(":5555"))
}
