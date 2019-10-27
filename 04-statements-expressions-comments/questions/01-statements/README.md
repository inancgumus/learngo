# Questions: Statements #

---

## (1) Which of the following best describes a statement? ##

1. A statement instructs Go to do something.
2. A statement produces a value.
3. A statement can't change the execution flow of a program.

---

**Solution: 1**

That's right. A statement like `fmt.Println("hi")` tells go to print something to the console.

> 2. A statement can't produce a value. However, it can indirectly help produce a value.
> 3. Statements like `if` may change the execution flow of a program.

---

## (2) In what direction does Go execute code? ##

1. From left to right.
2. From top to bottom.
3. From right to left.
4. From bottom to top.

---

**Solution: 2**

That's right. Go executes code from top to bottom, executing one statement at a time.

---

## (3) Which of the following best describes an expression? ##

1. An expression instructs Go to do something.
2. An expression produces a value.
3. An expression can change the execution flow of a program.

---

**Solution: 2**

1. Only statements can instruct Go to do something and3. change the execution flow of a pr
gram.

---

## (4) Which one of the following best describes an operator? ##

1. An operator instructs Go to do something.
2. An operator can change the execution flow of a program.
3. An operator can combine expressions.

---

**Solution: 3**

If you combine string values with the operator `+` between them, the resulting value is a string literal.

> 1. Only statements can instruct tell Go to do something
> 2. Only statements can change the execution flow of a program.

---

## (5) Please review the following program and determine why it will not run. ##

```go
package main
import "fmt"

func main() {
    "Hello"
}

```

1. Because "Hello" is an expression and cannot exist on a single line of code without a statement.
2. Because the double-quotes surrounding "Hello" must be removed.
3. Because "Hello" must be removed out of the `main()` function.

---

**Solution: 1**

> 3. `"Hello"` is just a string literal value and without being used within a statement, the Go compiler will produce an error.

---

## (6) Please review the following program and determine if the program will run. ##

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

1. It will run: You can type expressions by separating them using semicolons.
2. It won't work: You can only use a single statement on a single line of code.
3. It will run: Go automatically adds a semicolon at the end of every statement behind the scenes.

---

**Solution: 3**

Whether a semicolon is explicitly used or not, Go adds them automatically. To the compiler, each statement is on its own line.

---

## (7) Please review the following program and determine why it will run. ##

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

1. Operators can combine expressions.
2. Statements can be used with operators.
3. Expressions can return multiple values.

---

**Solution: 1**

That's right. The `+` operator combines `10` and `runtime.NumCPU()`.

> 3. Expressions can return multiple values. However, that is not the reason why the program will not run, since the expression inside `fmt.Println()` only produces a single value.