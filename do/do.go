package do

import (
	"encoding/json"
	"fmt"
	"log"
)

type Place struct {
	id         int
	name       string
	place_type string
	image      string
	open       string
	close      string
	goods      []Good
}

type Good struct {
	id          int
	name        string
	price       float32
	ingredients []Ingredient
}

type Ingredient struct {
	id   int
	name string
}

func do(data []byte) {
	var p Place
	err := json.Unmarshal(data, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

}

//Ingredients_goods
//id
//I_id
//G_id
