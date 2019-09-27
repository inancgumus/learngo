# Questions: Code Your Own Program

---

## 1. Which of the following keywords is used to define a package?

```go
package main
import "fmt"
func main() {
    fmt.Println("Hello world!")
}
```

(A) `func`

(B) `package`

(C) `Println()`

(D) `import`

---

**Solution**: (B)

The keyword `package`, placed at the top of the program code, allows you to define a new 
Go package.

(A) The keyword `func` is used to declare a new function.

(C) `Println()` is a function function from the "fmt" package.

(D) The keyword `import` is used to import a package

---
## 2. What is the meaning of `package main` in the following program?
```go
package main

func main() {
    // ...code goes here
}
```

(A) It creates a library package.

(B) It tells Go to exit the program.

(C) It creates an executable Go program.

---

**Solution**: (C)
When the Go compiler reads the `package main` declaration, it creates an executable program from the code within that file.

---

## 3. What does `func main()` do in the following program?
```go
package main

func main() {
    // ...
}
```

(A) It tells Go the name of the program is.

(B) It tells Go where to begin executing the program.

(C) It prints messages written inside `{}` to the console.

---

**Solution**: (B)

The `func main()` is a special function that tells Go where to start executing a program after it compiles successfully.

---

## 2. What does `import "fmt"` do in the following program?

(A) It prints "fmt" to the console.

(B) It defines creates a package called "fmt".

(C) It imports the `fmt`, allowing the file to use its functionalities.

---

**Solution**: (A)
The package `fmt` is part of the standard Go library that contains functions that allow 
Go programs to print output to the console in

---

## 3. Which keyword is used to declare a new function?

(A) `func`

(B) `package`

(C) `Println()`

(D) `import`

---

**Solution**: (A)
The keyword `func` tells declares a new function in Go.

---

## 4. What is a function?

(A) It's like a mini-program that is reusable and executable.

(B) It tells Go to import.

(C) It tells Go to import a package called `function`.

(D) It prints a message to the console.

---

**Solution**: (A)
A function containing blocks of code can be re-used within the the program to execute 
the same logic as many times as needed.

(B) Go looks for `package main` and `func main()` to begin executing the program code.

(C) The `import` keyword is reserved for importing packages into the program. 
(D) functions from the `fmt` package can print text to the console.

---
## 5. Are you required to call the `main()` function yourself?

(A) Yes, the program cannot execute without it.

(B) No, Go automatically calls `main()`.

---

**Solution**: (A)

As long as the `main()` function is defined, Go will automatically find and execute it.

---
## 6. Except for the `main()` function, do you have to call a function to execute it?

(A) Yes; otherwise go will not execute the code.

(B) Yes; go will execute any program in the program.

(C) No; Go call all functions automatically.

---

**Solution**: (A)
All functions must be declared before Go can execute it.

---
## 7. What will the following program do when executed?
```go
package main

func main() {}
```

(A) It prints a message to the console.

(B) It's a won't print anything to the console.

(C) The program won't run because the compiler will throw an error.

---

**Solution**: (B)

(A) The function doesn't call the `fmt.Println()` function.

(C) The program will compile successfully but won't print anything to the console because there are no declarations or statements within the `main()` function to execute.

---

## 8. What does the following program print?

```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}
```

(A) "Hi, I want to be a Gopher!"

(B) The program won't print anything to the console.

(C) The program won't compile successfully.

---

**Solution**: (C)

That's right! The program requires a few more things before it can print something to the console.

The value inside `fmt.Println()` is not a valid string literal and should be wrapped in double-quotes like this: `"Hi! I want to be a Gopher!"`

The `import` statement is also missing, which is needed to used `Println()` from the package `"fmt"`.