## What's a magic value?
1. A value which pops up from somewhere
2. Merlin the Wizard's spell
3. An unnamed constant value in the source code *CORRECT*
4. A named constant


## What's a named constant?
1. A constant with a cool name
2. A constant value declared with a name *CORRECT*
3. A literal value converted to a name


## How to declare a constant?
1. `Const version int = 3`
2. `const version int := 3`
2. `const version int = 3` *CORRECT*

> **1:** "C"onst should be just "c"onst.
>


## Which code below is correct?
1. `s := "pick me"; const length = len(s)`
2. `const message = "pick me!"; const length = len(message)` *CORRECT*
3. `const length = utf8.RuneCountInString("pick me")`

> **1:** `s` not a constant.
>
> **2:** `len` function can be used as an initial value to a constant, when the argument to `len` is also a constant.
>
> **3:** You cannot call functions while initializing a constant.
>


## Which explanation below is correct for the following code?
```go
const speed = 100
porsche := speed * 3
```
1. speed is typeless and porsche's type is int *CORRECT*
2. speed's type is int and porsche's type is also int
3. speed and porsche are typeless

> **2:** speed has no type.
>
> **3:** A variable cannot be typeless.
>


## How to fix the following code?
```go
const spell string
spell = "Abracadabra"
```
1. `const spell = "Abracadabra"` *CORRECT*
2. `spell := "Abracadabra"`
3. `var spell = "Abracadabra"`

> **1:** A constant always have to be initialized to a value. And, sometimes the type declaration is not necessary.
>
> **2-3:** That's a variable not a constant.
>


## How to fix the following code?
```go
const total int8 = 10
x := 5

fmt.Print(total * x)
```

```go
// #1 - *CORRECT*
const total = 10
x := 5

fmt.Print(total * x)

// #2
const total int64 = 10
x := 5

fmt.Print(total * x)

// #3
const total int64 = 10
x := 5

fmt.Print(int64(total) * x)
```

> **1:** Now, the total constant is typeless, so it can be used with the x variable.
>
> **2:** There's still a type mismatch. x is int not int64.
>
> **3:** total is already int64. No need to convert it again.
>


## What are the values of the following constants?
```go
const (
    Yes = (iota * 5) + 2
    No
    Both
)
```
1. Yes=0 No=1 Both=2
2. Yes=2 No=3 Both=4
3. Yes=7 No=12 Both=17
4. Yes=2 No=7 Both=12 *CORRECT*

> **3:** iota starts at 0, not 1.
>