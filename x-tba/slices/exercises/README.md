# Slice Exercises

---

# ANNOUNCEMENT

I teach you what the other courses don't even care to teach. 

**What's new?**
* New Section: Advanced Slice Operations
* New Exercises for the Slice Internals
* New Exercises for the Slices: Advanced Operations

**What are you going to learn?**
* Full Slice Expressions: Limiting access to the backing array
* Make(): Preallocation
* Copy(): Efficiently and safely copy elements without using a loop
* Multi-Dimensional Slices

**What's coming next?**
* Empty Filer Finder: Your first taste of file operations.
* Bouncing Ball: Create a bouncing ball animation on a 2D surface.
* Png Parser: Parse a PNG file by hand and tell its dimensions.

These lectures will be added in the next 3 weeks.

**Statistics:**
* +1 hour of additional content!
* +5 new lectures
* +20 new questions
* 3 + ? new exercises

**Total content in the slices section:**
* ? hours
* ? lectures
* ? questions
* ? exercises

---

## Full Slice Exp + Make + Copy + Multi-Dim Slices

# FIX THIS
1. **[Limit the backing array sharing](https://github.com/inancgumus/learngo/tree/master/16-slices/exercises/??-limit-the-backing-array-sharing)**

* fix the excess memory allocation
  * return a huge slice from a func, ask the student fix it
* full slice exp: https://play.golang.org/p/SPrLspRBXdI
* copy: https://play.golang.org/p/SPrLspRBXdI
  * + put \n for the beatles exercise using copy
  ```go
    s = append(s, 0 /* use the zero value of the element type */)
    copy(s[i+1:], s[i:])
    s[i] = x
  ```


* multi dim slices batches
```go
actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
batchSize := 3
var batches [][]int

for batchSize < len(actions) {
    actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
}
batches = append(batches, actions)
```