# Questions: Short If #

## 1. Which of the choices below would fix the following program? ##

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	if err == nil; d, err := time.ParseDuration("1h10s") {
		fmt.Println(d)
	}
}
```

(A) Swapping the simple statement with `err != check`.

(B) Removing the error handling.

(C) Removing the remove the semicolon.

(D) Changing the short declaration to an assignment.


---

**Solution**: (A)

That's right! When initializing short if-statements, the simple statement (the short-declaration) should come first. Only after the semicolon should there be a conditional expression.

Removing the error handler is not the issue and semicolons are valid end-of-statement separators in Go. 

Changing the short declaration to an assignment declaration would not work because then it would no longer be a conditional.

---

## 2. What does the following program print to the console? ##

```go
package main 
import "fmt"

func main() {
	done := false
	if done := true; done {
		fmt.Println(done)
	}
	fmt.Println(done)
}
```

(A) true and false

(B) false and false

(C) true and false

(D) false and true


---

**Solution**: 

That's right. The program shadows the variable `done` in the `main()` function. 

Inside the if-statement, it prints `true`. The program prints `false` at the end because the values inside the if-statement don't affect the values of variables outside of it.

---

## 3. There's a shadowing issue in the following program. The program should print `10` but instead prints `0`. Which of the following choices would fix the program? ##

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	
	if n, err := strconv.Atoi("10"); err != nil {
		fmt.Printf("error: %s (n: %d)", err, n)
		return
	} 
	
	fmt.Println(n)
}
```

(A) Removing the first declaration (`var n int`)

(B) Removing the declaration in the short if (`if n, err..`)

(C) Declaring an error variable outside of `main()`.

(D) Declaring an error variable along with the variable `n` in `main()` and changing the short `if` declaration to an assignment (`if n, err = strconv.Atoi("10");...`)

---

**Solution**: (D)

The short if-statement initializes a second variable called `n` to the value of 10. After reaching the end of the if-statement, the second `n` goes out of scope. Therefore, when the `main()` function calls `fmt.Println(n)`, it prints the initial value of the first `n` variable.

Click [here](https://play.golang.org/p/6jc-nZJNHYb) to see the solution on Go Playground.