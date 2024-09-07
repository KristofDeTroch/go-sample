package main

import (
	"context"
	"go-sample/auth"
	"go-sample/db"
	controller "go-sample/server/controllers"
	router "go-sample/server/routers"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	// Set up the database connection
	dbConn, err := db.ConnectToDatabase("host=localhost port=5432 user=root dbname=postgres password=root sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer dbConn.Close()

	opt := option.WithCredentialsFile("./credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Create a Firebase auth client instance
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase auth client: %v", err)
	}

	// Run database migrations
	dbConn.AutoMigrate(&auth.User{})

	// Set up the authentication service
	authService := &auth.AuthService{
		DB:       dbConn,
		FireAuth: authClient,
	}
	authController := controller.NewAuthController(authService)

	// Set up the HTTP router
	r := router.NewRouter(authController)

	// Start the server
	r.Run(":8080")
}
