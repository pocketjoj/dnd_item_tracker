package main

import (
	"fmt"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
)

type Server struct {
	db firestore.Client
}

// Default Page Handler
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// Server Method to either post the json for an Item by ID or Name (using GET) or allow adding of a custom item formatted in JSON (POST).
func (s *Server) HandleItems(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "items/")
	id := urlPathSegments[1]
	switch r.Method {
	case http.MethodGet:
		item, err := GetItemByID(id, s.db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// itemJSON, err := json.Marshal(item)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		w.Header().Set("Content-Type", "application/json")
		w.Write(item)
		fmt.Println("Item retrieved")
	// case http.MethodPost:
	// 	b, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	_, err = AddItem(b, s.items)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		fmt.Println(err)
	// 		return
	// 	}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) GetAllItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		items, err := ListItems(s.db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(items)
		fmt.Println("Items sent")
	}
}

// Will display all characters and inventories -- would be good to use templates for this maybe.

// func (s *Server) DisplayCharacters(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		charactersJSON, err := json.Marshal(s.characters)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(charactersJSON)
// 		fmt.Println("Characters retrieved successfully")

// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }
