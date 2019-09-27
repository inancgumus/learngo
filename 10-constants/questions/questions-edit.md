# Questions: Constants #

---

## 1. What is a magic value? ##

(A) An unexpected value that pops up from somewhere.

(B) A spell cast by Merlin the Wizard.

(C) An unnamed constant value in the source code.

(D) A named constant.

---

**Solution**: (C)

A `magic value` is a constant value explicitly stated value in a program that can prevent an error from being detected until runtime, making it harder to understand.

---

## 2. What is a named constant? ##

(A) A constant with a cool name.

(B) A constant value with a declared name.

(C) A literal value converted to a name.

---

**Solution**: (C)

A named constant uses the syntax `const name = value`.

---

## 3. Which of the following choices shows the proper syntax for declaring a constant? ##

(A) `Const version int = 3`

(B) `const version int := 3`

(C) `const version int = 3`

---

**Solution**: (C)

the `const` keyword is case-sensitive so `Const` would not work.

---

## 4. Which of the following choices shows correct syntax? ##

(A) `s := "pick me"; const length = len(s)`

(B) `const message = "pick me!"; const length = len(message)`

(C) `const length = utf8.RuneCountInString("pick me")`

---

**Solution**: (B)

(A) `s` is not a `const` because it is missing the `const` keyword before its name.

(C) You cannot initialize constants using functions.

---

## 5. Which explanation below best explains the following code segment? ##

```go
const speed = 100
porsche := speed * 3
```

(A) `speed` is typeless and `porsche` is an `int`

(B) `speed` is an `int` and `porsche` is also an `int`

(C) `speed` and `porsche` are both typeless


---

**Solution**: (A)

That's right! `speed` is declared without a `type`, and `porsche` is an `int` because the expression (`speed * 3`) produces an integer value.

---

## 6. Which of the following choices fixes the error in the following code segment? ##

(A)
```go
const total = 10
x := 5

fmt.Print(total * x)
```

(B)
```go
const total = 10
x := 5

fmt.Print(total * x)
```

(C)
```go
const total int64 = 10
x := 5

fmt.Print(total * x)
```

(D)
```go
const total int64 = 10
x := 5

fmt.Print(int64(total) * x)
```

---

**Solution**: (A)

The constant `total` is typeless, so it can now be used with the variable `x`.

(B) There is still a  type mismatch since `x` is an `int` and not `int64`. 

(C) `total` is already an `int64`, so converting the type is redundant.

---

## 7. Which of the following choices fixes the error in the following code segment? ##

```go
const spell string
spell = "Abracadabra"
```

(A) `const spell = "Abracadabra"`

(B) `spell := "Abracadabra"`

(C) `var spell = "Abracadabra"`

---

**Solution**: (A)

A constant should always be initialized with an assigned value. Note that declaring a constant's `type` is not always necessary

The other choices are variables and not constants.

---

## 8. What are the values of the following constants? ##

```go
const (
    Yes = (iota * 5) + 2
    No
    Both
)
```

(A) `Yes=0`; `No=1`; `Both=2`

(B) `Yes=2`; `No=3`; `Both=4`

(C) `Yes=7`; `No=12`; `Both=17`

(D) `Yes=2`; `No=7`; `Both=12`


---

**Solution**: (C)

`iota` starts at 0 and not 1.