## What's the difference between `go build` and `go run`?
1. `go run` just compiles a program; whereas `go build` both compiles and runs it.
2. `go run` both compiles and runs a program; whereas `go build` just compiles it. *CORRECT*

> **1:** It's opposite actually.
>
>
> **2:** `go run` compiles your program and puts it in a temporary directory. Then it runs the compiled program in there.
>
>

## Which one below is true for runtime?
1. It happens when your program starts running on a computer *CORRECT*
2. It happens while your program is being compiled


## Which one below is true for the compile-time?
1. It happens when your program starts running in a computer
2. It happens while your program is being compiled  *CORRECT*


## In which stage your program can print a message to the console?
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
