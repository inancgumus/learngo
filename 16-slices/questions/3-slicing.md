# Slicing Quiz

## What does this code print?
```go
nums := []int{9, 7, 5}
nums = append(nums, 2, 4, 6)

fmt.Println(nums[2:4])
```
1. [9 7 5 2 4 6]
2. [5 2] *CORRECT*
3. [4 6]
4. [7 2]
5. [9 7]

> **2:** nums is [9 7 5 2 4 6]. So, nums[2:4] is [5 2]. Remember, in nums[2:4] -> 2 is the starting index, so nums[2] is 5; And 4 is the stopping position, so nums[4-1] is 2 (-1 because the stopping position is the element position). So, nums[2:4] returns a new slice that contains the elements at the middle of the nums slice.


## What does this code print?
```go
nums := []int{9, 7, 5}
nums = append(nums, 2, 4, 6)

fmt.Println(nums[:2])
```
1. [9 7 5 2 4 6]
2. [5 2]
3. [4 6]
4. [7 2]
5. [9 7] *CORRECT*

> **5:** nums is [9 7 5 2 4 6]. So, nums[:2] is nums[0:2] which in turn returns [9 7].


## What does this code print?
```go
nums := []int{9, 7, 5}
nums = append(nums, 2, 4, 6)

fmt.Println(nums[len(nums)-2:])
```
1. [9 7 5 2 4 6]
2. [5 2]
3. [4 6] *CORRECT*
4. [7 2]
5. [9 7]

> **3:** nums is [9 7 5 2 4 6]. So, nums[len(nums)-2:] is nums[4:6] (len(nums) is 6) which in turn returns [4 6].


## What does this code print?
```go
names := []string{"einstein", "rosen", "newton"}
names = names[:]
fmt.Println(names[:1])
```
1. [einstein rosen newton]
2. [einstein rosen]
3. [einstein] *CORRECT*
4. []

> **3:** names[:] is names[0:3] -> [einstein rosen newton]. names[:1] is names[0:1] -> [einstein].


## What is the type of the marked expression below?
```go
names := []string{"einstein", "rosen", "newton"}
names[2:3] // <- marked
```

1. []string *CORRECT*
2. string
3. names
4. []int

> **1:** Yes! A slicing expression returns a slice.
> 
> **2:** Remember, a slicing expression returns a slice. Did I give you the answer? Oops.


## What is the type of the marked expression below?
```go
names := []string{"einstein", "rosen", "newton"}
names[2] // <- marked
```

1. []string
2. string *CORRECT*
3. names
4. []int

> **1:** Remember, an index expression returns an element value, not a slice.
> 
> **2:** Yep! An index expression returns an element value. The element type of the []string slice is string, so the returned value is a string value.


## Which index expression returns the "rosen" element?
```go
names := []string{"einstein", "rosen", "newton"}
names = names[1:len(names) - 1]
```
1. names[0] *CORRECT*
2. names[1]
3. names[2]

> **1:** That's right: names2 is ["rosen"] after the slicing.
> 
> **2:** That's not right. Remember, indexes are relative to a slice. names is ["einstein" "rosen" "newton"] but names[1:len(names)-1] is ["rosen"]. So, names2[1] is an error, it's because, the length of the last slice is 1.


## What does this code print?
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

> **3:** Remember, slicing returns a new slice. Here, each `names = names[1:]` statement overwrites the names slice with the newly returned slice from the slicing. At first, the names was [einstein rosen newton]. After the first slicing, the names becomes [rosen newton]. After the second slicing, names becomes [newton]. See this for the complete explanation: https://play.golang.org/p/EsEHrSeByFR


## What does this code print?
```go
i := 2
s := fmt.Sprintf("i = %d * %d = %d", i, i, i*i)
fmt.Print(s)
```

1. i = i * i = i*i
2. i = %d * %d = %d
3. i = 2 * 2 = 2
4. i = 2 * 2 = 4 *CORRECT*

> **4:** Awesome! Sprintf works just like Printf. Instead of printing the result to standard out (usually to command-line prompt), it returns a string value.