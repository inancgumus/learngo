# Array Basics Quiz

## What's an array?
1. Accelerated charged particle array gun from Star Wars
2. Collection of values with dynamic length and type
3. Collection of values with fixed length and type *CORRECT*


## Where is the 2nd variable below stored in memory?
```go
// Let's say that the first variable below is stored in this memory location: 20th
var (
    first  int32 = 100
    second int32 = 150
)
```
1. 21
2. 22
3. 24
4. It can be stored anywhere *CORRECT*

> **4:** That's right. It can be anywhere. Because, unlike arrays, there isn't any guarantee that variables will be stored in contiguous memory locations.


## Where is the 3rd element of the nums array stored in memory?

```go
// Let's say: nums array is stored in this memory location: 500th
var nums [5]int64
```
1. 3
2. 2
3. 502
4. 503
5. 516 *CORRECT*

> **2:** Nope, that's the index of an element.
>
> **3, 4:** 500+index? You're getting closer.
>
> **5:** Perfect. Array elements are stored in contiguous memory locations. Here, the array's location is 500, and each element is 8 bytes (int64). So, 1st element is stored in: 500th, 2nd element: 508th, 3rd element: 516th, and so on. Formula for math lovers: 516 = 500 + 8 * (3 - 1). General formula: Starting position + The size of each element * (Element's index position - 1).


## How many values this variable stores?
**HINT:** _I'm asking about the variable not about the array inside the variable._
```go
var gophers [10]string
```
1. 0
2. 1 *CORRECT*
3. 2
4. 10

> **2:** That's right! A variable can only store one value. Here, it stores a single array value. However, through the variable, you can also access to individual string values inside the array value.
>
> **4:** That's the length of the array. It's not the number of values that the gophers variable stores.


## What's the length of this array?
```go
var gophers [5]int
```
1. 5 *CORRECT*
2. 1
3. 2

> **1:** That's right! It stores 5 int values.
>

## What's the length of this array?
```go
const length = 5 * 2
var gophers [length - 1]int
```
1. 10
2. 9 *CORRECT*
3. 1

> **2:** That's right! 5 * 2 - 1 is 9. You can use constant expressions while declaring the length of an array.


## What's the element type of this array?
```go
var luminosity [100]float32
```
1. [100]float32
2. luminosity
3. float32 *CORRECT*


## What's the type of this array?
```go
var luminosity [100]float32
```
1. [100]float32 *CORRECT*
2. luminosity
3. float32

> **1:** That's right. Array's type is consisting of its length and its element type together.
>


## What does this program print?
```go
package main
import "fmt"

func main() {
	var names [3]string

	names[len(names)-1] = "!"
	names[1] = "think" + names[2]
	names[0] = "Don't"
	names[0] += " "

	fmt.Println(names[0] + names[1] + names[2])
}
```
1. !think!Don't
2. Don't think!! *CORRECT*
3. This program is incorrect

> **2:** "Don't think!! Just do!". Explanation is here: https://play.golang.org/p/y_Tqwn_XRlg
>


## What does this program print?
_It's OK if you can't solve this question. This is a hard question. You may try it on Go Playground here: https://play.golang.org/p/o0o0UM7Ktyy_
```go
package main
import "fmt"

// This program sets the next element of the array,
// then it quits just before the last element.
func main() {
	var sum [5]int

	for i, v := range sum {
	    if i == len(sum) - 1 {
	        break
	    }

	    sum[i+1] = 10
	    fmt.Print(v, " ")
	}
}
```
1. 0 0 0 0 *CORRECT*
2. 10 10 10 10
3. 0 10 10 10

> **1:** That's right! the for range clause copies the sum array. So, changing the elements of the sum array while inside of the loop won't effect the original sum array. However, if you try to print it right after the loop, you'll see that it has changed. Try printing it like so on Go Playground.
>
> **2:** That's not right. Because, the for range clause copies the sum array.
>

