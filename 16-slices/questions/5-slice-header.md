# Slice Header Quiz

## What is a slice header?
1. The first element of a slice value
2. The first element of the backing array
3. A tiny data structure that describes all or some part of a backing array *CORRECT*
4. A data structure that contains the elements of a slice

> **3:** Yes! It's just a tiny data structure with three numeric fields.
> 
> **4:** A slice doesn't contain any elements on its own.


## What are the fields of a slice value?
1. Pointer, length, and capacity *CORRECT*
2. Length and capacity
3. Only a pointer


## Which slice value does the following slice header describe?
SLICE HEADER:
+ Pointer : 100th
+ Length  : 5
+ Capacity: 10

Assume that the backing array is this one:
```go
var array [10]string
```
1. array[5:]
2. array[:5] *CORRECT*
3. array[3:]
4. array[100:]

> **1**: This slice's capacity is 5, it can only see the elements beginning with the 6th element.
> 
> **2**: That's right. `array[:5]` returns a slice with the first 5 elements of the `array` (len is 5), but there are 5 more elements in the backing array of that slice, so in total its capacity is 10.
> 
> **3**: This slice's capacity is 7, it can only see the elements beginning with the 4th element.
> 
> **4**: This is an error. The backing array doesn't have 100 elements.
> 


## Which one is the slice header of the following slice?
```go
var tasks []string
```
1. Pointer: 0, Length: 0, Capacity: 0 *CORRECT*
2. Pointer: 10, Length: 5, Capacity: 10
3. Pointer: 0, Length: 1, Capacity: 1

> **1:** A nil slice doesn't have backing array, so all the fields are equal to zero.


## What is the total memory usage of this code?
```go
var array [1000]int64

array2 := array
slice := array2[:]
```

1. 1024 bytes
2. 2024 bytes
3. 3000 bytes
4. 16024 bytes *CORRECT*

> **4:** `array` is 1000 x int64 (8 bytes) = 8000 bytes. Assigning an array copies all its elements, so `array2` adds additional 8000 bytes. A slice doesn't store anything on its own. Here, it's being created from array2, so it doesn't allocate a backing array as well. A slice header's size is 24 bytes. So in total: This program allocates 16024 bytes.


## What value does this code pass to the sort.Ints function?
```go
nums := []int{9, 7, 5, 3, 1}
sort.Ints(nums)
```
1. [9 7 5 3 1] — All the values of the nums slice
2. A pointer to the backing array of the nums slice
3. A pointer, length and capacity as three different arguments
4. The slice header that is stored in the nums variable *CORRECT*

> **1:** No, a slice value doesn't contain any elements. So it cannot pass the elements.
> 
> **2:** Sorry but not only that.
> 
> **3:** Nope. Remember, they are packed in a tiny data structure called the ....?
> 
> **4:** Yep! A slice value is a slice header (pointer, length and capacity). A slice variable stores the slice header.
> 

