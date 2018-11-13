package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [...][3]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record[:]); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	// Output:
	// first_name,last_name,username
	// Rob,Pike,rob
	// Ken,Thompson,ken
	// Robert,Griesemer,gri
}
