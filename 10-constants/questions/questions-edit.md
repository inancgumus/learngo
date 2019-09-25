# Questions: Constants #

## 1. What's a magic value? ##
(A) An unexpected value that pops up from somewhere. <br>
(B) A spell case by Merlin the Wizard. <br>
(C) An unnamed constant value in the source code. <br>
(D) A named constant. <br>

---
**Solution**: (C) <br>
A `magic value` is a value explicitly stated value in a program that can prevent an error from being detected until runtime.

---
## 2. What's a named constant? ##
(A) A constant with cool name. <br>
(B) A constant value with a declared name. <br>
(C) A literal value converted to a name. <br>

---
**Solution**: (C) <br>
A named constant uses the syntax `const name = "value"`.

---
## 3. Which of the following choices shows the proper syntax for declaring a constant? ##
(A) `Const version int = 3` <br>
(B) `const version int := 3` <br>
(C) `const version int = 3` <br>

---
**Solution**: (C) <br>
the `const` keyword is case-sensitive so `Const` would not work.

---
## 4. Which of the following choices shows correct syntax? ##
(A) `s := "pick me"; const length = len(s)` <br>
(B) `const message = "pick me!"; const length = len(message)` <br>
(C) `const length = utf8.RuneCountInString("pick me")`

---
**Solution**: (B) <br>
d(A) `s` is not a `const` type because it is missing the `const` keyword before its name. <br>
(C) You cannot call functions to return a value in the same statement a constant is being initialized. <br>

---
## 5. Which explanation below best explains the following code block? ##
```go
const speed = 100
porsche := speed * 3
```
(A) `speed` is typless and `porsche` is an `int` <br>
(B) `speed` is an `int` and `porsche` is also an `int` <br>
(C) `speed` and `porsche` are both typeless <br>

---
**Solution**: (A) <br>
The `speed` is declared without a `type`, but `porsche` is an `int` because it produces the integer value of `speed * 3`.

---
## 6. Which of the following choices fixes the error in the original code? ##
(A) <br>
```go
const total int8 = 10
x := 5

fmt.Print(total * x)
```

(B) <br>
```go
const total = 10
x := 5

fmt.Print(total * x)
```

(C) <br>
```go
const total int64 = 10
x := 5

fmt.Print(total * x)
```

(D) <br> 
```go
const total int64 = 10
x := 5

fmt.Print(int64(total) * x)
```
---
**Solution**: (A) <br>
The constant `total` is typeless, so it can now be used with the variable `x`
<br>
(B) There is still a  type mismatch, since `x` is an `int` and not `int64`. 
<br>
(C) `total` is already an `int64`, so converting the type is redundant.

---
## 7. What are the values of the following constants? ##
```go
const (
    Yes = (iota * 5) + 2
    No
    Both
)
```
(A) `Yes`=0; `No`=1; `Both`=2 <br>
(B) `Yes`=2; `No`=3; `Both`=4 <br>
(C) `Yes`=7; `No`=12; `Both`=17 <br>
(D) `Yes`=2; `No`=7; `Both`=12 <br>

---
**Solution**: (C) <br>
`iota` starts at 0 and not 1.