# Questions: Run Your First Program

---

## (1) What is the difference between `go build` and `go run`? ##

1. `go run` only compiles the program; `go build` compiles and runs it.
2. `go run` compiles and runs the program; `go build` only compiles it.

---

**Solution: 1**

`go run` first compiles a program into a temporary directory, then runs it. That is why you can see text printed to the console after calling the command.

---

## (2) What is the name of the directory that Go saves compiled code? ##

1. The same directory where `go build` is called
2. The `$GOPATH/src` directory
3. The `$GOPATH/pkg` directory
4. Into a temporary directory

---

**Solution: 1**

Running `go build` will compile the program in the directory that it was called. If you look at files in the directory after building the program, you should be able to see a new executable program.

> 2. Only Go source-code files should be in the `$GOPATH/src` directory
> 3. Go will only save the compiled code if you run `go install`.

---

## (3) Which of the following statements is true about runtime? ##

1. Runtime occurs when your program starts running on a computer.
2. Runtime occurs when your program is being compiled.

---

**Solution: 1**

The keyword is `run`, which is used in the `go run` command.

---

## (4) Which of the following statements is true about compile-time? ##

1. Compile-time occurs when your program starts running on a computer.
2. Compile-time occurs when your program is being compiled.

---

**Solution: 2**

Compile-time describes the time when Go is building the program so that it can run.

---

## (5) When can Go programs print messages to the console? ##

1. While the program is being compiled.
2. While it is running, but only after compile-time.
3. While it is running, but only before compile-time.

---

**Solution: 2**

Correct! The only time your programs can interact with a computer and print messages to the console is after they are done compiling.

> 1. During compile-time, your program cannot print any messages because it is "dead".  
> 2. A Go program can run only after it has finished compiling. 