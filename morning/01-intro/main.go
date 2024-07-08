package main

import (
	"encoding/json"
	"log"
	"os"
)

type person struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func findPerson(people []*person, name string) (*person, bool) {
	for _, p := range people {
		if p.name == name {
			return p, true
		}
	}
	return nil, false
}

func main() {
	var people []*person
	data, err := os.ReadFile("people.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &people); err != nil {
		log.Fatal(err)
	}

	name := os.Args[1]
	p, ok := findPerson(people, name)
	if !ok {
		log.Fatal("person not found")
	}
	log.Default().Println("person found:", p.name)
}
