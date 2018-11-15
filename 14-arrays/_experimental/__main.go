// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

const (
	Mon Weekday = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

var names = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}


for d := Mon; d <= Sun; d++ {
	fmt.Println(names[d])
}