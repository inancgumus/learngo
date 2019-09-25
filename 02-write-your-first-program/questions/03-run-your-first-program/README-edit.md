# Questions: Run Your First Program
---
## 1. What is the difference between `go build` and `go run`? ##
(A) `go run` only compiles the program; `go build` compiles and runs it. <br>
(B) `go run` compiles and runs the program; `go build` it. <br>

---
**Solution**: (B)
`go run` first compiles a program into a temporary directory, then runs it. That is why you are able to see text printed to the console after calling the command. <br>

---
## 2. What is the name of the directory that Go saves compiled code? ##
(A) The same directory where `go build` isi called <br>
(B) The `$GOPATH/src` directory.  <br>
(C) The `$GOPATH/pkg` directory. <br>
(D) Into a temporary directory. <br>

---
**Solution**: (A) <br>
Running `go build` will compile the program in the directory that it was called. If you look at files in the directory, you should be able to see a new executable program. <br>
Go will only save the compiled code if you run `go install`.

---
## 3. Which of the following statements is true about runtime? ##
(A) Runtime occurs when your program starts running on a computer. <br>
(B) Runtime occurs when your program is being compiled. <br>

---
**Solution**: (A) <br>
The key word is `run`, which is used in the `go run` command.

---
## 4. Which of the following statements is true about compile-time? ##
(A) Compile-time occurs when your program starts running on a computer. <br>
(B) Compile-time occurs when your program is being compiled. <br>

---
**Solution**: (B)
The compile-time is the time before Go runs the program where Go is builds.the program.

---
## 5. When can Go programs print messages to the console? ##
(A) While the program is being compiled. <br>
(B) While it is running, but only after compile-time. <br>
(C) While it is running, but only before compile-time. <br>

---
**Solution**: (B) <br>
During compile-time, a go program is "dead". During compile-time is when Go builds the programs so that it **can** run. Only after compile-time  does a program begin executing.