package routes

import (
	// "app/middleware"
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router) {
	user := api.Group("/user")
	user.Get("/", controllers.GetUser)
	// user.Post("/", controllers.CreateUser)
	// user.Patch("/:id", middleware.Protected(), controllers.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), controllers.DeleteUser)
}
