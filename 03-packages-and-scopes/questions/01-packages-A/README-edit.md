# Questions: Packages and Scopes #
---

## 1. Where should a package's source code files be stored? ##
(A) Each file should be saved in a different directory. <br>
(B) In a single directory. <br>

---
**Solution**: (B) <br>
A packages source should be in the same directory so that features/functionality can be shared between the various source files.

---
## 2. How are package clauses used in a `.go` files? ##
(A) They are used to import packages. <br>
(B) They are used to inform Go whether or not a file belongs to a package. <br>
(C) They are used for declaring new functions. <br>

---
**Solution**: (B) <br>
(A) The `import` statement is used to import packages. <br>
(C) The keyword `func` is used for declaring new functions. <br>

---

## 3. Where should `package clause` be declared in a `.go` file? ##
(A) It should be declared in the first line of code. <br>
(B) It should always be the finals line of code. <br>
(C) The `package clause` can be placed anywhere. <br>
(D) The `package clause` can only be execute. <br>

---
**Solution**: (A)
The `package clause` must always be the first line of readable code interpreted by the Go compiler.

---

## 4. Which of the following is true? ##
(A) All files belonging to the same package cannot call each other's function. 
<br>
(B) All files belonging to the same package can call each other's functions. 
<br>

---
**Solution**: (B) <br>
Functions that are defined in a Go package are exportable and readable by other functions in the same directory as long as they are defined within the package scope.

---
## 5. Which of the following is is the command-line argument used to fun multiple Go files. ##
(A) `go run *go` <br>
(B) `go build *go` <br>
(C) `go run go` <br>
(D) `go run *.go` <br>

---
**Solution**: (D)
The `*` represents a wild-card that will execute any file that ends in `.go`. 
<br>
A more verbose way to run all `.go` files in a directory is to explicitly call each file as an argument following `go run <file1> <file2> <file3>`