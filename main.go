package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

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

// Declaring variable outside of func main so that helper functions will have access to it.
var Items []Item

//Helper Functions --------------------

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
		if strings.EqualFold(i.Name, n) {
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

// Handler functions --------------------

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

	// msg := "Welcome to the Handy Haversack Web Server\n\nTo use this web server, place a call to https://handyhaversack.herokuapp.com/items/ and place the item name or ID (int) after 'items/'."

	// Firebase connection
	ctx := context.Background()
	sa := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	iter := client.Collection("Items").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}

	// ctx := context.Background()
	// config := &firebase.Config{
	// 	DatabaseURL: "https://dnd-project-b2aa4.firebaseio.com",
	// }
	// app, err := firebase.NewApp(ctx, config, option.WithCredentialsFile("credentials.json"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client, err := app.Database(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// q := client.NewRef()
	// var result map[string]Item
	// if err := q.Get(ctx, &result); err != nil {
	// 	log.Fatal(err)
	// }

	// // Results will be logged in no specific order.
	// for key, acc := range result {
	// 	log.Printf("%s => %v\n", key, acc)
	// }

	// http.Handle("/", &defaultHandler{Message: msg})
	// http.HandleFunc("/items/", itemHandler)
	// http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

/*
TO DO LIST

-- Format certain fields (per notes in Item Struct above).
-- Create Character structure.
-- Add method for adding item(s)
-- Add method for removing item(s)
-- Add method for moving items between characters.
-- Designate handlers for all of the above methods.
-- Figure out if we want to plug in database for this.
-- Look into refactoring to make code cleaner... use different packages?
-- Review comments; ensure clarity and brevity.
-- Write concise readme for project.

*/
