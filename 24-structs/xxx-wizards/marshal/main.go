package main

import (
	"encoding/json"
	"fmt"
)

// Wizard is one of the greatest of people
type Wizard struct {
	// name won't be marshalled (should be exported)
	Name     string `json:name`
	Lastname string `json:"-"`
	Nick     string `json:"nick"`
}

func main() {
	wizards := []Wizard{
		{Name: "Albert", Lastname: "Einstein", Nick: "emc2"},
		{Name: "Isaac", Lastname: "Newton", Nick: "apple"},
		{Name: "Stephen", Lastname: "Hawking", Nick: "blackhole"},
		{Name: "Marie", Lastname: "Curie", Nick: "radium"},
		{Name: "Charles", Lastname: "Darwin", Nick: "fittest"},
	}

	bytes, err := json.Marshal(wizards)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(bytes))
}
