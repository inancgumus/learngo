# Slicing Quiz

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