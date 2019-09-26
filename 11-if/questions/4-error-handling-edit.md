# Questions: Error Handling
---

## 1. Why do we need error handling in our programs? ##
(A) I don't know. <br>
(B) To control the program's execution flow. <br>
(C) To make the program safe from dangerous malware. <br>
(D) Because many things can go with the program. <br>

---

**Solution**: (D) <br>
Many things can potentially go wrong when running a program, so catching an error can help you write code in the program to prepare for where something cause problems.

---

## 2. What is a nil value? ##
(A) The dark matter that rules the universe. <br>
(B) It's a zero-value for pointers or pointer-based types, meaning the value has not been initialized. <br>
(C) It's holds the same value as an empty strings (`""` == nil) <br>

---

**Solution**: (B) <br>

---

## 3. What is an error value? ##
(A) It stores the error details. <br>
(B) A global variable that stores the error status. <br>
(C) A global constant which stores the error status. <br>

---

**Solution**: (A) <br>
In Go, errors are given values, and therefore can be stored into a variable. 
<br>
There are no global variables in Go -- only package-level variables.

---

## 4. How does Go handle errors? ##
(A) Using a throw/catch statement <br>
(B) Using a simple if statement with nil comparison  <br>
(C) Using a mechanism called tagging <br>

---

**Solution**: (B) <br>
There are not throw/catch statements in Go. <br>
Unlike Java, C#, and other languages, Go is very explicit about error handling.

---

## 5. When should you handle errors? ##
(A) After the `main()` function. <br>
(B) Before calling any function. <br>
(C) Immediately after calling a function that returns an error value. <br>

---

**Solution**: (C) <br>
When calling the function, it is a good idea initialize and assign an explicit variable to the error value returned from a function and use and if statement against the nil value to handle the errors.

---

## 6. Which of the following functions should you consider handling errors for?  ##
```go
func Read() error {...}
func Write() error {...}
func String() string {...}
func Reset() {...}
```
(A) `Read()` and `Write()` <br>
(B) `String()` and `Reset()` <br>
(C) `Read()`, `Write()`, and `Reset()` <br>
(D) None of the functions need error handling <br>
(E) All of the functions need error handling <br>

---

**Solution**: (B) <br>
The return `Read()` and `Write()` functions are the only functions that have an return value that has and error type. <br>
`String()` and `Reset()` don't return any error values so checking for them may not be necessary. <br>

---

## 7. What does it mean when a function returns a `nil` error value? ##
(A) It means the function call failed. <br>
(B) It means the function call was successful. <br>
(C) It means the function call is currently in an indeterministic state. <br>

---

**Solution**: (B) <br>
A `nil` error value is a good thing meaning, that the code can continue executing while there is more code to execute.

---

## 8. What does it mean when a function returns a `non-nil` error value?
(A) It means the function call failed. <br>
(B) It means the function call was successful. <br>
(C) It means the function call is in an indeterministic state; we don't know.

---

**Solution**: (A) <br>
A simplified explanation is that the function call failed. <br>
Later on you'll learn that a function can return a `non-nil` value, and the returned value may indicate an error. <br>



> **1:** Yep. Later on you'll learn that, this is not always true. Sometimes a function can return a non-nil error value, and the returned value may indicate something rather than an error. <br>

Search on Google: golang io EOF error if you're curious.

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
}
if err != nil {
    fmt.Println(d)
}
```
(A) Yes; it prints the parsed duration if it's successful. <br>
(B) No; it doesn't check for errors. <br>
(C) No; it prints the duration even when there's an error.

---

**Solution**: (C) <br>
The function begins the error handling logic out well, but fails to use the `return` keyword which stops the execution of the program's `main()` function and exits the program.

---

## 10. Does the following program use error handling correctly? ##
```go
// The beginning of the `time.ParseDuration` function looks like this:
func ParseDuration(s string) (Duration, error) {...}
```
For further information, run `go doc time.ParseDuration` in the terminal do see the documentation.
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
(A) Yes; it prints the parsed duration if successful. <br>
(B) No; it doesn't check for the errors. <br>
(C) No; it prints the duration even when there's an error. <br>

---

**Solution**: (A) <br>
If there is an error, the error program will print the error message and then exit the program. <br>
If there is no error, the return value from `time.ParseDuration()` is printed to the console.

---
