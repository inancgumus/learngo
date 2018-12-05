# Arrays Quiz

## What's the length of this array literal?
```go
gadgets := [...]string{"Mighty Mouse", "Amazing Keyboard", "Shiny Monitor"}
```
1. 0
2. 1
3. 2
4. 3 *CORRECT*

> **4:** Yes! There are 3 elements in the element list. So, Go sets the length of the array to 3.
>


## What's the type and length of this array literal?
```go
gadgets := [...]string{}
```
1. [0]string and 0 *CORRECT*
2. [0]string{} and 0
3. [1]string and 1 
4. [1]string{} and 1 

> **1:** Yes! There are no elements in the element list. So, Go sets the length of the array to 0.
>

## What does this program print?
```go
package main
import "fmt"

func main() {
	gadgets := [3]string{"Confused Drone"}
	fmt.Printf("%q\n", gadgets)
}
```
1. [3]string{"Confused Drone", "", ""}
2. [1]string{"Confused Drone"}
3. ["Confused Drone" "" ""] *CORRECT*
4. ["Confused Drone"]

> **1:** %q verb doesn't print the type of an array.
>
> **2, 4:** Array's length cannot change depending on the elements.
>
> **3:** Yes! Go sets the uninitialized elements to their zero values.
>


## Are these arrays comparable?
```go
gadgets := [3]string{"Confused Drone"}
gears   := [...]string{"Confused Drone"}

fmt.Println(gadgets == gears)
```
1. Yes, because they have identical types and elements
2. No, because their types are different *CORRECT*
3. No, because their elements are different

> **2:** Yes! gadget's type is [3]string whereas gears's type is [1]string.
>


## What does this program print?
```go
gadgets := [3]string{"Confused Drone", "Broken Phone"}
gears   := gadgets

gears[2] = "Shiny Mouse"

fmt.Printf("%q\n", gadgets)
```
1. ["Confused Drone" "Broken Phone" "Shiny Mouse"]
2. ["Confused Drone" "Broken Phone" ""] *CORRECT*
3. ["" "" "Shiny Mouse"]
4. ["" "" ""]

> **2:** Yes! When you assign an array, Go creates a copy of the original array. So, gadgets and gears arrays are not connected. Changing one of them won't effect the other one.
>


## What's the type of the digits array?
```go
digits := [...][5]string{
	{
		"## ",
		" # ",
		" # ",
		" # ",
		"###",
	},
	[5]string{
		"###",
		"  #",
		"###",
		"  #",
		"###",
	},
}
```
1. [...][5]string
2. [2][2]string
3. [2][5]string *CORRECT*
4. [5][5]string

> **3:** Awesome! There are two inner arrays, so the outer array's length becomes 2. Also note that, `[5]string` in front of the second element is unnecessary.
>


## What does this program print?
```go
rates := [...]float64{
    5: 1.5,
    2.5,
    0: 0.5,
}

fmt.Printf("%#v\n", rates)
```
1. [7]float64{0.5, 0, 0, 0, 0, 1.5, 2.5} *CORRECT*
2. [7]float64{1.5, 2.5, 0.5, 0, 0, 0, 0}
3. [3]float64{1.5, 2.5, 0.5}
4. [3]float64{0.5, 2.5, 1.5}

> **1:** That's right! For the explanation check out the example in the course repository here: https://github.com/inancgumus/learngo/tree/master/14-arrays/11-keyed-elements/06-keyed-and-unkeyed
>



## Are these arrays equal?
```go
type three [3]int

nums  := [3]int{1, 2, 3}
nums2 := three{1, 2, 3}

fmt.Println(nums == nums2)
```
**Note:** _To solve this question you need to watch the comparison and unnamed types lectures._
1. Yes, because they have identical underlying types and elements *CORRECT*
2. No, because their types are different
3. No, because their length is different

> **1:** Yes! They both have the same underlying types: [3]int
>


## Are these array variables equal?
```go
type (
    threeA [3]int
    threeB [3]int
)

nums  := threeA{1, 2, 3}
nums2 := threeA(threeB{1, 2, 3})

fmt.Println(nums == nums2)
```
**Note:** _To solve this question you need to the watch comparison and unnamed types lectures._
1. Yes, because they have identical underlying types and elements *CORRECT*
2. No, because their types are different
3. No, because their length is different

> **1:** Yes! Actually, arrays have different types, so normally they're not comparable. However, when you convert `threeB{1, 2, 3}` array to `threeA` type, they become comparable.
>