package main

import "fmt"

func main() {
	type book struct {
		name, isbn string
	}

	kafka := book{"Kafka's Revenge", "S-001"}
	golden := book{"Stay Golden", "S-002"}

	books := make(map[book]int, 2)
	books[kafka] = 100
	books[golden] = 50

	fmt.Printf("%s sold %d times\n", kafka.name, books[kafka])
	fmt.Printf("%s sold %d times\n", golden.name, books[golden])
}
