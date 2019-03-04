# Advanced Slice Operations Quiz

## What are the length and capacity of the 'part' slice?
```go
lyric := []string{"show", "me", "my", "silver", "lining"}
part  := lyric[1:3:5]
```
1. Length: 1 - Capacity: 5
2. Length: 1 - Capacity: 3
3. Length: 3 - Capacity: 5
4. Length: 2 - Capacity: 4 *CORRECT*

> **4:** General Formula: `[low:high:max]` => `length = high - max` and `capacity = max - low`. `lyric[1:3]` is `["me" "my"]`. `lyric[1:3:5]` is `["me" "my" "silver" "lining"]`. So, `[1:3]` is the returned slice, length: 2. `[1:3:5]` limits the capacity to four because after the 1st element there are only four more elements.


## What are the lengths and capacities of the slices below?
```go
lyric := []string{"show", "me", "my", "silver", "lining"}
part  := lyric[:2:2]
part   = append(part, "right", "place")
```
1. lyric's len: 5, cap: 5 — part's len: 5, cap: 5
2. lyric's len: 3, cap: 1 — part's len: 2, cap: 3
3. lyric's len: 5, cap: 5 — part's len: 4, cap: 4 *CORRECT*
4. lyric's len: 3, cap: 1 — part's len: 2, cap: 3

> **3:** `lyric[:2:2]` = ["show" "me"]. After the append the part becomes: ["show" "me" "right" "place"] — so it allocates a new backing array. `lyric` stays the same: `["show" "me" "my" "silver" "lining"]`.


## When you might want to use the make function?
1. To preallocate a backing array for a slice with a definite length *CORRECT*
2. To create a slice faster
3. To use less memory

> **1:** Yes! You can use the make function to preallocate a backing array for a slice upfront.


## What does the program print?
```go
tasks := make([]string, 2)
tasks  = append(tasks, "hello", "world")

fmt.Printf("%q\n", tasks)
```
1. ["" "" "hello" "world"] *CORRECT*
2. ["hello" "world"]
3. ["hello" "world" "" ""]

> **1:** `make([]string, 2)` creates a slice with len: 2 and cap: 2, and it sets all the elements to their zero-values. `append()` appends after the length of the slice (after the first two elements). That's why the first two elements are zero-valued strings but the last two elements are the newly appended elements.


## What does the program print?
```go
tasks := make([]string, 0, 2)
tasks  = append(tasks, "hello", "world")

fmt.Printf("%q\n", tasks)
```
1. ["" "" "hello" "world"]
2. ["hello" "world"] *CORRECT*
3. ["hello" "world" "" ""]

> **2:** `make([]string, 0, 2)` creates a slice with len: 0 and cap: 2. `append()` appends after the length of the slice (at the beginning). That's why the first two elements are overwritten with the newly appended elements. This is a common usage pattern when you want to use the `make` and the `append` functions together.


## What does the program print?
```go
lyric := []string{"le", "vent", "nous", "portera"}
n := copy(lyric, make([]string, 4))

fmt.Printf("%d %q\n", n, lyric)

// -- USEFUL INFORMATION (but not required to solve the question) --
//      In the following `copy` operation, `make` won't allocate
//      a slice with a new backing array up to 32768 bytes
//      (one string value is 8 bytes on a 64-bit machine).
//
//      This is an optimization made by the Go compiler.
```
1. 4 ["le" "vent" "le" "vent"]
2. 4 ["le" "vent" "nous" "portera"]
3. 4 ["" "" "" ""] *CORRECT*
4. 0 []

> **3:** `copy` copies a newly created slice with four elements (`make([]string, 4)`) onto `lyric` slice. They both have 4 elements, so the `copy` copies 4 elements. Remember: `make()` initializes a slice with zero-values of its element type. Here, this operation clears all the slice elements to their zero-values.


## What does the program print?
```go
spendings := [][]int{{200, 100}, {}, {50, 25, 75}, {500}}
total := spendings[2][1] + spendings[3][0] + spendings[0][0]

fmt.Printf("%d\n", total)
```
1. 725 *CORRECT*
2. 650
3. 500
4. 750

> **1:** `spendings[2][1]` = 25. `spendings[3][0]` = 500. `spendings[0][0]` = 200. 25 + 500 + 200 = 725


## What does the program print?
```go
spendings := [][]int{{1,2}}
    
// REMEMBER: %T prints the type of a given value
fmt.Printf("%T - ", spendings)
fmt.Printf("%T - ", spendings[0])
fmt.Printf("%T", spendings[0][0])
```
1. [][]int{{1, 2}} - []int{1, 2} - int(2)
2. [][]int - []int - int *CORRECT*
3. []int - int - 2
4. [][]int - [][]int - []int

> **2:** `spendings` is a 2-dimensional int slice, so its type is [][]int. Its element type is: `[]int`, so: `spendings[0]` is `[]int`. `spendings[0]`'s element type is: `int`. So `spendings[0][0]`'s type is `int`.


## What is the 'element type' of the slice?
```go
[][][3]int{{{10, 5, 9}}}
```    
1. [][][3]int
2. [][]int
3. [][3]int *CORRECT*
4. [3]int

> **3:** `[][][3]int` is a multi-dimensional slice of `[][3]int` elements. `[][3]int` is a multi-dimensional slice of `[3]int` elements. `[3]int` is an array of 3 `int` values.