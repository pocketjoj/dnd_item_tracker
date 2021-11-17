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
	Name       string          `json:"name,omitempty"`
	Sources    json.RawMessage `json:"sources,omitempty"`
	Rarity     string          `json:"rarity,omitempty"`
	Entries    json.RawMessage `json:"entries,omitempty"`
	Attunement json.RawMessage `json:"attunement,omitempty"`
	Type       string          `json:"type,omitempty"`
	Properties json.RawMessage `json:"properties,omitempty"`
	Damage     json.RawMessage `json:"damage,omitempty"`
	Tier       string          `json:"tier,omitempty"`
	Srd        json.RawMessage `json:"srd,omitempty"`
	Charges    json.RawMessage `json:"charges,omitempty"` // Charges info can remain in different types, conversion is not needed.
	Image      bool            `json:"image,omitempty"`
	Range      string          `json:"range,omitempty"`
	Container  bool            `json:"container,omitempty"`
	Extends    json.RawMessage `json:"extends,omitempty"` // It is fine for this data to remain as json, as it does not really need to be converted.
	Custom     bool            `json:"custom"`
	ID         int             `json:"id"`
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

// Initially used to help format JSON -- keeping method available for now. This is used if the raw data is updated and I need to parse it again using https://handyhaversack.herokuapp.com/items/refresh

func RefreshSourceItems(filename string) (map[int]Item, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not refresh items: %w", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", f.Name(), err)
	}
	var items []Item
	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, fmt.Errorf("could not refresh items: %w", err)
	}
	c := 1

	formatted_items := make(map[int]Item)

	for _, i := range items {
		i.ID = c
		formatted_items[i.ID] = i
		c++
	}

	j, err := json.Marshal(formatted_items)
	if err != nil {
		return nil, fmt.Errorf("problem marshalling formatted items into json: %w", err)
	}

	err = ioutil.WriteFile(item_source, j, 0644)
	if err != nil {
		return nil, fmt.Errorf("could not write json file: %w", err)
	}

	return formatted_items, nil
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
func AddItem(b []byte, i map[int]Item) (map[int]Item, error) {
	var item Item
	err := json.Unmarshal(b, &item)
	if err != nil {
		return nil, fmt.Errorf("could not add item: %w", err)
	}
	item.ID = len(i) + 1
	i[item.ID] = item

	j, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("problem marshalling formatted items into json: %w", err)
	}
	err = ioutil.WriteFile(updated_items, j, 0644)
	if err != nil {
		return nil, fmt.Errorf("could not write json file: %w", err)
	}
	return i, nil
}
