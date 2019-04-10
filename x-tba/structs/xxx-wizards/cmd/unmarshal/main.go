package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/inancgumus/learngo/x-tba/structs/xxx-json/wizards"
)

func main() {
	file, err := ioutil.ReadFile("../marshal/wizards.json")
	if err != nil {
		panic(err)
	}

	wizards := make([]wizards.Wizard, 10)
	if json.Unmarshal(file, &wizards) != nil {
		panic(err)
	}

	fmt.Printf("%-15s %-15s\n%s",
		"Name", "Nick", strings.Repeat("=", 25))

	for _, w := range wizards {
		fmt.Printf("%-15s %-15s\n", w.Name, w.Nick)
	}
}
