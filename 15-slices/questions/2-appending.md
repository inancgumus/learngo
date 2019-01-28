# Appending Quiz

## How does the append function work?
1. It appends new elements to the given slice on the fly and returns a new slice, normally, it doesn't change the given slice *CORRECT*
2. It appends new elements to the given slice and returns the existing slice
3. It appends new elements to the given slice and returns a new slice, it also changes the given slice

> **1:** Yes, the append function doesn't change the given slice unless you overwrite the result of the append function back to the original slice.  Most of the times, this is true.
>
> **2:** It doesn't return the existing slice, it returns a new slice. That's why you usually save the result of the append back to the original slice.
> 
> **3:** It doesn't change the given slice, it creates and returns a new slice. Most of the times, this is true.


## When you call the append function, where does it append the new elements?
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
3. It tries to get an element that doesn't exist yet *CORRECT*

> **3:** That's right. The append function returns a new slice with the newly added elements. But here, the code doesn't save the result of the append, so the nums slice only has 3 elements. So, `nums[3]` is invalid, because there are only 3 elements in the nums slice.


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

> **1:** It overwrites the nums slice with the new slice that is returned from the append function. So the nums slice has got the newly appended elements.