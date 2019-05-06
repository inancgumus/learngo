## Why are maps used for?
For example, here is an inefficient program that uses a loop to find an element among millions of elements.
```go
millions := []int{/* millions of elements */}
for _, v := range millions {
    if v == userQuery {
        // do something
    }
}
```
1. Maps allow fast-lookup for map keys in O(1) time *CORRECT*
2. Maps allow fast-lookup for map keys in O(n) time
3. Maps allow fast-traversal on map keys in O(1) time

> **1:** That's right. Maps work in O(1) in average for fast-lookup.
>
> **2:** Map doesn't work in O(n) time for key lookup.
>

## When should you not use a map?
1. To find an element through a key
2. To loop over the map keys *CORRECT*
3. To add structured data to your program

> **1:** That is when you use a map.
> 
> **2:** Right! Looping over map keys happen in O(n) time. So, maps are the worst data structures for key traversing.
>
> **3:** Maps don't allow you to add structured data to your program.


## Which type below cannot be a map key?
1. map[string]int
2. []string
3. []int
4. []bool
5. All of them *CORRECT*

> **5:** Slices, maps, and function values are not comparable. So, they cannot be map keys.
>

## Which are the key and element types of the map below?
```go
map[string]map[int]bool
```
1. Key: string Element: bool
2. Key: string Element: int
3. Key: string Element: map[int]
4. Key: string Element: map[int]bool *CORRECT*

> **4:** The map contains other maps. The element type of a map can be of any type.
>

## What is a map value behind the scenes?
1. A map header
2. A pointer to a map header *CORRECT*
3. Tiny data structure with 3 fields: Pointer, Length and Capacity

> **2**: That's right. Maps are complex data structures. However, each map value is only a pointer to a map header (which is a more complex data structure).