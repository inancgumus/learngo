## Which keyword below that you need use to define a package?
```go
package main

func main() {
}
```
1. func
2. package *CORRECT*
3. fmt.Println
4. import

> **1:** func keyword is used to declare a new function.
>
>
> **2:** That's right! package keyword allows you to define a package for a Go file.
>
>
> **3:** It's not a keyword, it's a function of the fmt package.
>
>
> **4:** import keyword is used to import a package.
>
>


## What is the purpose of using package main in the following program?
```go
package main

func main() {
}
```
* To create a library package
* To properly exit from the program
* To create an executable Go program *CORRECT*


## What is the purpose of func main in the following program?
```go
package main

func main() {
}
```
1. It defines a package called main
2. It allows Go to start executing the program *CORRECT*
3. It prints a message to the console

> **1:** main function doesn't create a package.
>
>
> **2:** That's right. Go automatically calls the main function to execute a program.
>
>
> **3:** It doesn't print anything (at least directly).
>
>


## What is the purpose of import "fmt" in the following program?
```go
package main
import "fmt"

func main() {
    fmt.Println("Hi!")
}
```
1. It prints "fmt" to the console
2. It defines a new package called "fmt"
3. It imports the `fmt` package; so you can use its functionalities *CORRECT*

> **1:** `fmt.Println` prints a message not the `import "fmt"`.
>
>
> **2:** `package` keyword does that, not the `import` keyword.
>
>
> **3:** Yes. For example, after you import the fmt package you can call its Println function to print a message to the console.
>
>


## Which keyword is used to declare a new function?
* func *CORRECT*
* package
* Println
* import


## What is a function?
1. It's like a mini-program. It's a reusable and executable block of code. *CORRECT*
2. It allows Go to execute a program.
3. It allows Go to import a package called function.
4. It prints a message to the console.

> **2:** Go looks for package main and func main to do that. A function doesn't do that on its own.
>
>
> **3:** `import` keyword does that.
>
>
> **4:** For example: `fmt.Println` does that.
>
>


## Do you have to call the main function yourself?
1. Yes, so that, I can execute my program.
2. No, Go calls the main function automatically. *CORRECT*

> **1:** No, you don't need to call the main function. Go automatically executes it.
>
>


## Do you have to call a function to execute it?
_(except the main func)_
1. Yes, so that, Go can execute that function. *CORRECT*
2. Yes, so that, Go can execute my program.
3. No, Go calls the functions automatically.

> **1:** That's right. You need to call a function yourself. Go won't execute it automatically. Go only calls the main function automatically (and some other functions which you didn't learn about yet).
>
>

> **2:** That's only the job of the `func main`. There's only one `func main`.
>
>

> **3:** Go doesn't call any function automatically except the main func (and some other functions which you didn't learn about yet). So, except the main func, you need to call the functions yourself.
>


## What does the following program print?
```go
package main

func main() {
}
```
1. It prints a message to the console
2. It's a correct program but it doesn't print anything *CORRECT*
3. It's an incorrect program

> **1:** It doesn't print a message. To do that you can use fmt.Println function.
>
>

> **2:** Yes, it's a correct program, however since it doesn't contain fmt.Println it doesn't print anything.
>
>

> **3:** It's a correct program. It uses the package keyword and it has a main function. So, this is a valid and an executable Go program.
>
>


## What does this program print?
```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}
```
* Hi! I want to be a Gopher!
* It doesn't print anything
* This program is incorrect *CORRECT*

> **1:** It doesn't pass the message to Println wrapped between double-quotes. It should be like: fmt.Println("Hi! I want to be a Gopher")
>
>

> **3:** It doesn't import "fmt" package. Also see #1.
>
>


## What does this program print?
```go
package main
import "fmt"

func main() {
    fmt.Println("Hi there!")
}
```
* Hi there! *CORRECT*
* fmt
* This program is incorrect; it imports the wrong package or there isn't a function called `Println`

> **2:** import "fmt" imports the `fmt` package; so you can use its functionalities.
>
>
> **3:** Actually, this program is correct.
>
>