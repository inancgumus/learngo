# Backing Array Quiz

## Where does a slice store its elements?
1. In the slice value
2. In a global backing array that is shared by all the slices
3. In a backing array that is specific to a slice
4. In a backing array that the slice references *CORRECT*

> **1:** A slice value doesn't store any elements. It's just a simple data structure.
> 
> **2:** There is not a global backing array.
> 
> **3:** A backing array can be shared among slices. It may not be specific to a slice.
> 
> **4:** Yep! A slice stores its elements in a backing that the slice references (or points to).
> 


## When you slice a slice, what value does it return?
```go
// example:
s := []string{"I'm", "a", "slice"}
s[2:] // <-- slicing
```
1. It returns a new slice value with a new backing array
2. It returns the existing slice value with a new backing array
3. It returns a new slice value with the same backing array *CORRECT*

> **3:** Yes! Slicing returns a new slice that references to some segment of the same backing array.


## Why are slicing and indexing a slice efficient?
1. Slices are fast
2. Backing arrays are contiguous in memory *CORRECT*
3. Go uses clever algorithms

> **2:** Yes. A slice's backing array is contiguous in memory. So, accessing an element of a slice is very fast. Go can look at a specific memory location to find an element's value very fast.


## Which one is the backing array of "slice2"?
```go
arr := [...]int{1, 2, 3}
slice1 := arr[2:3]
slice2 := slice1[:1]
```

1. arr *CORRECT*
2. slice1
3. slice2
4. A hidden backing array

> **1:** Yes! When a slice is created by slicing an array, that array becomes the backing array of that slice.
> 
> **4:** Nope. That only happens when a slice doesn't being created from an array.
>


## Which one is the backing array of "slice"?
```go
arr := [...]int{1, 2, 3}
slice := []int{1, 2, 3}
```

1. arr
2. slice1
3. slice2
4. A hidden backing array *CORRECT*

> **1:** Nope, the slice hasn't created by slicing an array.
> 
> **4:** Yes! A slice literal always creates a new hidden array.
>


## Which answer is correct for the following slices?
```go
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
```
1. Their backing array is the same.
2. Their backing arrays are different. *CORRECT*
3. They don't have any backing arrays.

> **2:** That's right. A slice literal always creates a new backing array.


## Which answer is correct for the following slices?
```go
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
slice3 := slice1[:]
slice4 := slice2[:]
```
1. slice1 and slice2 have the same backing arrays.
2. slice1 and slice3 have the same backing arrays. *CORRECT*
3. slice1 and slice4 have the same backing arrays.
4. slice3 and slice4 have the same backing arrays.

> **2:** Yep! A slice that is being created by slicing shares the same backing with the sliced slice. Here, slice3 is being created from slice1. That is also true for slice2 and slice4.


## What does the backing array of the nums slice look like?
```go
nums := []int{9, 7, 5, 3, 1}
nums = nums[:1]
	
fmt.Println(nums) // prints: [9]
```
1. [9 7 5 3 1] *CORRECT*
2. [7 5 3 1]
3. [9]
4. []


## What does this code print?
```go
arr   := [...]int{9, 7, 5, 3, 1}
nums  := arr[2:]
nums2 := nums[1:]

arr[2]++
nums[1]  -= arr[4] - 4
nums2[1] += 5

fmt.Println(nums)
```
1. [5 3 1]
2. [6 6 6] *CORRECT*
3. [9 7 5]

> **2:** Yes! Because the backing array of `nums` and `nums2` is the same: `arr`. See the explanation here: https://play.golang.org/p/xTy0W0S_8PN
