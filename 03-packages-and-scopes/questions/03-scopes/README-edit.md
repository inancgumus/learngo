<!-- [jarrodchung-dev]: Edits -->
<!-- 2019-09-24 -->

# Questions: Code Your Own Program

## 1. Which of the following keywords is used to define a package?
```go
package main
import "fmt"
func main() {
    fmt.Println("Hello world!")
}
```
(A) `func` <br>
(B) `package` <br>
(C) `Println()` <br>
(D) `import` <br>

---
**Solution**: (B) <br>
The keyword `package`, placed at the top of the program code, allows you to define a new Go package. <br>
(A) The keyword `func` is used to declare a new function. <br>
(C) `Println()` is a function function from the "fmt" package. <br>
(D) The keyword `import` is used to import a package <br>

---
## 2. What is meaning of `package main` in the following program?
```go
package main

func main() {
    // ...code goes here
}
```
(A) It creates a library package. <br>
(B) It tells Go to exit the program. <br>
(C) It creates an executable Go program. <br>

---
**Solution**: (C) <br>
When the Go compiler reads the `package main` declaration, it creates an executable program from the code within that file.

---
## 3. What does `func main()` do in the following program?
```go
package main

func main() {...}
```
(A) It tells Go the name of the program is. <br>
(B) It tells Go where to begin executing the program. <br>
(C) It prints a messages written inside `{}` to the console. <br>

---
**Solution**: (B) <br>

The `func main()` is a special function that tells Go where to start executing a program after it compiles successfully.

---
## 2. What does `import "fmt"` do in the following program?
(A) It prints "fmt" to the console. <br>
(B) It defines creates package called "fmt". <br>
(C) It imports the `fmt`, allowing the file to use its functionalities. <br>

---
**Solution**: (A) <br>
The package `fmt` is part of the standard Go library that contains functions that allow Go programs to print output to the console in various formats.

---
## 3. Which keyword is used to declare a new function?
(A) `func` <br>
(B) `package` <br>
(C) `Println()` <br>
(D) `import` <br>

---
**Solution**: (A) <br>
The keyword `func` tells declares a new function in Go.

---

## 4. What is a function?
(A) It's like a mini-program that is reusable and executable. <br>
(B) It tells Go to import. <br>
(C) It tells Go to import a package called `function`. <br>
(D) It prints a message to the console. <br>

---
**Solution**: (A) <br>
A function containing blocks of code can be re-used within the the program to execute the same logic as many times as needed. <br>
(B) Go looks for `package main` and `func main()` to begin executing the program code. <br>
(C) The `import` keyword is reserved for importing packages into the program. 
<br>
(D) functions from the `fmt` package can print text to the console.

---
## 5. Are you required to call the `main()` function yourself?
(A) Yes, the program cannot execute without it. <br>
(B) No, Go automatically calls `main()`. <br>

---
**Solution**: (A) <br>
As long as the `main()` function is defined, Go will automatically find and execute it. <br>

---
## 6. Except the `main()`, you have to call a function to execute it?
(A) Yes; otherwise go will not execute the code. <br>
(B) Yes; go will execute any program in the program. <br>
(C) No; Go call all functions automatically. <br>

---
**Solution**: (A) <br>
All functions must be declared before Go can execute it.

---
## 7. What will the following program do when executed?
```go
package main

func main() {}
```
(A) It prints a message to the console. <br>
(B) It's a won't print anything to the console. <br>
(C) The program won't run because the compiler will throw an error. <br>

---
**Solution**: (B) <br>
(A) The function function doesn't call the `fmt.Println()` function. <br>
(C) The program will compile successfully but won't print anything to the console because there are no declarations or statements within the `main()` function to execute.

---
## 8. What does the following program print?

```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}
```
(A) "Hi, I want to be a Gopher!" <br>
(B) The program won't print anything to the console. <br>
(C) The program won't compile successfully. <br>

---
**Solution**: (C) <br>
(A) The values inside aren't wrapped in between two double-quotes; In order to be valid syntax, the string the value must be "Hi, I want to be a Gopher!" <br>
(C) The program is also missing an import statement that would allow the program to use the functions from the `fmt` package. 