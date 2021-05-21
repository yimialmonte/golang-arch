package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string `json:"first"`
}

func main() {
	p1 := person{
		First: "Yelmi",
	}

	p2 := person{
		First: "Juan",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Json Format", string(bs))

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Back to Go structure", xp2)
}
