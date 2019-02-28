full slice exp
make
copy

# Slice Internals Quiz

## Where does a slice value stores its elements?
1. In the global backing array that is shared by all the slices
2. In a separate backing array that a slice value points to *CORRECT*
3. In the slice value itself (like an array)

> **1:** There isn't something called the global backing array.
> 
> **2:** That's right. A slice doesn't store any elements direcly. It points to an array that is stored separately on the computer memory, and a slice points to that array, a slice is just a window to that array.
> 
> **3:** A slice values doesn't store any elements by itself. It's just a description of its backing array.


## What's a backing array?
1. An array of trustworthy friends that you can count on always
2. An array that is stored in the slice value
3. A slice value stores its elements in it *CORRECT*

> **1:** Oh, come on!
> 
> **2:** Nope, a slice value doesn't store its backing array, the slice value just points to it. The backing array is stored separately from the slice value.


## What does this program print?
```go
nums := []int{10, 15, 10, 2, 3, 4}
digits := nums[len(nums)-3:]

nums[len(nums)-1] = 8
digits[0] += nums[3]

fmt.Println(digits)
```
1. [2 3 4]
2. [4 3 4]
3. [4 3 8] *CORRECT*
4. [10 15 10]
5. [4 15 8]

> **3:** Awesome! At first, digits is [2 3 4]. After `nums[len(nums)-1] = 8`, the digits becomes [2 3 8] (it's because, digits is created by slicing the nums slice, so they share the same backing array). And lastly, after `digits[0] += nums[3]`, the digits becomes [4 3 8].


## Find the correct explanation below
1. A slice variable can store a slice value and the slice value is actually a slice header behind the scenes *CORRECT*
2. A slice variable can only store the same slice value and it cannot be changed afterward
3. A slice header stores the slice value

> **1:** That's right. You can change the slice values that is being stored in a slice variable, and the slice value is actually a slice header behind the scenes.
> 
> **3:** Nope, actually a slice header and a slice value is the same thing.


* after slicing, does it copy all the values to the new slice?
* why indexing a slice is fast? (array is contagious on memory)
* separate backing arrays question
* sliced array share the same explicit array

## 
backing array
slice header
capacity
append mechanics