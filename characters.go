package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Character struct {
	Name      string `json:"name"`
	Inventory []Item `json:"inventory"`
}

//Character Helper Functions

func (c Character) CheckInventory() []Item {
	return c.Inventory
}

func AddItemByID(id int, i map[int]Item, c Character) ([]Item, error) {
	value, ok := i[id]
	if ok == true {
		c.Inventory = append(c.Inventory, value)
		return c.Inventory, nil
	}
	return nil, fmt.Errorf("item with that ID could not be located")
}

func RemoveItemByID(id int, c Character) ([]Item, error) {
	for i, value := range c.Inventory {
		if value.ID == id {
			c.Inventory = append(c.Inventory[:i], c.Inventory[(i+1):]...)
			return c.Inventory, nil
		}
	}
	return nil, fmt.Errorf("item id given was not found within this character's inventory")
}

func GiveItem(id int, c1 Character, c2 Character, i map[int]Item) ([]Item, []Item, error) {
	c1Inventory, err := RemoveItemByID(id, c1)
	if err != nil {
		return nil, nil, fmt.Errorf("problem with removing item from %s: %w", c1.Name, err)
	}
	c2Inventory, err := AddItemByID(id, i, c2)
	if err != nil {
		return nil, nil, fmt.Errorf("problem with add item to %s: %w", c1.Name, err)
	}
	return c1Inventory, c2Inventory, nil
}

//Unmarshaling json data into a collection of Characters.
func LoadCharacters(filename string) (map[int]Character, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not load characters: %w", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read character data: %w", err)
	}
	characters := make(map[int]Character)
	err = json.Unmarshal(b, &characters)
	if err != nil {
		return nil, fmt.Errorf("could not load character data: %w", err)
	}
	return characters, nil
}
