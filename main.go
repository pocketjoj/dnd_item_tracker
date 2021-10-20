package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Main struct that will handle data; each "item" will have some of the following properties, so we can then call these properties as needed.

type Item struct {
	Name    string                 `json:"name"`
	Sources map[string]interface{} `json:"sources"`
	Rarity  string                 `json:"rarity"`
	Entries []interface{}          `json:"entries"`
	//Attunement requires formatting based on entry.
	Attunement string                 `json:"reqAttune"`
	Type       string                 `json:"type"`
	Properties []interface{}          `json:"properties"`
	Damage     map[string]interface{} `json:"damage"`
	Tier       string                 `json:"tier"`
	//Srd has bool and string - will need switch case.
	Srd interface{} `json:"srd"`
	//Charges appears to contain a mixture of strings and numbers; will need to check this.
	Charges   interface{}            `json:"charges"`
	Image     bool                   `json:"image"`
	Range     string                 `json:"range"`
	Container bool                   `json:"container"`
	Extends   map[string]interface{} `json:"extends"`
	ID        int                    `json:"id"`
}

// Declaring variable outside of func main so that functions will have access to it.
var Items []Item

//Helper Functions

// Sets all IDs on initial load of program.
func SetIDs(i []Item) {
	for index := range i {
		i[index].ID = index + 1
	}
}

// Pass in ID, get item returned.
func GetItemByID(id int) *Item {
	for _, i := range Items {
		if i.ID == id {
			return &i
		}
	}
	fmt.Println("ID not found")
	return nil
}

//Pass in name, get item returned.
func GetItemByName(n string) *Item {
	for _, i := range Items {
		if i.Name == n {
			return &i
		}
	}
	fmt.Println("Item not found")
	return nil
}

//Add item by passing in json data.
func AddItem(data string, items []Item) []Item {
	bytes := []byte(data)
	var i Item
	err := json.Unmarshal(bytes, &i)
	if err != nil {
		log.Fatal(err)
	}
	i.ID = len(items) + 1
	return append(items, i)
}

// Parsing if request is using item ID or item Name

func parseRequest(s string) *Item {
	itemRequest, err := strconv.Atoi(s)
	if err != nil {
		return GetItemByName(s)
	}
	if err == nil {
		return GetItemByID(itemRequest)
	}
	return nil
}

// Basic default handler

type defaultHandler struct {
	Message string
}

func (f *defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

// Handler function for requests made for items
func itemHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "items/")
	item := parseRequest(urlPathSegments[len(urlPathSegments)-1])
	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
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

func main() {
	log.SetFlags(log.Lshortfile)

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

	SetIDs(Items)

	msg := "Welcome to the Handy Haversack Web Server\n\nTo use this web server, place a call to https://handyhaversack.herokuapp.com/items/ and place the item name or ID (int) after 'items/'."

	http.Handle("/", &defaultHandler{Message: msg})
	http.HandleFunc("/items/", itemHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
