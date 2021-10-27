package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

// Use json.RawMessage for the interface information. Will need to write a function to decode those items.

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

// ItemList Helper Methods --------------------

func LoadItems(filename string) (map[int]Item, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not load items: %w", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", f.Name(), err)
	}
	items := make(map[int]Item)
	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, fmt.Errorf("could not load items: %w", err)
	}
	return items, nil
}

//Pass in name, get item returned.
func GetItemByName(n string, i map[int]Item) (Item, error) {
	for _, i := range i {
		if strings.EqualFold(i.Name, n) {
			return i, nil
		}
	}
	return Item{}, fmt.Errorf("item with that name not found")
}

//Add item by passing in json data.
func AddItem(data string, i map[int]Item) (map[int]Item, error) {
	bytes := []byte(data)
	var item Item
	err := json.Unmarshal(bytes, &item)
	if err != nil {
		return nil, fmt.Errorf("could not add item: %w", err)
	}
	i[len(i)+1] = item
	return i, nil
}
