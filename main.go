package main

import (
	"encoding/json"
	"io/ioutil"
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

	var Items databasehelper.ItemList

	// Opening json file and reading data from it.
	data, err := os.Open("items.json")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}

	//Unmarshaling json data into a collection of Items.
	err = json.Unmarshal(byteValue, &Items)
	if err != nil {
		log.Fatal(err)
	}

	Items.SetIDs()

	msg := "Welcome to the Handy Haversack Web Server\n\nTo use this web server, place a call to https://handyhaversack.herokuapp.com/items/ and place the item name or ID (int) after 'items/'."

	http.Handle("/", &defaultHandler{Message: msg})
	http.HandleFunc("/items/", Items.IDHandler)
	http.Handle("/items/?name=", &defaultHandler{Message: "this is working"})
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

*/
