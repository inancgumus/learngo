// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("••••• ARRAYS")
	arrays()

	fmt.Println("\n••••• SLICES")
	slices()

	fmt.Println("\n••••• MAPS")
	maps()

	fmt.Println("\n••••• STRUCTS")
	structs()
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(myHouse)

	// fmt.Printf("%+v\n", myHouse)
	fmt.Printf("structs()     : %p %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)
	fmt.Printf("structs()     : %p %+v\n", &myHouse, myHouse)
}

func addRoomPtr(h *house) {
	h.rooms++ // same: (*h).rooms++
	fmt.Printf("addRoomPtr()  : %p %+v\n", h, h)
	fmt.Printf("&h.name       : %p\n", &h.name)
	fmt.Printf("&h.rooms      : %p\n", &h.rooms)
}

func addRoom(h house) {
	h.rooms++
	fmt.Printf("addRoom()     : %p %+v\n", &h, h)
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)

	// &confused["one"]
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

func slices() {
	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	fmt.Printf("slices list   : %p %q\n", &dirs, dirs)

	upPtr(&dirs)
	fmt.Printf("slices list   : %p %q\n", &dirs, dirs)
}

func upPtr(list *[]string) {
	lv := *list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "HEISEN BUG")

	fmt.Printf("upPtr list    : %p %q\n", list, list)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
		fmt.Printf("up.list[%d]    : %p\n", i, &list[i])
	}

	list = append(list, "HEISEN BUG")

	fmt.Printf("up list       : %p %q\n", &list, list)
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("arrays nums   : %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
}

func incr(nums [3]int) {
	fmt.Printf("incr nums     : %p\n", &nums)
	for i := range nums {
		nums[i]++
		fmt.Printf("incr.nums[%d]  : %p\n", i, &nums[i])
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr nums: %p\n", &nums)
	for i := range nums {
		nums[i]++ // same: (*nums)[i]++
	}
}
