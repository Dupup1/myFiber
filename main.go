package main

import (
	"github.com/Dupup1/myFiber/database"
	"github.com/Dupup1/myFiber/database/migration"
	"github.com/Dupup1/myFiber/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/user", user.GetUsers)
	app.Get("/api/user/:id", user.GetUser)
	app.Post("/api/user", user.NewUser)
	app.Delete("/api/user/:id", user.DeleteUser)
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	database.InitDatabase()
	migration.RunMigration()

	setupRoutes(app)

	app.Static("/", "./public")

	app.Listen(":3000")
}
