## Why you want to define new types?
1. To declare new methods
2. For type-safety
3. For readability and conveying meaning
4. All of the options above *CORRECT*

> **1-3:** Yes, that's only one of the reasons.
>


## Let's suppose that you've declared the following defined type. Which property below the new type doesn't get from its underlying type?

```go
// For example, let's say that you've defined a new type
// using time.Duration type like this:
type Millennium time.Duration
```
1. Methods *CORRECT*
2. Representation
3. Size
4. Range of values

> **1:** That's right. A defined type doesn't get its source type's methods.
>
> **2-4:** Actually the defined type gets it from its underlying type.
>


## How to define a new type using float32?
1. `var radius float32`
2. `radius = type float32`
3. `type radius float32` *CORRECT*
4. `type radius = float32`

> **1-2:** This is not a correct syntax.
>
> **3:** `radius` is a new type, defined using `float32`.
>
> **4:** This declares `radius` as an alias to `float32`. So, they become the same types.
>


## How to fix the following code?
```go
type Distance int

var (
    village Distance = 50
    city = 100
)

fmt.Print(village + city)
```
1. `int(village + city)`
2. `village + int(city)`
3. `village(int) + city`
4. `village + Distance(city)` *CORRECT*

> **1-3:** There's a type mismatch in this code. But, this won't fix it.
>
> **4:** That's right. Now, the `city`'s type is Distance in the expression.
>


## For the following program which types you might want to declare?
```go
package main
import "fmt"

func main() {
	celsius := 35.
	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%g ºC is %g ºF\n", celsius, fahrenheit)
}
```

1. Celsius and Fahrenheit using int64
2. Celsius and Fahrenheit using float64 *CORRECT*
3. Temperature using int64
4. Temperature using uint32

> **1:** But a degree value has a floating part. So, using an integer may not the best way.
>
> **2:** float64 can represent a degree value.
>
> **3-4:** But a degree value has a floating part. So, using an integer may not the best way. Also, there are two different temperature units here: Celsius and Fahrenheit. Isn't it better to create two distinct types?
>


## What's the underlying type of the Millennium type?
```go
type (
    Duration int64
    Century Duration
    Millennium Century
)
```
1. `int64` *CORRECT*
2. `Duration`
3. `Century`
4. Another type

> **1:** That's right. Go's type system is flat. So, the defined type's underlying type is a type with a real structure. int64 is not just a name, it has its own structure, it's a predeclared type.
>
> **2:** Duration is just a new type name. It doesn't have its own structure.
>
> **3:** Century is just a new type name. It doesn't have its own structure.
>


## Which types do not need to be converted between each other?
**HINT:** Aliased types do not require type conversion.

1. byte and uint8 *CORRECT*
2. byte and rune
3. rune and uint8
4. byte and uint32

> **1:** byte data type is an alias to uint8 data type. So, they don't need conversion between each other. They're the same types.
>


## Which types do not need to be converted between each other?
**HINT:** Aliased types do not require type conversion.

1. byte and uint32
2. byte and rune
3. rune and int32 *CORRECT*
4. byte and int8

> **3:** rune data type is an alias to int32 data type. So, they don't need conversion between each other. They're the same types.
>