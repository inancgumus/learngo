# Capacity and Append Mechanics Quiz

## What is the difference between the length and capacity of a slice?
1. They are the same
2. The length is always greater than the capacity
3. The capacity is always greater than the capacity
4. The length describes the length of a slice but a capacity describes the length of the backing array beginning from the first element of the slice *CORRECT*

> **2:** The length is never greater than the capacity.
> 
> **3:** The length and capacity of a slice can be equal.


## What is the capacity of a nil slice?
1. It is equal to its length + 1
2. It is nil
3. 0 *CORRECT*
4. 1

> **2:** The capacity's type is int, it cannot be nil.


## What are the length and capacity of the slice value?
```go
[]string{"I", "have", "a", "great", "capacity"}
```
1. Length: 5 - Capacity: 5 *CORRECT*
2. Length: 0 - Capacity: 5
3. Length: 5 - Capacity: 10
4. Length: 10 - Capacity: 10

> **1:** That's right! A slice literal creates a new slice value with equal length and capacity.


## What are the length and capacity of the 'words' slice?
```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = words[:0]
```
1. Length: 0 - Capacity: 0
2. Length: 6 - Capacity: 6
3. Length: 0 - Capacity: 6 *CORRECT*
4. Length: 5 - Capacity: 10

> **3:** Right! `words[:0]` slices for 0 elements, which in turn returns a slice with zero-length. Because the `words` slice points to the same backing array, its capacity is equal to 6.


## What are the length and capacity of the 'words' slice?
```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = words[0:]
```
1. Length: 0 - Capacity: 0
2. Length: 6 - Capacity: 6 *CORRECT* 
3. Length: 0 - Capacity: 6
4. Length: 5 - Capacity: 10

> **2:** Right! `words[0:]` slices for the rest of the elements, which in turn returns a slice with the same length as the original slice: 6. Beginning from the first array element, the `words` slice's backing array contains 6 elements; so its capacity is also 6.


## What are the length and capacity of the 'words' slice?
```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = words[2:cap(words)-2]
```
1. Length: 4 - Capacity: 6
2. Length: 6 - Capacity: 4
3. Length: 2 - Capacity: 6
4. Length: 2 - Capacity: 4 *CORRECT*

> **4:** Right! `words[2:cap(words)-2]` is equal to `words = words[2:4]`, so it returns: `["the" "sky"]`. So, its length is 2. But there are 4 more elements (`["the" "sky" "with" "diamonds"]`) in the backing array, so the capacity is 4.


