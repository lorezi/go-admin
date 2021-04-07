package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lorezi/go-admin/controllers"
	"github.com/lorezi/go-admin/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.AuthUser)
	app.Post("/api/logout", controllers.Logout)

}
