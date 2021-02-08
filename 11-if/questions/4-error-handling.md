## Why error handling is needed?
1. I don't know.
2. To control the execution flow.
3. To make a program malware safe.
4. Because, things can go wrong. *CORRECT*

> **1:** Then, please watch the lecture again! :)
>
> **2:** Actually yes, but that's not the main reason. 
> 
> **3:** Come on!


## What's a nil value?
1. The dark matter that rules the universe.
2. It's a zero value for pointers or pointer-based types. It means the value is uninitialized. *CORRECT*
3. It's equal to empty string: `"" == nil` is true.


## What's an error value?
1. It stores the error details *CORRECT*
2. A global variable which stores the error status.
3. A global constant which stores the error status.

> **2, 3:** There aren't any global variables in Go. There are only package level variables. And, since the error value is just a value, so it can be stored in any variable.
> 


## How Go handles error handling?
1. Using a throw/catch statement
2. Using a simple if statement with nil comparison *CORRECT*
3. Using a mechanism called tagging

> **1:** There isn't a throw/catch statement in Go; unlike Java, C#, and so on... Go is explicit.


## When you should handle the errors?
1. After the main func ends.
2. Before calling a function.
3. Immediately, after calling a function which returns an error value. *CORRECT*


## For which one of the following functions that you might want to handle the errors?
```go
func Read() error
func Write() error
func String() string
func Reset()
```
1. Read and Write *CORRECT*
2. String and Reset
3. Read, Write and Reset
4. For neither of them
5. For all of them

> **1:** They return error values. So, you might want to handle the errors after you call them.
> 
> **2:** They don't return error values. So, you don't have to handle any errors.
> 
> **3:** Partially true. Try again.


## Let's say a function returns a nil error value. So, what does that mean?
1. The function call failed.
2. The function call is successful. *CORRECT*
3. The function call is in an indeterministic state. We can't know.


## Let's say a function returns a non-nil error value. So, what does that mean?
1. The function call failed. *CORRECT*
2. The function call is successful.
3. The function call is in an indeterministic state. We can't know.

> **1:** Yep. Later on you'll learn that, this is not always true. Sometimes a function can return a non-nil error value, and the returned value may indicate something rather than an error. Search on Google: golang io EOF error if you're curious.


## Does the following program correctly handles the error?

**NOTE:** This is what the `ParseDuration` function looks like:

```go
func ParseDuration(s string) (Duration, error)
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h10s")
	if err != nil {
		fmt.Println(d)
	}
}
```
1. Yes. It prints the parsed duration if it's successful.
2. No. It doesn't check for the errors.
3. No. It prints the duration even when there's an error. *CORRECT*

> **1:** Yes, it handles the error; however it does so incorrectly. Something is missing here. Look closely.
> 
> **2:** Actually, it does.
> 
> **3:** That's right. It shouldn't use the returned value when there's an error.


## Does the following program correctly handles the error?

**NOTE:** This is what the `ParseDuration` function looks like:

```go
func ParseDuration(s string) (Duration, error)
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h10s")
	if err != nil {
        fmt.Println("Parsing error:", err)
        return
    }
    fmt.Println(d)
}
```
1. Yes. It prints the parsed duration if it's successful. *CORRECT*
2. No. It doesn't check for the errors.
3. No. It prints the duration even when there's an error.

> **1:** That's right. When there's an error, it prints a message and it quits from the program.
> 
> **2:** Actually, it does.
> 
> **3:** No, it does not. It only prints it when there isn't an error.
