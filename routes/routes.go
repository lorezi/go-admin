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

	app.Patch("/api/users/profile", controllers.UpdateInfo)
	app.Patch("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.AuthUser)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.Users)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Patch("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.Roles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Patch("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/permissions", controllers.GetPermissions)

	app.Get("/api/products", controllers.GetProducts)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Patch("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

}
