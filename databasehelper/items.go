package databasehelper

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

// Interface with methods that will apply to all database types (e.g. items, characters, etc.)
//Not sure if I need this at all...
// type DB interface {
// 	SetIDs()
// 	GetItemByID(id int) *Item
// 	GetItemByName(n string) *Item
// 	AddItem(data string, items []Item) []Item
// 	idHandler(w http.ResponseWriter, r *http.Request)
// 	NameHandler(w http.ResponseWriter, r *http.Request)
// }

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

type ItemList map[int]Item

// ItemList Helper Methods --------------------

func (i *ItemList) RefreshList() {
	c_data, err := os.Open("json_data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	defer c_data.Close()

	c_byteValue, err := ioutil.ReadAll(c_data)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(c_byteValue, &i)
	if err != nil {
		log.Fatal(err)
	}
}

// Pass in ID, get item returned.
func (i ItemList) GetItemByID(id int) Item {
	return i[id]
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
func (i ItemList) AddItem(data string) ItemList {
	bytes := []byte(data)
	var item Item
	err := json.Unmarshal(bytes, &item)
	if err != nil {
		log.Fatal(err)
	}
	i[len(i)+1] = item
	return i
}

func (i ItemList) ItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	query := r.URL.Query()
	if len(query) == 0 {
		urlPathSegments := strings.Split(r.URL.Path, "items/")
		itemRequest, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		item = i.GetItemByID(itemRequest)
	} else {
		item = *i.GetItemByName(query["name"][0])
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
