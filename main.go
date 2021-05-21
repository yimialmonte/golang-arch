package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string `json:"first"`
}

const port = "127.0.0.1:8080"

func main() {
	// p1 := person{
	// 	First: "Yelmi",
	// }

	// p2 := person{
	// 	First: "Juan",
	// }

	// xp := []person{p1, p2}

	// bs, err := json.Marshal(xp)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Json Format", string(bs))

	// xp2 := []person{}

	// err = json.Unmarshal(bs, &xp2)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// fmt.Println("Back to Go structure", xp2)
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	http.ListenAndServe(port, nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Juan",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Enocde error: ", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decode data", err)
	}

	log.Println(p1)
}
