# Slice Basics Quiz

## How does the append function work?
1. It appends new elements to the given slice and returns a new slice, it doesn't change the given slice *CORRECT*
2. It appends new elements to the given slice and returns the existing slice
3. It appends new elements to the given slice and returns a new slice, it also changes the given slice

> **1:** Yes, the append function doesn't change the given slice unless you overwrite the result of the append function back to the original slice. 
>
> **2:** It doesn't return the existing slice, it returns a new slice. That's why you usually save the result of the append back to the original slice.
> 
> **3:** It doesn't change the given slice, it creates and returns a new slice.


## When you call the append function, where it does appends the new elements?
1. It appends them at the beginning of the given slice
2. It appends them at the middle of the given slice
3. It appends them after the length of the given slice *CORRECT*

> **3:** Yes! The append function appends the new elements by looking at the length of the given slice and then returns a new slice with the newly appended elements.


## What's the problem with the following code?
```go
nums := []int{9, 7, 5}
append(nums, []int{2, 4, 6}...)

fmt.Println(nums[3])
```
1. It can't append to an int slice
2. It can't append a slice to another slice
3. It gets an element that doesn't exist yet *CORRECT*

> **3:** That's right. The append function returns a new slice with the newly added elements. But here, it doesn't save the result of the append, so the nums slice only has 3 elements. So, `nums[3]` is invalid, it doesn't exist.


## What does this program print?
```go
nums := []int{9, 7, 5}
evens := append(nums, []int{2, 4, 6}...)

fmt.Println(nums, evens)
```
1. [9 7 5] [2 4 6]
2. [9 7 5] [9 7 5 2 4 6] *CORRECT*
3. [9 7 5 2 4 6] [2 4 6]
4. [9 7 5 2 4 6] [9 7 5 2 4 6]

> **2:** It appends the new elements to the nums slice but it saves the returned slice to the evens slice. So, the nums slice doesn't change. That's why, it prints the original elements from the nums slice first, then it prints the evens slice with the newly added elements.
>
> **3, 4:** It doesn't save the result of the append call back into the nums slice, so the nums slice doesn't contain the new elements.


## What does this program print?
```go
nums := []int{9, 7, 5}
nums = append(nums, 2, 4, 6)

fmt.Println(nums)
```
1. [9 7 5 2 4 6] *CORRECT*
2. [9 7 5]
3. [2 4 6]
4. [2 4 6 9 7 5]

> **1:** It overwrites the nums slice with the new slice that is returned from the append function. So the nums slice has the newly appended elements.


## What does this program print?
```go
nums := []int{9, 7, 5}
nums = append(nums, 2, 4, 6)

fmt.Println(nums[2:4])
```
1. [9 7 5 2 4 6]
2. [5 2] *CORRECT*
3. [4 6]
4. [7 2]

> **2:** nums is [9 7 5 2 4 6]. So, nums[2:4] is [5 2]. Remember, in nums[2:4] -> 2 is the starting index, so nums[2] is 5; And 4 is the stopping position, so nums[4-1] is 2 (-1 because the stopping position is the element position not an index). So, nums[2:4] returns a new slice that contains the elements at the middle of the nums slice.


## What does this program print?
```go
names := []string{"einstein", "rosen", "newton"}
fmt.Println(names[:])
```
1. [einstein rosen newton] *CORRECT*
2. [einstein rosen]
3. [einstein]
4. []

> **1:** The start index's default value is 0, and the stop position's default value is the length of the slice. So, `names[:]` equals to `names[0:3]`. It returns a new slice with the same elements.


## How can you get "rosen" element?
```go
names := []string{"einstein", "rosen", "newton"}
names2 := names[1:len(names) - 1]
```
1. names2[0] *CORRECT*
2. names2[1]
3. names2[2]

> **1:** That's right: names2 is ["rosen"] after the slicing.
> 
> **2:** That's not right. names is ["einstein" "rosen" "newton"] but names2 is ["rosen"] after the slicing. So, names2[1] is an error, it's because, the length of the names2 is 1.


## What does this program print?
```go
names := []string{"einstein", "rosen", "newton"}
names = names[1:]
names = names[1:]
fmt.Println(names)
```
1. [einstein rosen newton]
2. [rosen newton]
3. [newton] *CORRECT*
4. []

> **3:** Remember, slicing returns a new slice. Here, each slicing statement overwrites the names slice with the newly returned slice from the slicing. At first, the names was [einstein rosen newton. After the first slicing, the names becomes [rosen newton]. After the second slicing, names becomes [newton]. See this for the complete explanation: https://play.golang.org/p/EsEHrSeByFR