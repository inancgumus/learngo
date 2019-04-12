// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
)

// #4: BTW, I can also use a package variable.
// Here, this variable belongs to the main package.
// So, any function inside the main package can see it.
//
// REF: http://wiki.c2.com/?GlobalVariablesAreBad
var even = 16

func main() {
	// #1C: Let me declare a variable in the main function.
	number := 1
	// Now, I'm going to call the function with the variable.
	incr(number)
	fmt.Println("main.number      :", number)
	// As you can see, it prints 1 instead of 2. Why?
	// It's because, Go is a pass-by-value language. So, when you pass an argument to a function, it gets copied.

	// #1E: For example, I cannot print it here. It isn't available here because it's only in the scope of the incr function.
	// fmt.Println("Inside main:", num)

	// #2: The name of the incr function's input value is num.
	// So, what would happen if I would declare a variable with the same name in the main function, and pass it to the incr function like so?
	num := 10
	incr(num)
	fmt.Println("main.num         :", num)

	// As you can see, it doesn't matter whether the variable has the same name or not. There are two num variables with the same name but they are in different scopes. Inside the incr function, the num variable is a different variable even though the main function also has a variable with the same name.

	// #3: So, how can you change the number in a function?
	// Well, you can return a value from the function.
	// This time, I'm using a function that decrements the given number.
	num = 5
	num = decr(num)
	fmt.Println("main.num         :", num)
	// #3C: As you can see, this time, it has decremented the number.
	// Here, I've simply got a copy of the number from the decr function.
	// So, it's still not the decr function's local variable.

	// #4C: Let me call the new function here like so.
	square()
	fmt.Println("even             :", even)
	// As you can see, this time the main function sees the change.
	// It's because, as I've said, the even variable belongs to the main package, any code inside the main package can access it. However, doing so is not good because anyone can change it.

	// #4E: Now, check out the result, it is 0! This is why pass-by-value is a good thing because it isolates the local variables of a function from another one. So the other functions cannot change the same data.
}

// #1: Let's say I want to increase the given number inside a function.
// So, I'm going to declare a function to do so.
// It accepts an int value, and it doesn't return anything.
func incr(num int) {
	// #1D: Here, it receives the number value and then it declares a hidden variable behind the scenes like so.
	// var num int
	//
	// This variable is local only to the incr function. So the other functions cannot see it from the outside of this function.

	// #1F: The main function, passes 1 to this function. So, after declaring the hidden variable here, Go sets it to the given argument like so.
	// num = 1
	//
	// And then, this statement increments the local copy of the variable.

	// #1B: Here, I'm going to increase the number.
	num++

	// #1G: So, if I print it from here, it's going to say: 2.
	fmt.Println("incr.num         :", num)
	// It prints 2 because the variable is a copy of the number variable.
}

// #3B: Let me also declare the function.
// It gets an int as an input value.
// And it decrements and returns the value.
func decr(num int) int {
	// #4D: Let's say, for some weird reason, the decr function secretly changes the even variable. In a large code base this can easily go unnoticed.
	even = 0
	return num - 1
}

// #4B: Let me declare a function that squares this number.
// This time, this function doesn't accept or return any values because it operates on the package level variable.
func square() {
	even *= even
}
