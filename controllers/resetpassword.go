package controllers

import (
	"log"
	"os"
	"regexp"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type PasswordResetRequest struct {
	Email string `json:"email"`
}

var FireAuth *auth.Client

func isValidEmail(email string) bool {
	// Expressão regular para validar o formato do email
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}

func sendPasswordResetEmail(email, link string) error {
	from := mail.NewEmail("Go Messenger", "gomessengerapp@gmail.com")
	subject := "Password Reset Request"
	to := mail.NewEmail("", email)
	plainTextContent := "Click the link to reset your password: " + link
	htmlContent := "<p>Click the link to reset your password: <a href=\"" + link + "\">Reset Password</a></p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	log.Printf("Email sent: %v", response.StatusCode)
	return nil
}

func PasswordReset(c *fiber.Ctx) error {
	var request PasswordResetRequest

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	if request.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email required"})
	}

	if !isValidEmail(request.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email format"})
	}

	authClient := FireAuth // Reuse the initialized Firebase Auth client

	link, err := authClient.PasswordResetLink(c.Context(), request.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate password reset link")
	}

	if err := sendPasswordResetEmail(request.Email, link); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to send password reset email")
	}

	return c.JSON(fiber.Map{
		"message": "Password reset link sent",
	})
}
