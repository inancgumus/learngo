## Which one is the correct description for a statement?
1. A statement instructs Go to do something *CORRECT*
2. A statement produces a value
3. A statement can't change the execution flow

> **2:** A statement can't produce a value. However, it can indirectly help to produce a value.
>
>
> **3:** It surely can.
>
>


## What's the direction of execution in a Go code?
1. From left to right
2. From top to bottom *CORRECT*
3. From right to left
4. From bottom to top

> **2:** That's right. Go executes the code from top-to-bottom, one statement at a time.
>
>


## Which one is the correct description for an expression?
1. An expression instructs Go to do something
2. An expression produces a value *CORRECT*
3. An expression can change the execution flow

> **1:** It can't. Only a statement can do that.
>
>
> **3:** It can't. Only a statement can do that.
>
>


## Which one is the correct description for an operator?
1. An operator instructs Go to do something
2. An operator can change the execution flow
3. An operator can combine expressions *CORRECT*

> **1:** It can't. Only a statement can do that.
>
>
> **2:** It can't. Only a statement can do that.
>
>


## Why doesn't the following program work?
```go
package main
import "fmt"

func main() {
    "Hello"
}
```

1. "Hello" is an expression and it can't be on its own on a single line of code without a statement. *CORRECT*
2. By removing the double-quotes surrounding the "Hello". Like this: Hello
3. By moving "Hello" out of the func main.


## Does the following program work?
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU()); fmt.Println("cpus"); fmt.Println("the machine")
}
```

1. It works: Expressions can be typed by separating them using semicolons
2. It doesn't work: Statements should be on their own on a single line of code
3. It works: Go adds semicolons behind the scenes for every statement already *CORRECT*

> **1:** It works but that's not the reason. And, expressions can't be typed like that.
>
>
> **2:** Are you sure?
>
>
> **3:** That's right. Whether there's a semicolon or not; Go adds them automatically. Those statements are still assumed as they're on a different code line of their own.
>
>


## Why does this program work?
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

1. Operators can combine expressions *CORRECT*
2. Statements can be used with operators
3. Expressions can return multiple values

> **1:** That's right. + operator combines `runtime.NumCPU()` and `10` expressions.
>
>
> **2:** No, they can't be. For example, you can't do this: `import "fmt" + 3`. Some statement can allow expressions. However, this doesn't mean that they can be combined using expressions.
>
>
> **3:** That's right however it's irrelevant to why this code works.
>
>
