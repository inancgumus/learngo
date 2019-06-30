## What is a pointer?
1. A variable that contains an hexadecimal value
2. A variable that contains a memory address
3. A value that can contain a memory address of a value *CORRECT*
4. A value that points to a function

> **2:** Although a pointer can be put into a variable, it's not solely a variable. You're almost there! But this distinction is very important.
> 
> **3:** A pointer is just another value that can contain a memory address of a value.


## Which one is a pointer to a computer?
```go
type computer struct {
	brand string
}
```
1. `*computer{}`
2. `var c computer`
3. `var *c computer`
4. `var c *computer` *CORRECT*

> **4:** * in front of a type denotes a pointer type.


## Which one gets the pointed composite value by the following pointer?
```go
type computer struct {
	brand string
}

c := &computer{"Apple"}
```
1. `*c` *CORRECT*
2. `&c`
3. `c`
4. `*computer`

> **1:** * in front of a pointer value gets the value that is pointed by the pointer.
> 
> **2:** & in front of a value gets the memory address of that value
> 
> **4:** * in front of a type denotes a pointer type.


## What is the result of the following code?
```go
type computer struct {
    brand string
}

var a, b *computer
fmt.Print(a == b)

a = &computer{"Apple"}
b = &computer{"Apple"}
fmt.Print(" ", a == b, " ", *a == *b)
```
1. false false false
2. true true true
3. true false true *CORRECT*
4. false true true

> **3:** a and b are nil at the beginning, so they are equal. However, after that, they get two different memory addresses from the composite literals, so their addresses are not equal but their values (that are pointed by the pointers) are equal.


## How many variables are there in the following code?
```go
type computer struct {
    brand string
}

func main() {
    a = &computer{"Apple"}
    b := a
    change(b)
    change(b)
}

func change(c *computer) {
    c.brand = "Indie"
    c = nil
}
```
1. 1
2. 2
3. 3
4. 4 *CORRECT*

> **4:** Every time a func runs, it creates new variables from its input params and named result values (if any). There two pointer variables: a and b. Then there are, two more pointer variables because: change is called two times.


## Why you cannot take the address of a map's elements?
1. It's an addressable value
2. It's an unaddressable value *CORRECT*
3. Doing so can crash your program

> **2:** Map elements are not addressable, so you cannot take their addresses.