package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const raw string = "json_data/raw_items.json"
const item_source string = "json_data/source_items.json"
const updated_items string = "json_data/updated_items.json"

func main() {
	// add flags

	log.SetFlags(log.Lshortfile)

	// Setting port value so that it will work on heroku but will be set at 5000 for local use.

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	//-------- Setting up item database. --------
	items, err := LoadItems(item_source)
	if err != nil {
		log.Fatal((err))
	}

	// // -------- Setting up character database. --------
	characters, err := LoadCharacters("json_data/characters.json")
	if err != nil {
		log.Fatal((err))
	}

	server := Server{
		items:      items,
		characters: characters,
	}

	http.HandleFunc("/items/", server.HandleItems)
	http.HandleFunc("/items/admin/reload", server.ReloadItems)
	http.HandleFunc("/characters", server.DisplayCharacters)
	http.HandleFunc("/", ServeIndex)
	http.ListenAndServe(":"+port, nil)
	// SetIDs(Items)

	// msg := "Welcome to the Handy Haversack Web Server\n\nTo use this web server, place a call to https://handyhaversack.herokuapp.com/items/ and place the item name or ID (int) after 'items/'."

	// Firebase connection
	ctx := context.Background()
	sa := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	iter := client.Collection("Items").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	// ctx := context.Background()
	// config := &firebase.Config{
	// 	DatabaseURL: "https://dnd-project-b2aa4.firebaseio.com",
	// }
	// app, err := firebase.NewApp(ctx, config, option.WithCredentialsFile("credentials.json"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client, err := app.Database(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// q := client.NewRef()
	// var result map[string]Item
	// if err := q.Get(ctx, &result); err != nil {
	// 	log.Fatal(err)
	// }

	// // Results will be logged in no specific order.
	// for key, acc := range result {
	// 	log.Printf("%s => %v\n", key, acc)
	// }

	// http.Handle("/", &defaultHandler{Message: msg})
	// http.HandleFunc("/items/", itemHandler)
	// http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
