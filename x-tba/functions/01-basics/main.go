// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// #1: Let's say you want to print a message three times like so.
	// fmt.Println("hello gopher!")
	// fmt.Println("hello gopher!")
	// fmt.Println("hello gopher!")

	// It works. However, when you want to do the same thing over and over again, it's better to use a function. Functions allow us to reuse the same block of code over and over again.

	// For example, instead, let's say there is a function that prints the same message. Let's call it three times.
	greet()
	greet()
	greet()

	// #4: As you can see, it prints the same message three times because I've called the function three times. The benefit of using a function here is that I don't have to duplicate the code inside the function, and I can change it in one place like so (see #4B).

	// #5: Now, let's say I want to print different messages like so.
	greetWith("i'm bond. james bond.")
	greetWith("i'm gopher. lovely gopher.")

	// greetWith()
	// #6B: BTW, if a function expects a parameter, I cannot call it without an argument like so. A parameter is a definition of what a function expects. An argument is the actual data that you send to a function. Here, the string value is an argument, it's a real value. But here, it's just a parameter declaration.

	// #7: Now it prints different messages. You can pass and call this function with any string value. However, as you know, you cannot call it with other types like so.
	// greetWith(42)
	// Doing so is an error because the function only accepts a string.

	// #8: Now, let's say I want to negate an integer value. For example, when I give it 100, it's going to give me -100.
	n := negate(100)
	fmt.Printf("negate(%d) = %d\n", 100, n)
	// #10: As you can see, the negate function returns -100.

	// #11: I can also pass multiple inputs to a function like so.
	// Here, this function is going to multiply the given numbers and it will return a float64 number.
	f := multiply(10, 20)
	fmt.Printf("%d * %d     = %g\n", 10, 20, f)

	// #13: Now, let's take a look at a function that returns multiple values.
	// I'm going to give it two numbers, and it's going to divide them, it's also going to return back an error if it cannot divide the numbers like so.
	n, err := divide(10, 5)

	// So, I'm going to the handle error as usual.
	if err != nil {
		fmt.Println(err)

		// As you can see, here, this time, I'm returning from the main function. The main function is running the whole show. It orchestrates the other functions. So, when I return from the main function your program will terminate, or in other words, it will quit.
		//
		// BTW note that, here, I don't have to provide any arguments to the return statement because the main function doesn't return anything. As you can see, it doesn't declare any result parameters.
		return
	}

	// If there isn't an error, I'm going to print the result like so.
	fmt.Printf("%d / %d      = %d\n", 20, 0, n)

	// #15: As you can see, the divide function prints an error, and the program terminates because the main function returns. So, it doesn't execute the rest of the statements.
	//
	// Let me change the arguments so that the divide function can divide the numbers without any error.
	// divide(10, 5)
	//
	// As you can see, this time, it prints the result of the division.

	// #16B: Let me call it from here, with an empty string argument.
	greetWith("")
}

// #2: Of course, the greet function doesn't exist yet, so let's declare it.
// It looks like the main function, however, its name is different, and it doesn't get called automatically, you need to call it by yourself. That's why I need to call it here [above].
//
// Inside the function, I'm going to print the hello message like so.
func greet() {
	fmt.Println("hello gopher!")

	// #4B: Now, it prints a different message. However, nothing interesting here, so let's continue :)
	// fmt.Println("hello there!")
}

// #3: BTW, every declared function occupies the same name space of its package. I mean, you cannot redeclare a function with the same name of an existing function. So, there can be only one function within the same package with the same name. Also note that, Go doesn't support function overloading if you're wondering about it.
// func greet() {
// 	fmt.Println(hello)
// }

// #6: This time, I'm going to name the function greetWith. Here, inside the parentheses, I'm going to type the input parameter that the function is expecting. It expects only a single string parameter: message.
//
// Inside the function body, I'm going to print the given message like so.
func greetWith(message string) {
	// #16: BTW, I don't have to return from this function as well because it doesn't declare a result value.
	// However, if I want to, I can use the return statement in this function as well.
	// Let's say I want to return from this function when it's called with an empty argument like so.
	if message == "" {
		// If so, I'm going to print the default message then I'm going to terminate the function.
		fmt.Println("i can greet!")
		return
	}
	fmt.Println(message)

	// #16C: As you can see, when I call it without any arguments, it prints the default message instead.
}

// #9: To do that, I'm going to return an int value so the main function can print it.
// It's better for a function to have a single responsibility. Here, it only needs to negate the number, it's not responsible from printing it. So the caller of this function can use the returned data however it wants.
func negate(num int) int {
	// The return statement returns a value to the calling function. Here, I'm going to return the number by negating it like so.
	//
	// As you can see, I can use an expression next to a return statement. Here, it returns the result of this expression.
	return -num

	// BTW, here, I can only return an int value because the function declares that it returns an int. For example, I cannot return a string value.
	// return "nope"
}

// #12: So here, I'm going to declare the multiply function, and I'm going to declare it with two parameters. And lastly, I'm going to return a float64 value.
// As you can see, I can use the serial declaration, I could also have declared it like so:
// multiply(a float64, b float64)
// It's the same thing.
func multiply(a, b float64) float64 {
	// Now, I'm going to multiply the given arguments and I'm going to return the result of the expression like so.
	return float64(a) * b
}

// #14: OK, now let's declare the divide function.
// It accepts two int parameters, and returns two values like so.
//
// As you can see, when you return multiple values, you need to separate them with a comma, and you need to put them inside the parentheses. You can return more than two result values, but usually, we return two result values for making the function easy to work with.
func divide(a, b int) (int, error) {
	// Inside the function, I'm going to check for the division by zero error. To do that I'm going to check whether the b is zero or not.
	if b == 0 {
		// If so, I'm going to return two values from the function.
		// First, I'm going to return 0 because it's the zero value of an int type.
		// I could have returned any other value as well, but idiomatically, it's better to return a zero value.
		//
		// Then, I'm going to return an error value using the Errorf function like so. This function is just like the Printf function, the only difference is that, instead of printing the result to the console, it returns an error value with the given message.
		return 0, fmt.Errorf("error: divide by zero: %d/%d", a, b)
	}

	// OK, if the b argument is not zero, the execution will continue from here. So here, I'm going to divide the numbers, and as the second result value, I'm going to return nil because that is the zero value for the error type because the error type is an interface.
	return a / b, nil
}
