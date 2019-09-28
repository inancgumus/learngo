# Questions: `if` #

---

## 1. What does "control flow" mean?

(A) Changing the top-to-bottom execution of a program

(B) Changing the left-to-right execution of a program

(C) Deciding which statements are executed

---

**Solution**: (C)

Go will always execute code from top-to-bottom going left-to-right on each line. 

Control flow allows us to decide which parts of a program will run depending on conditional values like `true` and `false`.

---

## 2. Which of the following is choices simplifies the `if` statement below? ##

```go
if (mood == "perfect") {
    // ...
}
```

(A) `if {mood == "perfect"} ...`

(B) `if [mood == "perfect"] ....`

(C) `if mood = "perfect" ...`

(D) `if mood == "perfect"...`

---

**Solution**: (D)

Go does not require parentheses for conditional statements.

Using `{}` and `[]` will is will produce a syntax error. 

`=` is an assignment operator and cannot be used as a condition.

---

## 3. Which of the following choices will allow the code to run? ##

```go
package main

import "fmt"

func main() {
    // this program prints "cool"
    // when the mood is "happy"

    mood := "happy"

    if "happy" {
        fmt.Println("cool")
    }
}
```

(A) Wrapping the `"happy"`` after the `if` inside parentheses

(B) Comparing "happy" with the variable `mood`: `if mood == "happy"`

(C) Directly using `mood` instead of `"happy"`: `if mood {...}`

(D) Trick question - no changes are required for the program to run.

---

**Solution**: (B)

Conditional expressions always yield a `bool` value. By using `==` (the equality operator), we allow the program to continue executing.

---

## 4. How can the following program be simplified? ##

**Hint**: Only one condition needs to be changed.

```go
package main

import "fmt"

func main() {
    happy := true

    if happy == true {
        fmt.Println("cool!")
    }
}
```

(A) `happy != false`

(B) `!happy == false`

(C) `happy`

(D) `!happy == true`

---

**Solution**: (C)

The value of happy is already `true`, making the condition `true` as well, which will call the print function.

`!happy == true` returns `false` so the program will not print anything.

`happy != false` and `!happy == false` are also valid but are verbose compared to.

---

## 5. How can the following program be simplified? ##
```go
package main

import "fmt"

func main() {
    happy := false
    
    if happy == !true {
        fmt.Println("why not?")
    }
}
```
(A) Changing `happy != false` to `happy == !true`
(B) `!happy`
(C) `happy == false`
(D) `!happy == false`

---

**Solution**: (B)

Perfect! `!happy` already. returns `true` because it is already `false. There is no need to compare the value to `false` or `!true`.

`!happy == false` will cause the program not to print because the condition won't be met, since `happy` will be `true`.

---

## 6. The following program contains an error. Which of the following answer choices fixes the error? ##

```go
package main

import "fmt"

func main() {
    happy := false

    if happy {
        fmt.Println("cool!")
    } else if !happy {
        fmt.Println("why not?")
    } else {
        fmt.Println("why not?")
    } else {
        fmt.Println("why not?")
    }
}
```

(A) Remove one of the `else` branches.

(B) Moving the `else if` to the last branch.

(C) Changing the values in the print statements to be unique.

(D) Removing one of the `else if` branches.

---

**Solution**: (A)

Only one `else` branch can be used within the same `if` statement.

The same statements can be executed based on different conditions as long as they are defined after each `else if` or `else`.

If an `else` is at the end of the `if` block, the `else if` cannot be moved to the end.

---

## 7. Where is the error in the following program?
```go
package main

import "fmt"

func main() {
    happy := true
    energic := happy

    if happy {
        fmt.Println("cool!")
    } else if !happy {
        fmt.Println("why not?")
    } else if energic {
        fmt.Println("working out?")
    }
}
```

(A) The variable `energic` is declared but never used.

(B) More than one `else if` branch is used.

(C) The last branch will never run.

(D) The `else` branch is missing.

---

**Solution**: (C)

Because `happy` is a `bool`, it can only be one of two values: `true` or `false`, so only two possible outcomes are possible.

`else if` can be declared more than once and the `else` branch is not required to following an `if`.

---

## 8. How can the following program be simplified? ##

```go
package main

import "fmt"

func main() {
    happy := false
    
    if happy {
        fmt.Println("cool!")
    } else if happy != true {
        fmt.Println("why not?")
    } else {
        fmt.Println("why not?") 
    }
}
```

(A) Changing the `else if`'s condition to `!happy`.

(B) Moving the `else` branch before the `else if`

(C) Removing the `else` branch

(D) Removing the `else if` branch

---

**Solution**: (D)

The `else if` branch is not necessary because the `else` branch already handles the `not happy` situation.

It's simpler and less verbose to not have a condition in this case. 
