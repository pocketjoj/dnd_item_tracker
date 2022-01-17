package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// Setting port value so that it will work on heroku but will be set at 5000 for local use.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	cred := `{
		"projectId": process.env.FIREBASE_PROJECT_ID,
		"private_key": process.env.FIREBASE_PRIVATE_KEY,
		"client_email": process.env.FIREBASE_CLIENT_EMAIL,
	}`

	var sa option.ClientOption

	if port == "5000" {
		sa = option.WithCredentialsFile("credentials.json")
	} else {
		sa = option.WithCredentialsJSON([]byte(cred))
	}

	// Initialize Firebase DB
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	server := Server{
		db: *client,
	}

	http.HandleFunc("/items/", server.HandleItems)
	http.HandleFunc("/items", server.GetAllItems)
	// http.HandleFunc("/characters", server.DisplayCharacters)
	http.HandleFunc("/", ServeIndex)
	http.ListenAndServe(":"+port, nil)
}

// msg := "Welcome to the Handy Haversack Web Server\n\nTo use this web server, place a call to https://handyhaversack.herokuapp.com/items/ and place the item name or ID (int) after 'items/'."
