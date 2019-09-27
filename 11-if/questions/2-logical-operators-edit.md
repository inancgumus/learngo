# Questions: Logical Operators #

---

## 1. Which of the following choices is not a logical operator in Go? ##

(A) `||`

(B) `!=`

(C) `!`

(D) `&&`

---

**Solution**: (B)

`!=` is the "not equals" operator. It is a comparison operator, not a logical operator.

---

## 2. Which of the following types is returned by logical operators? ##

(A) `int`

(B) `byte`

(C) `bool` 

(D) `float64`

---

**Solution**: (C)

That's right. Logical operators return an untyped `bool` value (either `true` or `false`).

---

## 3. Which fo the following types can be used with logical operators?  ##

(A) `int`

(B) `byte`

(C) `bool`

(D) `float64` 

---

**Solution**: (C)

All logical operators expect a `bool` values (or an expression the yields a `bool` value).

---

## 4. Which of the following choices translates the sentence below into Go code? ##

`"age is greater than or equal to than 15 and hair color is yellow`

(A) `age > 15 || hairColor == "yellow"`

(B) `age < 15 || hairColor != "yellow"`

(C) `age >= 15 && hairColor == "yellow"`

(D) `age > 15 && hairColor == "yellow"`

--- 

**Solution**: (C)

`"greater than"` is expressed as `>=`.

`"and"` is expressed as `&&`.

`"is"` translates to `"equals to"`, which is expressed as `==`.

---

## 5. What does the following program print to the console? ##

```go
package main
import "fmt"

func main() {
	on = true
	off = !on
	
	fmt.Println(!on && !off)
	fmt.Println(!on || !off)
}

```

(A) true true

(B) true false

(C) false true

(D) false false 

(E) The compiler will produce an error.

---

**Solution**: (C)

Because `!on` is `false` and `!off` is true, `!on && !off` evaluates to `false` and `!on || !off` evaluates to `true`

---

## 6. What does the following program print to the console? ##

```go
package main
import "fmt"

func main() {
	on := 1
	fmt.Println(on == true)
}
```

(A) true

(B) false

(C) The compiler will produce an error.

---

**Solution**: (C)

The compiler will produce a type mismatch error because `on` is an `int` while `true` is a `bool`. 

Unlike other C-based languages, Go does not interpret `0` as `false` and `1` as `true`.

---

## 7. What will the following print to the console? ##
```go
a := "a" > "b"
b := "b" <= "c"
fmt.Println(a || b)
```
**Note**: `"a"` comes before `"b"`.

(A) `"a"`

(B) `"b"`

(C) `true`

(D) `false`

(E) The compiler will produce and error.

---

**Solution**: (C)

The order is as follows: `"a", "b", "c"`. 

`"a" > "b"` evaluates to `false` and `"b" <= "c"` evaluates to `true`, so `fmt.Println(a || b)` will print true.

Logical operators only return `bool` values "a" and "b" (`string` values) are not valid options.

---

## What will the following program print the console? ##

**Note**: Make to read the comments!

```go
/*
   There are two functions 
*/
```

## What will the following program print? ##
**Note**: Read the comments and focus on the `main()` function to solve the problem.

```go
package main
import "fmt"
/*
   There are two functions defined below:
   1. `isOn()`: returns `true` and prints "on "
   2. `isOff()`: returns `false` and prints "off "
*/
func main() {
	_ = isOff() && isOn()
	_ = isOn() || isOff()
}

func isOn() bool {
	fmt.Println("on ")
	return true
}

func isOff() bool {
	fmt.Println("off ")
	return false
}
```

(A) `on off on off `

(B) `off on `

(C) `off on on off `

(D) `on off `

---

**Solution**: (B)

**`_ = isOff() and isOn()`**: `isOff()` returns `false` so `&&` (the logical AND operator) will short-circuit the function, meaning `isOn()` never gets called, meaning `of ` prints to the screen.

**`_ = isOn() || isOff()`**: `isOn()` returns `true` so `||` (the logical OR operator) short circuits and `isOff()` is never called, printing `on ` to screen.

When `main()` finishes executing, the final printed result is `off on`.

You can view an example of the program [here](https://play.golang.org/p/6z3afaOf7yT)