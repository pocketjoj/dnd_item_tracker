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

	http.HandleFunc("/items/name/", Items.NameHandler)
	http.HandleFunc("/items/", Items.IDHandler)
	http.HandleFunc("/characters", Crew.Display)
	http.Handle("/", &defaultHandler{Message: msg})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

/*
TO DO LIST

-- Format certain fields (per notes in Item Struct above).
-- Create Character structure.
-- Add method for adding item(s)
-- Add method for removing item(s)
-- Add method for moving items between characters.
-- Designate handlers for all of the above methods.
-- Figure out if we want to plug in database for this.
-- Look into refactoring to make code cleaner... use different packages?
-- Review comments; ensure clarity and brevity.
-- Write concise readme for project.
-- Decide whether ItemList should be map or slice.

*/
