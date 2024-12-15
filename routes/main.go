package routes

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	gofiberfirebaseauth "github.com/sacsand/gofiber-firebaseauth"
	"google.golang.org/api/option"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/", logger.New())

	serviceAccount, fileExi := os.LookupEnv("SERVICE_ACCOUNT_JSON")

	opt := option.WithCredentialsFile(serviceAccount)
	fireApp, _ := firebase.NewApp(context.Background(), nil, opt)

	if !fileExi {
		log.Fatalf("Please provide valid firebbase auth credential json!")
	}

	app.Use(gofiberfirebaseauth.New(gofiberfirebaseauth.Config{
		FirebaseApp: fireApp,
		IgnoreUrls:  []string{"POST::/user"},
	}))

	// Register feature-based routes
	UserRoutes(api)
}
