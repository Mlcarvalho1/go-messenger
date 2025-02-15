package routes

import (
	"log"
	//"os"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	//"github.com/sendgrid/sendgrid-go"
	//"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type PasswordResetRequest struct {
	Email string `json:"email"`
}

var FireAuth *auth.Client

//func sendPasswordResetEmail(email, link string) error {
//from := mail.NewEmail("Your App Name", "no-reply@yourapp.com")
//subject := "Password Reset Request"
//to := mail.NewEmail("", email)
//plainTextContent := "Click the link to reset your password: " + link
//htmlContent := "<p>Click the link to reset your password: <a href=\"" + link + "\">Reset Password</a></p>"
//message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
//client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
//response, err := client.Send(message)
//if err != nil {
//    return err
//}
//log.Printf("Email sent: %v", response.StatusCode)
//return nil
//}

func PasswordReset(c *fiber.Ctx) error {
	var request PasswordResetRequest

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	authClient := FireAuth // Reuse the initialized Firebase Auth client

	link, err := authClient.PasswordResetLink(c.Context(), request.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate password reset link")
	}

	// Aqui você pode enviar o link por email usando um serviço de email de sua escolha
	log.Printf("Password reset link: %s", link)

	return c.JSON(fiber.Map{
		"message": "Password reset link sent",
	})
}
