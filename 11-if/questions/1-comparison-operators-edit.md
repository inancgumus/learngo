## Which one below is not one of the equality operators of Go? ##

# Questions: Comparison Operators
---
## 1. Which of the following is not an equailty operator in Go? ##
(A) `==` <br>
(B) `!=` <br>
(C) `>` <br>

---
**Solution**: (C) <br>
`>` is the "greater than" operator that checks if the value on the left is greater than that value on the right.

---
## 2. Which of the following is not an ordering operator in Go? ##

## Which one below is not one of the ordering operators of Go?
(A) `> <br>
(B) `<=` <br>
(C) `==` <br>
(D) `<` <br>

---
**Solution**: (C) <br>
`==` is the "equals" operator that checks if the value (or operand) on both sides have the same value.

---
## 3. Which of the following types if returns by a comparison operator? ##
(A) `int` <br>
(B) `byte` <br>
(C) `bool` <br>
(D) `float64` <br>

---
**Solution**: (C)
Comparison operators return an untyped `bool` value (true or false).

---
## 4. Which of the following cannot be used as operands with ordering operators? ##
**Hint**: Ordering operators are: `>`, `<`, `>=`, `<=`.

(A) `int` values <br>
(B) `byte` values <br>
(C) `string` values <br>
(D) `bool` values <br>
(E) All of the above

---
**Solution**: (D) <br>
`int` and `byte` are valid becuase both are ordered values. <br>
`string` values are ordered because they are interpreted as a series of numbers by the compiler.
`bool` values are not ordered and cannot be used with ordering operations.

---
## 5. Which of the following cannot be used with equality operators? ##
**Note**: Equality operators are: `==`, `!=`.
(A) `int` values <br>
(B) `byte` values <br>
(C) `string` values <br>
(D) `bool` values <br>
(E) All of the above can be used <br>

---
**Solution**: (E) <br>
Every value that can be compared to another value can be used with can be used as with equality operators.

---
## What does the following code print to the console? ##
```go
fmt.Println("go" != "go!")
fmt.Println("go" == "go!")
```
(A) true true <br>
(B) true false <br>
(C) false true <br>
(D) false false <br>
(E) The compiler will throw an error.

---
**Solution**: (C) <br>
Make sure to take note of the exclamation mark at the end of the second string value.

## 6. What will following code print to the console? ##
```go
fmt.Println(1 == true)
```
(A) true <br>
(B) 1 <br>
(C) false <br>
(D) 2 <br>
(E) The compiler will throw an error.

---
**Solution**: (E) <br>
Numeric constant values cannot be compared to boolean values.

---
## 7. What will the following code print to the console? ##
```go
fmt.Println(2.9 > 2.9)
fmt.Println(2.9 <= 2.9)
```
(A) true true <br>
(A) true false <br>
(A) false true <br>
(A) false false <br>
(A) error <br>

---
**Solution**: (C)

---
## 8. What will the following code print to the console? ##
```go
fmt.Println(false >= true)
fmt.Println(true <= false)
```
(A) true true <br>
(B) true false <br>
(C) false true <br>
(D) false false <br>
(E) error

---
**Solution**: (E) <br>
`bool` values are not ordered values, so they cannot be used with comparison operators.

---
## 9. Which of the following choices doinwould not alter the precision of the `weight`? ##
```go
package main
import "fmt"

func main() {
    weight, factor := 500, 1.5
    weight *= factor

    fmt.Println(weight)
}
```
(A) `weight *= float64(factor)` <br>
(B) `weight *= int(factor)` <br>
(C) `weight = float64(weight) * factor` <br>
(D) `weight = int(float64(weight) * factor)`  <br>
(E) None of the above.

---
**Solution**: (D) <br>
The resulting value of the expression is 750 (the same as before). <br>
(A) The types are mismtached because weight is an `int`. <br>
(B) The value loses precision with a factor of 1. <br>
(C) The value of weight is an `int` and cannot be assigned back. <br>