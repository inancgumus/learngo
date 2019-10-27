# Questions: Code Your Own Program

---

## (1) Which of the following keywords is used to define a package?

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

1. `func`
2. `package`
3. `Println()`
4. `import`

---

**Solution: 2**

That's right. The keyword `package`, placed at the top of the program code, allows you to define a new Go package.

> 1: The keyword `func` is used to declare a new function.
> 3: `Println()` is a function from the `"fmt"` package.
> 4: The keyword `import` is used to import a package.

---

## (2) What is the meaning of `package main` in the following program?

```go
package main

func main() {
    // ...code goes here
}
```

1. It creates a library package.
2. It tells Go to exit the program.
3. It creates an executable Go program.

---

**Solution: 3**

That's right. When the Go compiler reads the `package main` declaration, it creates an executable program from the code within that file.

---

## (3) What does `func main()` do in the following program?

```go
package main

func main() {
    // ...
}
```

1. It defines a package named `main`.
2. It tells Go where to begin executing the program.
3. It prints the message written inside `{}` to the console.

---

**Solution: 2**

Correct! The `func main()` is a special function that tells Go where to start executing a program after it compiles successfully.

> 1: The `main()` function doesn't create a package.
> 3: It doesn't print an

---

## (4) What does `import "fmt"` do in the following program?

```go
package main
import "fmt"

func main() {
	fmt.Println("Hi!")
}
```

1. It prints `"fmt"` to the console.
2. It defines a new package called `"fmt"`
3. It imports the `"fmt"` package, allowing the file to use its functionalities.

---

**Solution: 3**

The package `"fmt"` is part of the standard Go library that conatins functions that allow Go programs to print output to the console.

---

## (5) Which keyword is used to declare a new function?

1. `func`
2. `package`
3. `Println()`
4. `import`

---

**Solution: 1**

The keyword `func` declares a new function in Go.

---

## (6) What is a function?

1. It's like a mini-program that is reusable and executable.
2. It tells Go to execute a program.
3. It allows Go to import a package called `function`.
4. It prints a message to the console.

---

**Solution: 1**

A function containing blocks of code can be re-used within the program to execute the same logic as many times as needed.

> 2: Go looks for `package main` and `func main()` to begin executing the program code.
> 3: The `import` keyword is reserved for importing packages into the program. 
> 4: functions from the `fmt` package can print text to the console.

---

## (7) Are you required to call the `main()` function yourself?

1. Yes, the program cannot execute without it.
2. No, Go automatically calls `main()`.

---

**Solution: 1**

That's right. As long as the `main()` function is defined, Go will automatically find and execute it.

---

## (6) Except for the `main()` function, do you have to call a function to execute it?

1. Yes; otherwise go will not execute the code.
2. Yes; Go will execute any program in the program.
3. No; Go calls all functions automatically.

---

**Solution: 1**

That's right. You need to call all functions besides the `main()` function if you want Go to execute them.

---

## (7) What will the following program do when executed?

```go
package main

func main() {}
```

1. It prints a message to the console.
2. It's a won't print anything to the console.
3. The program won't run because there is an error.

---

**Solution: 2**

It won't print anything because the `main()` function doesn't contain the `fmt.Println()` function.

> 1: It doesn't print a message to the console. To do that, you can use the function `fmt.Println()`.
> 3: There program doesn't contain any errors because uses the `package` keyword and has a `main()` function.

---

## (8) What does the following program print?

```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}
```

1. `"Hi, I want to be a Gopher!"`
2. The program won't print anything to the console.
3. The program won't compile successfully.

---

**Solution: 3**

That's right! The program requires a few more things before it can print something to the console.

The value inside `fmt.Println()` is not a valid string literal and should be wrapped in double-quotes like this: `"Hi! I want to be a Gopher!"`

The `import` statement is also missing, which is needed to used `Println()` from the package `"fmt"`.

---

## (9) What does the following program print? ##

```go
package main
import "fmt"

func main() {
	fmt.Println("Hi there!")
}
```

1. `Hi there!` 
2. `fmt` 
3. This program contains an error; `Println()` is not a part of the `fmt` package.

---

**Solution: 1**

> 2: `import "fmt"` imports the `fmt` package, so you can use its functionalities.
> 3: The program does not contain any errors.
