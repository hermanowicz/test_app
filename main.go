package main

import (
	"embed"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsfs embed.FS

func main() {

	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		BodyLimit:             1 * 1024 * 1024,
		CaseSensitive:         true,
		StrictRouting:         true,
		DisableStartupMessage: false,
		ReadTimeout:           15 * time.Second,
		WriteTimeout:          15 * time.Second,
		IdleTimeout:           15 * time.Second,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		server, _ := os.Hostname()
		return c.Render("views/index", fiber.Map{"Server": server})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	app.Listen(":8080")
}
