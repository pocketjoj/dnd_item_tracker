package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Basic default handler
type defaultHandler struct {
	Message string
}

// Struct for handling items and characters.
type Server struct {
	items      map[int]Item
	characters map[int]Character
}

func (s Server) HandleItems(w http.ResponseWriter, r *http.Request) {
	var item Item
	var err error
	query := r.URL.Query()
	if len(query) == 0 {
		urlPathSegments := strings.Split(r.URL.Path, "items/")
		itemRequest, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		item = s.items[itemRequest]
	} else {
		item, err = GetItemByName(query["name"][0], s.items)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	switch r.Method {
	case http.MethodGet:
		itemJSON, err := json.Marshal(item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(itemJSON)
		fmt.Println("Item retrieved")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s Server) DisplayCharacters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		charactersJSON, err := json.Marshal(s.characters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(charactersJSON)
		fmt.Println("Characters retrieved successfully")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
