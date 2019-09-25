# Questions: Run Your First Program
---

## 1. What is the difference between `go build` and `go run`?
(A) `go run` compiles the program; `go build` compiles and runs it. <br>
(B) `go run` compiles and runs the program; `go build` only compiles it. <br>

---
**Solution**: (B) <br>

`go run` compiles and then runs the program, which allows the program to print output to the command-line directly after calling it. <br>

`go build` compiles the program into an executable binary file that must be called by the main.go directory name. <br>

---

## 2. What directory does Go compile code into? ##
(A) The same directory where `go build` was called. <br>
(B) The `$GOPATH/src` directory. <br>
(C) The `$GOPATH/pkg` <br>
(D) Into a temporary directory. <br>

---
**Solution**: (A) <br>
Running the command to build a `main.go` file will result in a executable file which was placed there after go compiled the code. <br>

(B) Only Go source code files exist in the `$GOPATH/usr` directory <br>
(C) Go will only put code there when `go install` is called. <br>

---

## 2. Which of the following statements is true about runtime? ##
(A) Runtime refers to when the program being running on a computer. <br>
(B) Runtime refers to the during which Go is compiling the code.  <br>

---
**Solution**: (B)
The key work is `run`; runtime occurs only after the compiler has successfully built the program. <br>

---
## 3. Which of the following statements is true about compile-time? ##
(A) Compile-time refers to when the program being running on a computer. <br>
(B) Compile-time refers to the time during which Go is compile the code. <br>

---
## When can a Go program print messages to the console? ##
(A) While the program is being compiled. <br>
(B) During runtime (after compile-time) <br>
(C) During runtime (before compile-time) <br>

---
**Solution**: (B) <br>
Go can only print messages during runtime. <br>
While the compiler is building the program, is "dead" and cannot run. <br>
Runtime can only happen after compile-time.

---
