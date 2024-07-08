// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type person struct {
	name string
	id   int
	age  int
}

func findPerson(people []*person, name string) *person {
	for i := range people {
		if people[i].name == name {
			return people[i]
		}
	}
	return nil
}

func populate(people []*person) []*person {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	names := []string{"Alicia", "Alex", "Erica", "Derek", "Julian", "Bob"}
	for i := range names {
		people = append(people, &person{
			id:   r.Int(),
			name: names[i],
			age:  rand.Intn(100),
		})
	}
	return people
}

func main() {
	people := make([]*person, 0, 6)
	people = populate(people)
	p := findPerson(people, "Dеrek")
	if p == nil {
		fmt.Println("could not find Derek")
	} else {
		fmt.Println("found our person")
	}
}
