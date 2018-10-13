## Which one is a correct type conversion expression?
* convert(40)
* var("hi")
* int(4.) *CORRECT*
* int[4]

## What does this code print?
```go
age := 6.5
fmt.Print(int(age))
```
* 6.5
* 65
* 6 *CORRECT*
* .5

> When you convert a float to integer
> It loses its fractional part

## What does this code print?
```go
fmt.Print(int(6.5))
```
* 6.5
* 65
* 6
* Compile-Time Error *CORRECT*

> Go can detect conversion errors at the compile-time

## What does this code print?
```go
area := 10.5
fmt.Print(area/2)
```
* 5.25 *CORRECT*
* 5 
* 0
* Error

## What does this code print?
```go
area := 10.5
div := 2
fmt.Print(area/div)
```
* 5.25
* 5
* ERROR *CORRECT*

> You can't divide different type of values.
> You need to convert: `area / float64(div)`

## Which code below can fix the following code?
```go
area := 10.5
div := 2
fmt.Print(area/div)
```
* `fmt.Print(int(area)/div)`      // 5
* `fmt.Print(area/int(div))`      // type mismatch
* `fmt.Print(int(area)/int(div))` // 5
* `fmt.Print(area/float64(div))`  *CORRECT*