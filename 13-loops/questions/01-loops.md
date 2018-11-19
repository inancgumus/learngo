## Which one of these is a valid loop statement in Go?
1. while
2. forever
3. until
4. for *CORRECT*

> **4:** Correct. There is only one loop statement in Go.


## What does this code print?
```go
for i := 3; i > 0; i-- {
    fmt.Println(i)
}
```
1. 3 2 1 *CORRECT*
2. 1 2 3
3. 0 1 2
4. 2 1 0


## What does this code print?
```go
for i := 3; i > 0; {
    i--
    fmt.Println(i)
}
```
1. 3 2 1
2. 1 2 3
3. 0 1 2
4. 2 1 0 *CORRECT*


## What does this code print?
```go
for i := 3; ; {
    if i <= 0 {
        break
    }

    i--
    fmt.Println(i)
}
```
1. 3 2 1
2. 1 2 3
3. 0 1 2
4. 2 1 0 *CORRECT*


## What does this code print?
```go
for i := 2; i <= 9; i++ {
    if i % 3 != 0 {
        continue
    }

    fmt.Println(i)
}
```
1. 3 6 9 *CORRECT*
2. 9 6 3
3. 2 3 6 9
4. 2 3 4 5 6 7 8 9


## How can you simplify this code?
```go
for ; true ; {
    // ...
}
```
1. ```go
   for true {
   }
   ```
2. ```go
   for true; {
   }
   ```
3. ```go
   for {
   }
   ```
   *CORRECT*
4. ```go
   for ; true {
   }
   ```


## What does this code print?
Let's say that you run the program like this:
```bash
go run main.go go is awesome
```

```go
for i, v := range os.Args[1:] {
    fmt.Println(i+1, v)
}
```
1. ```
   1 go
   2 is
   3 awesome
   ```
   *CORRECT*
2. ```
   go
   is
   awesome
   ```
3. ```
   0 go
   1 is
   2 awesome
   ```
4. ```
   1
   2
   3
   ```


## What does this code print?
Let's say that you run the program like this:
```bash
go run main.go go is awesome
```

```go
for i := range os.Args[1:] {
    fmt.Println(i+1)
}
```
1. ```
   1 go
   2 is
   3 awesome
   ```
2. ```
   go
   is
   awesome
   ```
3. ```
   0 go
   1 is
   2 awesome
   ```
4. ```
   1
   2
   3
   ```
   *CORRECT*


## What does this code print?
Let's say that you run the program like this:
```bash
go run main.go go is awesome
```

```go
for _, v := range os.Args[1:] {
    fmt.Println(v)
}
```
1. ```
   1 go
   2 is
   3 awesome
   ```
2. ```
   go
   is
   awesome
   ```
   *CORRECT*
3. ```
   0 go
   1 is
   2 awesome
   ```
4. ```
   1
   2
   3
   ```


## What does this code print?
Let's say that you run the program like this:
```bash
go run main.go go is awesome
```

```go
var i int

for range os.Args {
    i++
}

fmt.Println(i)
```
1. go is awesome
2. 1 2 3
3. 2
4. 4 *CORRECT*

> **4:** As you can see, you can also use a for range statement for counting things.
