package routes

import (
	"go.messenger/controllers"

	"github.com/gofiber/fiber/v2"
)

func ResetPasswordRoutes(api fiber.Router) {
	resetPassword := api.Group("/resetpassword")

	resetPassword.Post("/", controllers.PasswordReset)
}
