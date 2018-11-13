# Arrays

## Where is the 2nd variable below stored in memory?
```go
// Let's say that first variable below is stored in this memory location: 20th
var (
    first  int32 = 100
    second int32 = 150
)
```
1. 21
2. 22
3. 24
4. It can be stored anywhere *CORRECT*

> **4:** That's right. It can be anywhere. Because, there's no guarantee that variables will be stored in contiguous memory locations.


## Where is the 3rd element of the following array stored in memory?

```go
//
// Let's say that:
// nums array is stored in this memory location (cell): 500th
//
// So, this means: nums[0] is stored at 500th location as well.
//
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
> **5:** Perfect. Array elements are stored in contiguous memory locations (cells). Here, the array's location is 500, and each element is 8 bytes (int64). So, 1st element: 500th, 2nd element: 508th, 3rd element: 516th, and so on. Formula: 516 = 500 + (8 * (3 - 1)).


## How many values the following variable represents?
```go
var gophers [10]string
```
1. 0
2. 1 *CORRECT*
3. 2
4. 10

> **2:** That's right! A variable can only store one value. Here, it stores a single array value with all its elements. However, through the gophers variable, you can access to 10 string values individually of that array.
>
> **4:** That's the length of the array. It's not the number of values that the gophers variable represents.


## ?
1. text *CORRECT*
2. text

> **1:**
>


