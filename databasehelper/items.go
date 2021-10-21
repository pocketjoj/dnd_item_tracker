package databasehelper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Interface with methods that will apply to all database types (e.g. items, characters, etc.)

type DB interface {
	SetIDs()
	GetItemByID(id int) *Item
	GetItemByName(n string) *Item
	AddItem(data string, items []Item) []Item
	idHandler(w http.ResponseWriter, r *http.Request)
	NameHandler(w http.ResponseWriter, r *http.Request)
}

// Main struct that will handle data; each item will have some of the following properties, so we can then call these properties as needed.

type Item struct {
	Name       string                 `json:"name,omitempty"`
	Sources    map[string]interface{} `json:"sources,omitempty"`
	Rarity     string                 `json:"rarity,omitempty"`
	Entries    []interface{}          `json:"entries,omitempty"`
	Attunement string                 `json:"reqAttune,omitempty"` //Attunement requires formatting based on entry.
	Type       string                 `json:"type,omitempty"`
	Properties []interface{}          `json:"properties,omitempty"`
	Damage     map[string]interface{} `json:"damage,omitempty"`
	Tier       string                 `json:"tier,omitempty"`
	Srd        interface{}            `json:"srd,omitempty"`     //Srd has bool and string - will need switch case.
	Charges    interface{}            `json:"charges,omitempty"` //Charges appears to contain a mixture of strings and numbers; will need to check this.
	Image      bool                   `json:"image,omitempty"`
	Range      string                 `json:"range,omitempty"`
	Container  bool                   `json:"container,omitempty"`
	Extends    map[string]interface{} `json:"extends,omitempty"`
	ID         int                    `json:"id"`
}

type ItemList []Item

// ItemList Helper Methods --------------------

// Sets all IDs on initial load of program.
func (i ItemList) SetIDs() {
	for index := range i {
		i[index].ID = index + 1
	}
}

// Pass in ID, get item returned.
func (i ItemList) GetItemByID(id int) *Item {
	for _, i := range i {
		if i.ID == id {
			return &i
		}
	}
	fmt.Println("ID not found")
	return nil
}

//Pass in name, get item returned.
func (i ItemList) GetItemByName(n string) *Item {
	for _, i := range i {
		if strings.EqualFold(i.Name, n) {
			return &i
		}
	}
	fmt.Println("Item not found")
	return nil
}

//Add item by passing in json data.
func (i ItemList) AddItem(data string) []Item {
	bytes := []byte(data)
	var item Item
	err := json.Unmarshal(bytes, &item)
	if err != nil {
		log.Fatal(err)
	}
	item.ID = len(i) + 1
	return append(i, item)
}

// ItemList Handler Methods

func (i ItemList) IDHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "items/")
	itemRequest, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := i.GetItemByID(itemRequest)
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

func (i ItemList) NameHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "?name=")
	itemRequest, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	item := i.GetItemByID(itemRequest)
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
