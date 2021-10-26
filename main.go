package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pocketjoj/dnd_item_tracker/databasehelper"
)

// Handler functions --------------------

// Basic default handler
type defaultHandler struct {
	Message string
}

func (f *defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func main() {
	log.SetFlags(log.Lshortfile)

	// -------- Setting up item database. --------
	var Items databasehelper.ItemList
	Items.RefreshList()

	// -------- Setting up character database. --------
	Crew := make(databasehelper.CharacterList)
	Crew.RefreshList()

	msg := "Welcome to the Handy Haversack Web Server\n\nTo use this web server, place a call to https://handyhaversack.herokuapp.com/items/ and place the item ID (int) after 'items/'."

	// Setting port value so that it will work on heroku but will be set at 5000 for local use.
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	http.HandleFunc("/items/", Items.ItemHandler)
	http.HandleFunc("/characters", Crew.Display)
	http.Handle("/", &defaultHandler{Message: msg})
	http.ListenAndServe(":"+port, nil)
	// Below is for local testing.
	// http.ListenAndServe(":5000", nil)

}

/*
TO DO LIST

-- Format certain fields (per notes in Item Struct above).
-- Designate handlers for character methods.
-- Review comments; ensure clarity and brevity.
-- Write concise readme for project.

*/
