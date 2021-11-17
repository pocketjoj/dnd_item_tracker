package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Struct that provides access for handling items and characters.

type Server struct {
	items      map[int]Item
	characters map[int]Character
}

// Server Method to either post the json for an Item by ID or Name (using GET) or allow adding of a custom item formatted in JSON (POST).

func (s *Server) HandleItems(w http.ResponseWriter, r *http.Request) {
	var item Item
	var err error
	query := r.URL.Query()
	urlPathSegments := strings.Split(r.URL.Path, "items/")
	switch r.Method {
	case http.MethodGet:
		if len(query) == 0 {
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
		itemJSON, err := json.Marshal(item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(itemJSON)
		fmt.Println("Item retrieved")
	case http.MethodPost:
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = AddItem(b, s.items)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Will display all characters and inventories -- would be good to use templates for this maybe.

func (s *Server) DisplayCharacters(w http.ResponseWriter, r *http.Request) {
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

// Default Page Handler

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

// Admin function to re-format raw JSON data is source updates their JSON files.

func (s *Server) ReloadItems(w http.ResponseWriter, r *http.Request) {
	i, err := RefreshSourceItems(raw)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	s.items = i
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Request Successful. Please return to home page at ' https://handyhaversack.herokuapp.com/'."))

}
