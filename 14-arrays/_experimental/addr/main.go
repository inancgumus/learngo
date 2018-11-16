package main

import "fmt"

func main() {
	age1 := getAge(15)
	age2 := getAge(25)
	age3 := getAge(35)

	fmt.Printf("age1 lives in %p\n", age1)
	fmt.Printf("age2 lives in %p\n", age2)
	fmt.Printf("age3 lives in %p\n", age3)
}

func getAge(n byte) *byte {
	m := new(byte)
	*m = n
	return m
}
