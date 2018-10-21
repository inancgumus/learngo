## How to fix this program?
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	if err != nil; d, err := time.ParseDuration("1h10s") {
		fmt.Println(d)
	}
}
```
1. Swap the simple statement with the `err != nil` check. *CORRECT*
2. Remove the error handling.
3. Remove the semicolon.
4. Change the short declaration to an assignment.

> **1:** Yes. In a short if statement, the simple statement (the short declaration there) should be the first part of it. Then, after the semicolon separator, there should be a condition expression.
> 
> **2:** You don't want that. That's not the issue here.


## What does this program print?
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
1. true and true
2. false and false
3. true and false *CORRECT*
4. false and true

> **3:** Yes. It shadows the main()'s done variable, and inside the if statement, it prints "true". Then, after the if statement ends, it prints the main()'s done variable which is "false".


## How can you fix this code?
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
1. Remove the first declaration (main()'s done)
2. Remove the declaration in the short-if (if's done)
3. Change the done declaration of the main() to an assignment
4. Change the done declaration of the short-if to an assignment. And, after the if statement, assign false back to the done variable. *CORRECT*

> **1:** That will break the program. Because, the last line prints it.
> 
> **2:** The program wants to use it to print true.
> 
> **3:** There will be "undefined variable" error.
> **4:** Yes, that will solve the shadowing issue. Short-if will reuse the same done variable of the main(). And, after the short-if, done will be false because of the assignment, and it will print false.

