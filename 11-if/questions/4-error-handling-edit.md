# Questions: Error Handling

---

## 1. Why do we need error handling in our programs? ##

(A) I don't know.

(B) To control the program's execution flow.

(C) To make the program safe from dangerous malware.

(D) Because many things can go wrong with the program.

---

**Solution**: (D)

Many things can potentially go wrong in a running program, so catching errors help control the problems and properly handle the errors.

---

## 2. What is a nil value? ##

(A) The dark matter that rules the universe.

(B) It's a zero-value for pointers or pointer-based types, meaning the value has not been initialized.

(C) It is equal to an empty string (`"" == nil`)

---

**Solution**: (B)

---

## 3. What is an error value? ##

(A) It stores the error details.

(B) A global variable that stores the error status.

(C) A global constant that stores the error status.

---

**Solution**: (A)

In Go, errors are values, and therefore can be stored in a variable.

There are no global variables in Go -- only package-level variables.

---

## 4. How does Go handle errors? ##

(A) Using a throw/catch statement

(B) Using a simple if statement with nil comparison 

(C) Using a mechanism called tagging

---

**Solution**: (B)

There are no throw/catch statements in Go.

Unlike Java, C#, and other languages, Go is very explicit about error handling.

---

## 5. When should you handle errors? ##\ ##

(A) After the `main()` function.

(B) Before calling any function.

(C) Immediately after calling a function that returns an error value.

---

**Solution**: (C)

When calling a function, it's a good idea to initialize and assign an explicit variable to the error value returned from a function, and use an `if` statement against the `nil` value to handle the errors.

---

## 6. Which of the following functions should you consider handling errors for?  ##

```go
func Read() error {...}
func Write() error {...}
func String() string {...}
func Reset() {...}
```

(A) `Read()` and `Write()`

(B) `String()` and `Reset()`

(C) `Read()`, `Write()`, and `Reset()`

(D) None of the functions need error handling

(E) All of the functions need error handling

---

**Solution**: (B)

`Read()` and `Write()` functions are the only functions that return error values. Therefore, you might consider handling the errors they are called.

`String()` and `Reset()` don't return any error values so error checking is not necessary.

---

## 7. What does it mean when a function returns a `nil` error value? ##

(A) It means the function call failed.

(B) It means the function call was successful.

(C) It means the function call is currently in an indeterministic state.

---

**Solution**: (B)

A `nil` error value is a good thing, meaning that the code can continue executing while there is more code to execute.

---

## 8. What does it mean when a function returns a `non-nil` error value? ##

(A) It means the function call failed.

(B) It means the function call was successful.

(C) It means the function call is in an indeterministic state; we don't know.

---

**Solution**: (A)

A simplified explanation is that the function call failed.

Later on you'll learn that a function can return a `non-nil` value, and the returned value may indicate an error.

To learn more, Google this: `golang io EOF error`.

---

## 9. Does the following program use error handling correctly? ##

```go
func ParseDuration(s string) (Duration, error) {...}
```

**Note:** This is what the `time.ParseDuration()` function looks like.
To view the documentation for the function, run `go doc time.ParseDuration` at the command line.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h10s")
	if err != nil {
		// ...
	}
	fmt.Println(d)
}
```

(A) Yes; it prints the parsed duration if it's successful.

(B) No; it doesn't check for errors.

(C) No; it prints the duration even when there's an error.

---

**Solution**: (C)

The function begins the error handling logic out well, but fails to use the `return` keyword which stops the execution of the program's `main()` function and exits the program.

---

## 10. Does the following program use error handling correctly? ##

```go
// The beginning of the `time.ParseDuration` function looks like this:
func ParseDuration(s string) (Duration, error) {...}
```
**Note:** This is what the `time.ParseDuration()` function looks like.
To view the documentation for the function, run `go doc time.ParseDuration` at the command line.

```go
package main 

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h10s")
	if err != nil {
		fmt.Println("Parsing error:", err)
		return 
	}
	fmt.Println(d)
}
```

(A) Yes; it prints the parsed duration if successful.

(B) No; it doesn't check for the errors.

(C) No; it prints the duration even when there's an error.

---

**Solution**: (A)

If there is an error, the error program will print the error message and then exit the program.

If there is no error, the return value from `time.ParseDuration()` is printed to the console.

---
