package database

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitFirebaseApp() *firebase.App {
	serviceAccount, fileExi := os.LookupEnv("SERVICE_ACCOUNT_JSON")

	if !fileExi {
		log.Fatalf("Please provide valid firebbase auth credential json!")
	}

	opt := option.WithCredentialsFile(serviceAccount)
	// Initialize the Firebase app
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	return app
}
