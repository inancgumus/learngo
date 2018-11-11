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


## There's a shadowing issue in this program. The program should print: 10, now it prints: 0.

**How can you fix this it?**

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

_See the code also in here: https://play.golang.org/p/fDrmcXWGnQB_

1. Remove the first declaration (main()'s n variable).
2. Remove the declaration in the short-if (if's n).
3. Declare an error variable outside of the main.
4. Declare an error variable along with main's n variable, then change the short-if declaration to an assignment. *CORRECT*

> **1:** That will break the program. Because, the last line prints it.
> 
> **2:** The program uses it to set the n.
> 
> **3:** There will be "unused variable" error.
> **4:** Yes, that will solve the shadowing issue. Short-if will reuse the same 'n' variable (main()'s n). So, the converted number 'n' will be printed as: 10. See the solution here: https://play.golang.org/p/6jc-nZJNHYb

