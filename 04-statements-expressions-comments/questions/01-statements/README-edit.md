# Questions: Statements #
---
## 1. Which of the following is an best describes a statement? ##
(A) A statement instructs Go to do something. <br>
(B) A statement produces a value. <br>
(C) A statement can't change the execution flow of a program. <br>

---
**Solution**: (A) <br>
A statement, like `fmt.Println("hi")` tells go to print something to the console. <br>
Statements like `if` can defnintely change the execution flow of a program. 
<br>

---
## 2. In what direction does Go execute code? ##
(A) From left to right. <br>
(B) From top to bottom. <br>
(C) From right to left. <br>
(D) From bottom to top. <br>

---
**Solution**: (B) <br>
Go executes code from top to bottom, executing one statement at a time. <br>

---
## 3. Which of the following best describes an expression? ##
(A) An expression instructs Go to do something. <br>
(B) An expression produces a value. <br>
(C) An expression can change the execution flow of a program. <br>

---
**Solution**: (B) <br>
Remember that only statements can instruct Go to do something and change the execution flow of a program.

---
## 4. Which one of the following best describes an operator? ##
(A) An expression instructs Go to do something. <br>
(B) An expression can change the execution flow of a program. <br>
(C) An operator can combine expressions. <br>

---
**Solution**: (C)
If you have to strings variables with the operator `+` between them, the resulting value is a single string literal.

## 5. Review the program and determine why the program will not run. ##
```go
package main
import "fmt"

func main() {
    "Hello"
}
```
(A) Because "Hello" is an expression and cannot be on exist on a single line of code without a statement. <br>
(B) Because the double-quotes surrounding "Hello" must be removed. <br>
(C) Because "Hello" must be removed out of the `main()` function. <br>

---
**Solution**: (A) <br>
"Hello" is just a string literal value and without being used within a statement, the Go compiler will throw an error.

---
## 6. Review the program and determine if the program will work.  ##
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(
        runtime.NumCPU();
        fmt.Println("cpus");
        fmt.Println("on the machine");
    )
}
```
(A) It will run: Expressions can be type by separating them using semicolons. 
<br>
(B) It doesn't work; more than one statement on a line of code will throw an error.
<br>
(C) It will run, because Go automatically adds a semiocolons at the end of every statement behind the scenes. <br>

---
**Solution**: (C) <br>
Wheter a semicolon is explicity used or not, Go adds them automatically. To the compiler, each statement is on it's own line.

---
## 7. Review the program and select the choice that best explains why the program will run. ##
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU() + 10)
}
```
(A) Operators can combine expressions. <br>
(B) Statements can be used with oeprators. <br>
(C) Expressions can return multi <br>

---
**Solution**: (A)
The `+` operator combines `10` and `runtime.NumCPU()`. <br>
Statements can't be declared with operators. For example, you cannot combine `"fmt" + 3. <br>
Note: some statements scan be declared with expressions.

---
