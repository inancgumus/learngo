# Questions: Switch #

## 1. What is the difference between `if` and `switch` statements? ##

(A) `if` statements control a program's execution flow; `switch` statements do not

(B) `if` statements are a more readable than `switch` statements

(C) `switch` statements are much more readable than `if` statements

---

**Solution**: (C)

Both `if` and `switch` statements control execution flow.

`if` statements **can** be more readable than `switch` statements, but as `if` statements become increasingly complex, `switch` statements are more readable.

---

## 2. What type of values be used as the condition of a `switch` statement? ##

(A) `int`

(B) `bool`

(C) `string`

(D) Any value can be used.

---

**Solution**: (D)

Unlike most other C-based languages, Go uses `switch` statements as syntactical-sugar for `if` statements. This means that Go converts `switch` statements into `if` statements behind the scenes.

Because of this, any type that can be used as an `if` condition can also be used as `switch` conditions.

---

## 3. Which of the following types can the switch statement's case condition use? ##

```go
switch false {
case condition:
    //
}
```

(A) `int`

(B) `bool`

(C) `string`

(D) Any type may be used.

---

**Solution**: (B)

That's right. Because the `switch` statement's condition is a `bool`, the case conditions must also use `bool` values.

---

## 4. Which of the following types can the switch statement's case condition use? ##

```go
switch "go" {
case condition:
    // ...   
}
```

(A) `int`

(B) `bool`

(C) `string`

(D) Any type may be used.

---

**Solution**: (C)

Correct. Because the `switch` statement uses a `string`, the `case` conditions must also use strings.

---

## 5. Which of the following choices identifies the problem with the `switch` statement below? ##

```go
switch 5 {
case 5:
case 6
case 5:
}
```

(A) The `case` clauses don't have any statements.

(B) The same constants are used for more than one `case` condition.

(C) Switch conditions cannot be an `untyped` int.

---

**Solution**: (B)

You got it! `5` is the condition for multiple clauses. Each case clause's condition should unique.

---

## 6. Which clause will be executed when the program is run? ##

```go
package main
import "fmt"

func main() {
    weather := "hot"
    switch weather {
    case "very cold", "cold":
        // ...
    case "very hot", "hot":
        // ...
    default:
        // ...
    }
}
```

(A) The first case clause

(B) The second case clause

(C) The default clause

(D) None of the above

---

**Solution**: (C) 

That's right! The second case clause is executed because one of the conditions (`"very hot", "hot"`) matches the switch's condition (`"hot"`).
---

## 7. Which clause will be executed when the program is run? ##

```go
package main
import "fmt"

func main() {
    switch weather := "too hot"; weather {
    case "very cold", "cold":
        // ...
    case "very hot", "kind of":
        // ...
    }
}
```

(A) The first case clause

(B) The second case clause

(C) The default case clause

(D) None of the above

---

**Solution**: (A)

That's right! None of the clauses will be executed because there are no case conditions that match the switch's condition and there the switch does not have a default clause.

---

## 8. What will the following program print? ##

```go
package main
import "fmt"

func main() {
    richter := 2.5
    
    switch r := richter; {
    case r < 2:
        fallthrough
    case r >= 2 && r < 3:
        fallthrough
    case r >= 5 && r < 6:
        fmt.Println("Not important")
    case r >= 6 && r < 7:
        fallthrough
    case r >= 7:
        fmt.Println("Destructive")
    }
}
```

(A) "Not important"

(B) "Destructive"

(C) The program won't print anything

---

**Solution**: (B)

Spot on! When the fallthrough is used, the following clauses will be executed without checking its condition.

---

## 8. Which of the following choices identifies the problem in the following program? ##

```go
package main
import "fmt"

func main() {
    richter := 14.5

    switch r := richter; {
    case r >= 5 && r < 6:
        fallthrough
        fmt.Println("Moderate")
    default:
        fmt.Println("Unknown")
    case r >= 7:
        fmt.Println("Destructive")
    }
} 
```

(A) The default clause should be the last clause

(B) The default clause should have a fallthrough statement

(C) The first case should the fallthrough statement as its last statement

---

**Solution**: (C)

That's correct! The fallthrough statement should be the last statement in a case block.

---

## 9. What does the following program print? ##

```go
package main
import "fmt"

func main() {
    n := 8

    switch {
    case n > 5, n < 1:
        fmt.Println("n is big")
    case n == 8:
        fmt.Println("n is 8")
    }
}
```

(A) "n is 8"

(B) "n is big"

(C) Nothing 

---

**Solution**: (B)

That's right! The switch statement runs top-to-bottom and case conditions run left-to-right.

The first case's first condition expression (`n > 5`) yields true, so the first case will be executed.

---

