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

func main() {
	// #1: Go is a pass-by-value programming language. So, when you assign a variable to another one, or pass it to a function, it gets copied. Let me show you how it work in the context of the functions.

	// #1C: Let me declare a variable.
	number := 1
	// Now, I'm going to call the function with the variable.
	incr(number)
	// Let me print the variable here.
	fmt.Println("main.number      :", number)
	// As you can see, it prints 1 instead of 2. Why?
	// It's because, as I've said, Go is a pass-by-value programming language.

	// #1E: For example, I cannot print it here. It isn't available here because the num is local to the incr function. It's in the scope of the incr function. If you don't remember the scopes, please check out the earlier scope lectures.
	// fmt.Println("Inside main:", num)

	// #2: So what if you want to change the value in the function?
	// In that case, you need to use a pointer because a pointer can send the memory location of the variable to the function, so the function can change the value directly on the memory.
	//
	// For example, let me pass the address of the number to another function like so. Ptr is short for pointer, we use it a lot.
	//
	// Ampersand sign here returns the memory location of the given variable.
	incrByPtr(&number)

	// Let me print the variable.
	fmt.Println("&main.number     :", number)

	// #2C: Let's check it out.
	// As you can see, now the number has changed to 2.
	// The function could increment the number because it can directly access to the value of the variable by using its memory address.

	// If you look at the memory address of the number variable, you will see that the function sees the same address, let me show you.
	//
	// Let me print the address of the number variable like so.
	// Every pointer stores a memory location value.
	// So the %p verb prints the address that a pointer stores.
	fmt.Printf("&main.number     : %p\n", &number)
}

// #1B:
// Let's say I want to increase the given number inside the function.
//
// So, I'm going to declare a function named Incr to do so.
// It accepts an int value, and it doesn't return anything.
func incr(num int) {
	// #1D: Here, it receives the num and then it declares a variable behind the scenes like so.
	// var num int
	// This variable is local to the incr function. The other functions cannot see this variable from the outside of this function.

	// #1F: The main function, passes 1 to this function. So, after declaring the hidden variable here, Go sets it to the given argument like so.
	// num = 1
	//
	// And then, this statement works and it increments the num variable. However, it increases its own local copy of the num variable.

	// #1B: Inside the function, I'm going to try to increase the given number like so.
	num++

	// #1G: So, if I print it from here, it's going to say: 2.
	fmt.Println("incr.num         :", num)
	// As you can see, it prints 2 because it gets copied here.
}

// #2B: Let me also declare the function like so.
// This time, it accepts an int pointer.
// When you put an asterisk sign in front of a type, it becomes a pointer.
// Every type comes with its own pointer type like so.
// This is an int pointer because it has an asterisk in front of the int type. So, this pointer can only store the memory address of an int value.
func incrByPtr(num *int) {
	// Inside the function I'm going to change the num variable by deferencing it like so, then I'm going to increase it as usual.
	*num++

	// #2D: Let me also print the memory address that the num pointer stores.
	// Remember: num here is a pointer variable, it stores a memory address of an int value.
	fmt.Printf("incrByPtr.number : %p\n", num)

	// As you can see, the main function and the incrByPtr function use the same variable. They both are looking at the same memory location. That's why this function can change the value.

	// However, if I print the memory address of the local num variable, it will print a different address, let me show you.
	fmt.Printf("&incrByPtr.number: %p\n", &num)

	// As you can see, the memory address of the local num variable is different than the memory address of this variable. It's because, Go is a pass-by-value language. So, it even copies the pointer variables. So, the local num pointer variable is a new variable inside this function. However, it stores the same address of the number variable of main function.

	// Don't worry, I'm going to talk about pointers in detail soon. For now, just know that, in Go, everything is passed by value, even pointers, however, when the functions knows the memory address of a variable, it can change it directly. Otherwise it cannot change it because the variable belongs to the local scope of the function.

	// Pass by value is a very good mechanism because it isolates the variables of a function from another one.
}
