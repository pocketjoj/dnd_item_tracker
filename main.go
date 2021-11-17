package main

import (
	"log"
	"net/http"
	"os"
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
	http.HandleFunc("/items/refresh", server.RefreshItems)
	http.HandleFunc("/characters", server.DisplayCharacters)
	http.HandleFunc("/", ServeIndex)
	http.ListenAndServe(":"+port, nil)
}

/*
TO DO LIST

-- Format certain fields (per notes in Item Struct above).
-- Designate handlers for character methods.
-- Review comments; ensure clarity and brevity.
-- Write concise readme for project.
-- Add html formats so web server pages look better??

*/
