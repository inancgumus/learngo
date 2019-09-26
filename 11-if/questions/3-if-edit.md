# Questions: `if` #

## 1. What does "control flow" mean?
(A) Changing the top-to-bottom execution of a program <br>
(B) Changing the left-to-right execution of a program <br>
(C) Deciding which statements are executed <br>

---
**Solution**: (C) <br>

Go will always execute code from top-to-bottom going left-to-right on each line. Control flow allows us decide which parts of a program will run depending on conditional values like `true` and `false`.

---

## 2. How can the `if` statement below be simplified?
```go
if (mood == "perfect") {
    // this code is not important
}
```
1. `if {mood == "perfect"}`
2. `if [mood == "perfect"]`
3. `if mood = "perfect"`
4. `if mood == "perfect"` *CORRECT*

---

**Solution**: (D) <br>
Go does not required parentheses for conditional statements. <br>
(B) Using `{}` and `[]` will is will throw a syntax error. <br> 
(C) `=` is an assignment operator and cannot be used as a condition. <br>

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
(A) Wrapping the `"happy"`` after the `if` inside parentheses <br>
(B) Comparing "happy" with the variable `mood`: `if mood == "happy"` <br>
(C) Directly using `mood` instead of `"happy"`: `if mood {...}` <br>
(D) Trick question - no changes are required for the program to run. <br>

---

**Solution**: (B) <br>
By using `==`, the equality operator, we allow are program to continue executing because the logic within the `if` block. <br>

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
(A) `happy != false` <br>
(B) `!happy == false` <br>
(C) `happy` <br>
(D) `!happy == true` <br>

---

**Solution**: (C) <br>
The value of happy is already `true`, making the condition `true` as well, which will call the print function. <br>
`!happy == true` returns false so the program will not print anything. <br>
`happy != false` and `!happy == false` are also valid but are verbose compared to (C). <br>

---

## 5. How can the following program be simplified? ##
```go
package main

import "fmt"

func main() {
    happy != false
    
    if happy == !true {
        fmt.Println("why not?")
    }
}
```
(A) Changing `happy != false` to `happy == !true` <br>
(B) `!happy` <br>
(C) `happy == false` <br>
(D) `!happy == false` <br>

---

**Solution**: (B) <br>
`!happy` already returns`true` because the value is changed to `false` or `!true`. <br>
(D) won't print anything because the condition isn't met. <br>
(A) and (C) are valid examples but are verbose compared to (C)

---

## 6. The following contains an error. Which of the following answer choices fixes the error? ##

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
(A) Remove one of the `else` branches. <br>
(B) Moving the `else if` to the last branch. <br>
(C) Changing the values in the print statements to be unique. <br>
(D) Removing one of the `else if` branches. <br>

---

**Solution**: (A) <br>
Only one `else` branch can be used following an `if` statement. <br>

The same statements can be executed based on different conditions as long as they are defined after each `else if` or `else`. <br>
If an `else` is at the end of the `if` block, the `else if` cannot be moved to the end. <br>

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
(A) The variable `eneregic` is declared but enver used. <br>
(B) More than one `else if` branch is used. <br>
(C) The last branch will never run. <br>
(D) The `else` branch is missing. <br>

---

**Solution**: (C) <br>
Because `happy` is a `bool`, it can only be one of two values: `true` or `false`, so only two possible outcomes are possible.

`else if` can be declared as more than once and the `else` branch is not required to following an `if`.

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
(A) Changing `else if`'s condition to `!happy`. <br>
(B) Moving the `else` branch before the `else if` <br>
(C) Removing the `else` branch <br>
(D) Removing the `else if` branch <br>


1. Change `else if`'s condition to: `!happy`. 
2. Moving the else branch before else if.
3. Remove the else branch.
4. Remove the else if branch. *CORRECT*

---

**Solution**: (D) <br>
The `else if` branch is not necessary because the `else` branch already handles the `not happy` situation. <br>
It's simpler and less verbose to not have a condition in this case. 
