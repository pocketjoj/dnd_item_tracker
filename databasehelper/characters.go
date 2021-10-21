package databasehelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Character struct {
	Name      string `json:"name"`
	Inventory []Item `json:"inventory"`
}

func (c Character) CheckInventory() []Item {
	return c.Inventory
}

func (c *Character) AddItemByID(id int, i ItemList) {
	c.Inventory = append(c.Inventory, i[id])
}

func (c *Character) RemoveItemByID(id int) {
	for i, value := range c.Inventory {
		if value.ID == id {
			c.Inventory = append(c.Inventory[:i], c.Inventory[(i+1):]...)
			break
		}
	}
}

type CharacterList map[int]Character

func (cl *CharacterList) RefreshList() {
	c_data, err := os.Open("json_data/characters.json")
	if err != nil {
		log.Fatal(err)
	}
	defer c_data.Close()

	c_byteValue, err := ioutil.ReadAll(c_data)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(c_byteValue, &cl)
	if err != nil {
		log.Fatal(err)
	}
}

//Unmarshaling json data into a collection of Characters.
