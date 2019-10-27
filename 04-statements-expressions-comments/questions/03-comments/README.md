# Questions: Comments #

---

## (1) Why are comments used in source code files? ##

1. To combine different expressions together.
2. To provide explanations or generate automatic documentation for your code.

3. To make the code look nice and beautiful.

---

**Solution: (B)** 

That's right! Leaving comments throughout your code will make it easier for others (and yourself) to understand what the code does. Commenting on your code can also automatically generate documentation. 

---

## (2) Which of the following examples shows the correct usage of comments? #

1.
```go
package main

/ main function is an entry point /
func main() {
    fmt.Println("Hi")
}
```

2.
```go
package main

// main function is an entry point /*
func main() {
    fmt.Println(/* this will print Hi! */ "Hi")
}
```

3.
```go
package main

/*
main function is an entry point

It allows Go to find where to start executing an executable program.
*/
func main() {
    fmt.Println(// "this will print Hi!")
}
```

---

**Solution: (B)** 

Multiline comments can be put almost anywhere as long as the begin with `/*` and end with `*/`. Anything between `/* ... */` is ignored by Go. Whenever go sees `//` as the first two characters on a line of code, Go skips the entire line.

> 1. This example attempts to place a single-line comment between `/ ... /`. Single-line comments must begin with `//` and the comment must be placed after.
> 3. This example attempts to place a single-comment inside `fmt.Println()`, and by doing so, tells Go to ignore everything after.