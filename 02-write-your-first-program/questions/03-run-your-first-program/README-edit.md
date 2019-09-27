# Questions: Run Your First Program

---

## 1. What is the difference between `go build` and `go run`? ##

(A) `go run` only compiles the program; `go build` compiles and runs it.

(B) `go run` compiles and runs the program; `go build` it.

---

**Solution**: (B)

`go run` first compiles a program into a temporary directory, then runs it. That is why you can see text printed to the console after calling the command.

---

## 2. What is the name of the directory that Go saves compiled code? ##

(A) The same directory where `go build` is called

(B) The `$GOPATH/src` directory. 

(C) The `$GOPATH/pkg` directory.

(D) Into a temporary directory.

---

**Solution**: (A)

Running `go build` will compile the program in the directory that it was called. If you look at files in the directory, you should be able to see a new executable program.

Go will only save the compiled code if you run `go install`.

---

## 3. Which of the following statements is true about runtime? ##

(A) Runtime occurs when your program starts running on a computer.

(B) Runtime occurs when your program is being compiled.

---

**Solution**: (A)

The keyword is `run`, which is used in the `go run` command.

---

## 4. Which of the following statements is true about compile-time? ##

(A) Compile-time occurs when your program starts running on a computer.

(B) Compile-time occurs when your program is being compiled.

---

**Solution**: (B)

Compile-time describes the time when Go is building the program so that it can run.

---

## 5. When can Go programs print messages to the console? ##

(A) While the program is being compiled.

(B) While it is running, but only after compile-time.

(C) While it is running, but only before compile-time.

---

**Solution**: (B)

During compile-time, a go program is "dead". During compile-time is when Go builds the programs so that it **can** run. Only after compile-time does a program begin executing.