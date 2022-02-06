package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jagch/crud-goroutines/handler"
)

//SetupRoutes agrupa las rutas
func SetupRoutes(sigoa *fiber.App) {
	sigoa.Post("/doc", handler.DocImport)
}
