# The Mechanics of Append Quiz

## Which append call below allocates a new backing array for the following slice?
```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
```
1. words = append(words[:3], "crystals")
2. words = append(words[:4], "crystals")
3. words = append(words[:5], "crystals")
4. words = append(words[:5], "crystals", "and", "diamonds") *CORRECT*

> **1:** No, it just overwrites the 4th element.
> 
> **2:** No, it just overwrites the 5th element.
> 
> **3:** No, it just overwrites the last element.
> 
> **4:** Yes, it overwrites the last element, then it adds two element. However, there is not enough space to do that, so it allocates a new backing array.


## What does the program print?
```go
words := []string{"lucy", "in", "the", "sky", "with", "diamonds"}
words = append(words[:1], "is", "everywhere")
words = append(words, words[len(words)+1:cap(words)]...)
```
1. lucy in the sky with diamonds
2. lucy is everywhere in the sky with diamonds
3. lucy is everywhere with diamonds *CORRECT*
4. lucy is everywhere

> **3:** line #2 overwrites the 2nd and 3rd elements. line #3 appends ["with" "diamonds"] after the ["lucy" "is" "everwhere"].


## What are the length and capacity of the words slice?
```go
// The words slice has 1023 elements.
//
// Tip: The keyed slice works like the same as a keyed array.
// If you don't remember how it works, please check out the keyed elements in the arrays section.
//
words := []string{1022: ""}
words = append(words, "boom!")
```
1. Length: 1024 - Capacity: 1024
2. Length: 1025 - Capacity: 1025
3. Length: 1025 - Capacity: 1280
4. Length: 1024 - Capacity: 2048 *CORRECT*

> **4:** That's right! Append function grows by doubling the capacity of the previous slice.


## What are the length and capacity of the words slice?
```go
// The words slice has 1024 elements.
//
// Tip: The keyed slice works like the same as a keyed array.
// If you don't remember how it works, please check out the keyed elements in the arrays section.
//
words := []string{1023: ""}
words = append(words, "boom!")
```
1. Length: 1024 - Capacity: 1024
2. Length: 1025 - Capacity: 1025
3. Length: 1025 - Capacity: 1280 *CORRECT*
4. Length: 1025 - Capacity: 2048

> **3, 4:** After 1024 elements, the append function grows at a slower rate, about 25%.