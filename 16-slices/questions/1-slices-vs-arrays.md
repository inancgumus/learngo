# Slices vs Arrays Quiz

## Why you want to use a slice instead of an array?
1. I like arrays more
2. I want to create a dynamic collection, so I need an array
3. A slice's length is dynamic, so I can create dynamic collections *CORRECT*


## Where does the length of a slice belong to?
1. Compile-Time
2. Runtime *CORRECT*
3. Walk-Time
4. Sleep-Time

> **2:** A slice's length is not a part of its type. So its length can change at runtime.


## Which function call below is correct?
```go
// Let's say there's a function like this.
func sort(nums []int) {
    // ...
}
```
1. sort([...]int{3, 1, 6})
2. sort([]int32{3, 1, 6})
3. sort([]int{3, 1, 6}) *CORRECT*

> **1:** You can't call the sort function using an array. It expects an int slice.
> 
> **2:** You can't call the sort function using an int32 slice. It expects an int slice.
> 
> **3:** That's right! You can pass an int slice to the sort function.


## What is the zero value of this slice?
```go
var tasks []string
```
1. 0
2. 1
3. nil *CORRECT*
4. unknown

> **3:** This is a nil slice. Unlike an array, a slice's zero value is nil.


## What does this code print?
```go
var tasks []string
fmt.Println(len(tasks))
```

1. 0 *CORRECT*
2. 1
3. nil
4. It doesn't work.

> **1:** Yes, you can use the len function on a nil slice. It returns 0 because the slice doesn't contain any elements yet.


## What does this code print?
```go
var tasks []string
fmt.Println(tasks[0])
```

1. 0
2. 1
3. nil
4. It doesn't work. *CORRECT*

> **4:** You can't get an element that does not exist. A nil slice does not contain any elements.


## Which declaration below is a correct slice declaration?
1. [...]int{}
2. [2]string{"hello", "world"}
3. []string{"hello", "world"} *CORRECT*
4. string[2]{"hello", world"}


## This code doesn't work, why?
```go
colors := []string{"red", "blue", "green"}
tones := []string{"dark", "light"}

if colors == tones {
    // ...
}
```

1. The slices have different lengths
2. If statement doesn't contain any statements
3. Slices cannot be compared *CORRECT*

> **3:** That's right! A slice value can only be compared to a nil value.


## What is the length of this slice?
```go
[]uint64{}
```

1. 64
2. 1
3. 0 *CORRECT*
4. Error

> **3:** That's right. This is an empty slice, it doesn't contain any elements.


## What is the length of this slice?
```go
[]string{"I'm", "going", "to", "stay", "\"here\""}
```

1. 0
2. 1
3. 2
4. 3
5. 4
6. 5 *CORRECT*
