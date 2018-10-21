## Which one below is not one of the equality operators of Go?
1. `==`
2. `!=`
3. `>` *CORRECT*

> **3:** That's the greater operator. It checks whether an ordered value is greater than the other or not.


## Which one below is not one of the ordering operators of Go?
1. `>`
2. `<=`
3. `==` *CORRECT*
4. `<`

> **3:** That's the equal operator. In an expression, it checks whether a value (operand) is equal to another value (operand).


## Which one of these types is returned by the comparison operators?
1. int
2. byte
3. bool *CORRECT*
4. float64

> **3:** That's right. All the comparison operators return an untyped bool value (true or false).


## Which one of these below cannot be used as an operand to ordering operators (`>`, `<`, `>=`, `<=`)?
1. int value
2. byte value
3. string value
4. bool value *CORRECT*
5. all of them

> **1-2:** This is an ordered value, it can be used.
> 
> **3:** String is an ordered value because it's a series of numbers. So, it can be used as an operand.
> 
> **4:** That's right. A bool value is not an ordered value, so it cannot be used with ordering operators.


## Which one of these cannot be used as an operand to equality operators (`==`, `!=`)?
1. int value
2. byte value
3. string value
4. bool value
5. They all can be used *CORRECT*

> **5:** That's right. Every **comparable value** can be used as an operand to equality operators.


## What does this code print?
```go
fmt.Println("go" != "go!")
fmt.Println("go" == "go!")
```

1. true true
2. true false *CORRECT*
3. false true
4. false false
5. error

> **3-4:** Watch out for the exclamation mark at the end of the second string value.


## What does this code print?
```go
fmt.Println(1 == true)
```

1. true
2. 1
3. false
4. 2
5. error *CORRECT*

> **5:** That's right. A numeric constant cannot be compared to a bool value.


## What does this code print?
```go
fmt.Println(2.9 > 2.9)
fmt.Println(2.9 <= 2.9)
```

1. true true
2. true false
3. false true *CORRECT*
4. false false
5. error


## What does this code print?
```go
fmt.Println(false >= true)
fmt.Println(true <= false)
```

1. true true
2. true false
3. false true
4. false false
5. error *CORRECT*

> **5:** That's right. Bool values are not ordered values, so they cannot be compared using the comparison operators.


## How to fix this program without losing precision?
```go
package main
import "fmt"

func main() {
    weight, factor := 500, 1.5
    weight *= factor

    fmt.Println(weight)
}
```

1. It cannot be fixed
2. `weight *= float64(factor)`
3. `weight *= int(factor)`
4. `weight = float64(weight) * factor`
5. `weight = int(float64(weight) * factor)` *CORRECT*

> **1:** It can be fixed.
> 
> **2:** Type mismatch: weight is int.
>
> **3:** Lost precision: factor will be 1.
> 
> **4:** Type mismatch: weight is int (cannot assign back).
>
> **5:** That's right. The result would be 750.