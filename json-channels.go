package main

import (
	"log"
	"fmt"
	"strings"
	"encoding/json"
	"io"
)

type Golfer struct {
	Name string
	Hcp  int
}

func readGolfers(r io.Reader, golfers chan<- Golfer) {
	defer close(golfers)
	dec := json.NewDecoder(r)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var g Golfer
		// decode an array value Golfer
		err := dec.Decode(&g)
		if err != nil {
			log.Fatal(err)
		}
		golfers <- g
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%T: %v\n", t, t)
}

func main() {
	s := `[
	{
	"Name": "Jon",
	"Hcp": 24
	},
	{
	"Name": "Tiger",
	"Hcp": -2
	},
	{
	"Name": "Annika",
	"Hcp": 0
	},
	{
	"Name": "Bob",
	"Hcp": 44
	},
	{
	"Name": "Alice",
	"Hcp": 9
	}
]`

	golfers := make(chan Golfer, 2)
	go readGolfers(strings.NewReader(s), golfers)

	for g := range golfers {
		fmt.Printf("%v: %v\n", g.Name, g.Hcp)
	}
}
