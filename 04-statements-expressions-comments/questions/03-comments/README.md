## Why do you need to use comments sometimes?
1. To combine different expressions together
2. To provide explanations or generating automatic documentation for your code *CORRECT*
3. To make the code look nice and beautiful


## Which of the following code is correct?
1.
```go
package main

/ main function is an entry point /
func main() {
    fmt.Println("Hi")
}
```

2. *CORRECT*
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

> **1:** `/` is not a comment. It should be `//`.
>
>
> **2:** Multiline comments can be put almost anywhere. However, when a comment starts with `/*`, it also needs to end with `*/`. Here, Go doesn't interpret `/* ... */`, it just skips it. And, when Go sees `//` as the first two characters in a code, it skips the whole line.
>
>
> **3:** `//` prevents Go to interpret the rest of the code line. That's why this code doesn't work. Go can't interpret this part because of the comment: `"this will print Hi!")`
>
>

## How should you name your code so that Go can generate documentation from your code automatically?
1. By commenting each line of the code; then it will generate the documentation from whatever it sees
2. By starting the comments using the name of the declared names *CORRECT*
3. By using multi-line comments

> **1:** This won't help. Sorry.
>
>
> **3:** It doesn't matter whether you type your comments using single-line comments or multi-line comments.
>
>


## Which tool do you need to use from the command-line to print the documentation?
1. go build
2. go run
3. go doctor
4. go doc *CORRECT*


## What's the difference between `godoc` and `go doc`?
1. `go doc` is the real tool behind `godoc`
2. `godoc` is the real tool behind `go doc` *CORRECT*
3. `go` tool is the real tool behind `go doc`
4. `go` tool is the real tool behind `godoc`

> **2:** That's right. go doc tool uses godoc tool behind the scenes. go doc is just a simplified version of the godoc tool.
>
>
