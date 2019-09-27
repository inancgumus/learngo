# Questions: Statements #

---

## 1. Which of the following best describes a statement? ##

(A) A statement instructs Go to do something.

(B) A statement produces a value.

(C) A statement can't change the execution flow of a program.

---

**Solution**: (A) 

That's right. A statement like `fmt.Println("hi")` tells go to print something to the console.

Statements like `if` may change the execution flow of a program.

---

## 2. In what direction does Go execute code? ##

(A) From left to right.

(B) From top to bottom.

(C) From right to left.

(D) From bottom to top.

---

**Solution**: (B)

That's right. Go executes code from top to bottom, executing one statement at a time.

---

## 3. Which of the following best describes an expression? ##

(A) An expression instructs Go to do something.

(B) An expression produces a value.

(C) An expression can change the execution flow of a program.

---

**Solution**: (B)

Remember that only statements can instruct Go to do something and change the execution flow of a program.

---

## 4. Which one of the following best describes an operator? ##

(A) An operator instructs Go to do something.

(B) An operator can change the execution flow of a program.

(C) An operator can combine expressions.

---

**Solution**: (C)

If you combine string values with the operator `+` between them, the resulting value is a string literal.

Only statements can instruct tell Go to do something and change the execution flow of a program.

---

## 5. Please review the following program and determine why it will not run. ##

```go
package main
import "fmt"

func main() {
    "Hello"
}

```

(A) Because "Hello" is an expression and cannot exist on a single line of code without a statement.

(B) Because the double-quotes surrounding "Hello" must be removed.

(C) Because "Hello" must be removed out of the `main()` function.

---

**Solution**: (A)

`"Hello"` is just a string literal value and without being used within a statement, the Go compiler will produce an error.

---

## 6. Please review the following program and determine if the program will run. ##

```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU());
    fmt.Println("cpus");
    fmt.Println("on the machine");
}
```

(A) It will run: You can type expressions by separating them using semicolons.

(B) It won't work: You can only use a single statement on a single line of code.

(C) It will run: Go automatically adds a semicolon at the end of every statement behind the scenes.

---

**Solution**: (C)

Whether a semicolon is explicitly used or not, Go adds them automatically. To the compiler, each statement is on its own line.

---

## 7. Please review the following program and determine why it will not run. ##

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

(A) Operators can combine expressions.

(B) Statements can be used with operators.

(C) Expressions can return multiple values.

---

**Solution**: (A)

That's right. The `+` operator combines `10` and `runtime.NumCPU()`.

Statements can't be used with operators. For example, `import "fmt" + 3` will produce an error. However, this doesn't mean that they can be combined using expressions.

Expressions can return multiple values. However, that is not the reason why the program will run, since the expression inside `fmt.Println()` only produces a single value.
