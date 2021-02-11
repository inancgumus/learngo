## What's the difference between if and switch statements?
1. If statements control the execution flow but switch statements do not
2. If statements are much more readable alternatives to switch statements
3. Switch statements are much more readable alternatives to if statements *CORRECT*

> **1:** They both control the execution flow.
> 
> **2:** Sometimes true, but, for complex if statements, switch statement can make them much more readable.


## What type of values can you use as a switch condition?
```go
switch condition {
    // ....
}
```
1. int
2. bool
3. string
4. Any type of values *CORRECT*

> **4:** Unlike most other C based languages, in Go, a switch statement is actually a syntactic-sugar for a if statement. This means that, Go converts a switch statement into an if statement behind the scenes. So, any type of values can be used as a condition.


## What type of values can you use as the case conditions for the following switch statement?
```go
switch false {
case condition:
    // ...
}
```
1. int
2. bool *CORRECT*
3. string
4. Any type of values

> **2:** Yes, you can only use bool values because in the example, the switch's condition is a bool.


## What type of values can you use as the case conditions for the following switch statement?
```go
switch "go" {
case condition:
    // ...
}
```
1. int
2. bool
3. string *CORRECT*
4. Any type of values

> **3:** Yes, you can only use string values because in the example, the switch's condition is a string.


## What's the problem in the following code?
```go
switch 5 {
case 5:
case 6:
case 5:
}
```
1. Case clauses don't have any statements
2. The same constants are used multiple times in case conditions *CORRECT*
3. Switch condition cannot be an untyped int

> **2:** That's right. 5 is used by multiple case clauses as case conditions. It should be used only once.


## When the following code runs, which case will be executed?
```go
weather := "hot"
switch weather {
case "very cold", "cold":
    // ...
case "very hot", "hot":
    // ...
default:
    // ...
}
```
1. None of them
2. The first case clause
3. The second case clause *CORRECT*
4. The default clause

> **3:** That's right. When switch's condition is either "hot" or "very hot", the 2nd case will be executed.


## When the following code runs, which clause will be executed?
```go
switch weather := "too hot"; weather {
default:
    // ...
case "very cold", "cold":
    // ...
case "very hot", "hot":
    // ...
}
```
1. None of them
2. The first case clause
3. The second case clause
4. The default clause *CORRECT*

> **4:** That's right. The switch's condition doesn't match to any of the case conditions, so the default clause will be executed as a last resort. And the default clause can be in any place inside a switch statement.


## When the following code runs, which case will be executed?
```go
switch weather := "too hot"; weather {
case "very cold", "cold":
    // ...
case "very hot", "hot":
    // ...
}
```
1. None of them *CORRECT*
2. The first case clause
3. The second case clause
4. The default clause

> **1:** That's right. The switch's condition doesn't match to any of the case conditions, and there isn't a default clause. So, none of the clauses will be executed.


## What does the following program print?
```go
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
```
1. Nothing
2. "Not important" *CORRECT*
3. "Destructive"

> **2:** That's right! When fallthrough is used, the following clause will be executed without checking its condition.


## What's the problem in the following code?
```go
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
```
1. default clause should be the last clause
2. default clause should have a fallthrough statement
3. The first case should use the fallthrough statement as its last statement *CORRECT*

> **3:** That's right! In a case block, fallthrough can only be used as the last statement.


## What does the following program print?
```go
n := 8

switch {
case n > 5, n < 1:
    fmt.Println("n is big")
case n == 8:
    fmt.Println("n is 8")
}
```
1. Nothing
2. "n is 8"
3. "n is big" *CORRECT*

> **3:** That's right! Switch runs top-to-bottom and case conditions run left-to-right. Here, 1st case's 1st condition expression (which is n > 5) will yield true, so the 1st case will be executed.
