# Questions: Comparison Operators

---

## 1. Which of the following is not an equality operator in Go? ##

(A) `==`

(B) `!=`

(C) `>`

---

**Solution**: (C)

Correct! `>` is the "greater than" operator, that checks if the value on the left is greater than the value on the right.

---

## 2. Which of the following is not an ordering operator in Go? ##

(A) `>`

(B) `<=`

(C) `==`

(D) `<`

---

**Solution**: (C)

That's right. `==` is the "equals" operator that checks if the values (or operands) on both sides have the same value.

---

## 3. Which of the following types is returned by a comparison operator? ##

(A) `int`

(B) `byte`

(C) `bool`

(D) `float64`

---

**Solution**: (C)

Spot on. Comparison operators return an untyped `bool` value (true or false).

---

## 4. Which of the following cannot be used with ordering operators? ##

**Hint**: Ordering operators are: `>`, `<`, `>=`, `<=`.

(A) `int` values

(B) `byte` values

(C) `string` values

(D) `bool` values

(E) All of the above

---

**Solution**: (D)

That's right. `bool` values are not ordered and cannot be used with ordering operations.

`int` and `byte` are valid because both are ordered values.

`string` values are ordered because they are interpreted as a series of numbers by the compiler.

---

## 5. Which of the following cannot be used with equality operators? ##

**Note**: Equality operators are: `==`, `!=`.

(A) `int` values

(B) `byte` values

(C) `string` values

(D) `bool` values

(E) All of the above can be used

---

**Solution**: (E)

Every value that can be compared to another value can be used with can be used as with equality operators.

---

## 6. What will the following code print to the console? ##

```go
fmt.Println("go" != "go!")
fmt.Println("go" == "go!")
```

(A) true true

(B) true false

(C) false true

(D) false false

(E) The compiler will produce an error.

---

**Solution**: (B)

Make sure to take note of the exclamation mark at the end of the second string value.

---

## 7. What will the following code print to the console? ##

```go
fmt.Println(1 == true)
```

(A) true

(B) 1

(C) false

(D) 2

(E) The compiler will produce an error.

---

**Solution**: (E)

Numeric constant values cannot be compared to boolean values.

---

## 8. What will the following code print to the console? ##

```go
fmt.Println(2.8 > 2.9)
fmt.Println(2.9 <= 2.9)
```

(A) true true

(A) true false

(A) false true

(A) false false

(A) error

---

**Solution**: (C)

---

## 9. What will the following code print to the console? ##

```go
fmt.Println(false >= true)
fmt.Println(true <= false)
```

(A) true true

(B) true false

(C) false true

(D) false false

(E) error

---

**Solution**: (E)

`bool` values are not ordered values, so they cannot be used with comparison operators.

---

## 10. Which of the following choices would not alter the precision of `weight`? ##

```go
package main
import "fmt"

func main() {
	weight, factor := 500, 1.5
	weight *= factor

	fmt.Println(weight)
}
```

(A) `weight *= float64(factor)`

(B) `weight *= int(factor)`

(C) `weight = float64(weight) * factor`

(D) `weight = int(float64(weight) * factor)` 

(E) None of the above.

---

**Solution**: (D)

The resulting value of the expression is 750 (the same as before).

(A) The types are mismatched because weight is an `int`.

(B) The value loses precision, the factor will have a value of 1.

(C) The value of weight is an `int` and cannot be assigned back.