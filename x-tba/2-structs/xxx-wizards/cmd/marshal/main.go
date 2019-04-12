package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/inancgumus/learngo/x-tba/structs/xxx-json/wizards"
)

func main() {
	wizards := []wizards.Wizard{
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

	ioutil.WriteFile("wizards.json", bytes, 0644)

	//
	// PREVIOUSLY
	//

	// names := [...][3]string{
	// 	{"First Name", "Last Name", "Nickname"},
	// 	{"Albert", "Einstein", "emc2"},
	// 	{"Isaac", "Newton", "apple"},
	// 	{"Stephen", "Hawking", "blackhole"},
	// 	{"Marie", "Curie", "radium"},
	// 	{"Charles", "Darwin", "fittest"},
	// }

	// for i := range names {
	// 	n := names[i]
	// 	fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

	// 	if i == 0 {
	// 		fmt.Println(strings.Repeat("=", 50))
	// 	}
	// }
}
