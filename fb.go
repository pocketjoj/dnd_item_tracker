package main

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Item struct {
	Name       string      `json:"name,omitempty"`
	Sources    interface{} `json:"sources,omitempty"`
	Rarity     string      `json:"rarity,omitempty"`
	Entries    interface{} `json:"entries,omitempty"`
	Attunement interface{} `json:"attunement,omitempty"`
	Type       string      `json:"type,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
	Damage     interface{} `json:"damage,omitempty"`
	Tier       string      `json:"tier,omitempty"`
	Srd        interface{} `json:"srd,omitempty"`
	Charges    interface{} `json:"charges,omitempty"` // Charges info can remain in different types, conversion is not needed.
	Image      bool        `json:"image,omitempty"`
	Range      string      `json:"range,omitempty"`
	Container  bool        `json:"container,omitempty"`
	Extends    interface{} `json:"extends,omitempty"` // It is fine for this data to remain as json, as it does not really need to be converted.
	ID         string      `json:"id"`
}

type Character struct {
	Name      string `json:"name"`
	Inventory []Item `json:"inventory"`
}

// Firebase Item Functions

// I THINK below can be deleted - I don't really need to unmarsal Firestore data into my items. I can just marshal the map[string] interface to json and pass it along.
// func FS_Unmarshal(x map[string]interface{}) (Item, error) {
// 	json_string, err := json.Marshal(x)
// 	if err != nil {
// 		return Item{}, err
// 	}
// 	var i Item
// 	err = json.Unmarshal(json_string, &i)
// 	if err != nil {
// 		return Item{}, nil
// 	}
// 	return i, nil
// }

func ListItems(c firestore.Client) ([]byte, error) {
	iter := c.Collection("items").Documents(context.Background())
	var list []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		m := make(map[string]interface{})
		m["id"] = doc.Ref.ID
		m["name"] = doc.Data()["name"]
		list = append(list, m)
	}
	fmt.Println(len(list))
	r, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetItemByID(i string, c firestore.Client) ([]byte, error) {
	dsnap, err := c.Collection("items").Doc(i).Get(context.Background())
	if err != nil {
		return nil, err
	}
	m := dsnap.Data()
	m["id"] = i

	item, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// func GetItemsByID(i []string, c firestore.Client) ([]Item, error) {
// 	dsnap, err := c.Collection("items").GetAll(context.Background())
// 	if err != nil {
// 		return Item{}, err
// 	}
// 	var item Item
// 	item, err = FS_Unmarshal(dsnap.Data())
// 	return item, nil
// }
