package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/TanishkN/xunami-server/auth"
	"github.com/TanishkN/xunami-server/db"
	"github.com/TanishkN/xunami-server/server/controller"
	"github.com/TanishkN/xunami-server/server/router"
	"google.golang.org/api/option"
)

func main() {
	// Set up the database connection
	dbConn, err := db.ConnectToDatabase("host=localhost port=5432 user=postgres dbname=postgres password=member sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer dbConn.Close()

	/*another way to create a connection to local firebase and create a doc:
		ctx:=context.Background()
		conf:=&firebase.Config(ProjectID:"xunami-userbase")
		app,err:=firebase.NewApp(ctx,conf)

		client,err:=app/Firestore(ctx)
		ref:=client.Collection("todos").NewDoc()
		result,err:=ref.Set(ctx,map[string]interface{}{
		"title":"Name"
		"text1":"SOMETHING"
		})
	programmer here wants to create todos
	*/

	// Create a Firebase app instance
	opt := option.WithCredentialsFile("./service-account-key.json")
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

	// Set up the authentication service and middleware
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
