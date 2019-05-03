## What is the difference between `go build` and `go run`?
1. `go run` just compiles a program; whereas `go build` both compiles and runs it.
2. `go run` both compiles and runs a program; whereas `go build` just compiles it. *CORRECT*

> **1:** It's opposite actually.
>
>
> **2:** `go run` compiles your program and puts it in a temporary directory. Then it runs the compiled program in there.
>
>


## Go saves the compiled code in a directory. What is the name of that directory?
1. The same directory where you call `go build` *CORRECT*
2. $GOPATH/src directory
3. $GOPATH/pkg directory
4. Into a temporary directory.

> **2:** There only lives Go source-code files
>
>
> **3:** Go only puts your code there when you call `go install`.
>
>


## Which is true for runtime?
1. It happens when your program starts running on a computer *CORRECT*
2. It happens while your program is being compiled


## Which is true for the compile-time?
1. It happens when your program starts running on a computer
2. It happens while your program is being compiled  *CORRECT*


## When can a Go program print a message to the console?
1. While it's being compiled.
2. While it runs (after compile-time). *CORRECT*
3. While it runs (inside the compile-time).

> **1:** In the compilation step your program cannot print a message. In that stage, it's literally dead.
>
>
> **2:** That's right. That's the only time which your program can interact with a computer and instruct it to print a message to the console.
>
>
> **3:** Running can only happen after the compile-time
>
>