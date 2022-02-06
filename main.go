package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jagch/crud-goroutines/bd"
	"github.com/jagch/crud-goroutines/routers"
)

func main() {

	if bd.VerificarConexion() {
		log.Fatal("Sin conexión a la BD.")
		return
	}

	//sigoa
	sigoa := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // this is the default limit of 4MB
	})

	//cors-------------------------	-------------------------------------------------------------------------------------
	sigoa.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	//------------------------------------------------------------------------------------------------------------------

	routers.SetupRoutes(sigoa)

	//fmt.Print(runtime.NumCPU(), runtime.GOMAXPROCS(-1))

	sigoa.Listen("127.0.0.1:6000")
}
