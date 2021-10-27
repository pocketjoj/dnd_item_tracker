package main

import (
	"log"
	"net/http"
	"os"
)

func (f *defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func main() {
	// add flags

	log.SetFlags(log.Lshortfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// -------- Setting up item database. --------
	items, err := LoadItems("json_data/items.json")
	if err != nil {
		log.Fatal((err))
	}

	// -------- Setting up character database. --------
	characters, err := LoadCharacters("json_data/characters.json")
	if err != nil {
		log.Fatal((err))
	}

	server := Server{
		items:      items,
		characters: characters,
	}

	// Looking into adding html template.
	// t, err := os.Open("index.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer t.Close()

	// tmpl, err := ioutil.ReadAll(t)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// msg := template.New(string(tmpl))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Setting port value so that it will work on heroku but will be set at 5000 for local use.

	http.HandleFunc("/items/", server.HandleItems)
	http.HandleFunc("/characters", server.DisplayCharacters)
	http.Handle("/", &defaultHandler{Message: "Hello World"})
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
