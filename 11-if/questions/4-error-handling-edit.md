# Questions: Packages and Scopes #

---

## 1. Where should a package's source code files be stored? ##

(A) Each file should be saved in a different directory.

(B) In a single directory.

---

**Solution**: (B)

A packages source should be in the same directory so that features/functionality can be shared between the various source files.

---

## 2. How are package clauses used in a `.go` files? ##

(A) They are used to import packages.

(B) They are used to inform Go whether or not a file belongs to a package.

(C) They are used for declaring new functions.

---

**Solution**: (B)

(A) The `import` statement is used to import packages.

(C) The keyword `func` is used for declaring new functions.

---

## 3. Where should `package clause` be declared in a `.go` file? ##

(A) It should be declared in the first line of code.

(B) It should always be the finals line of code.

(C) The `package clause` can be placed anywhere.

---

**Solution**: (A)

The `package clause` must always be the first line of readable code interpreted by the Go compiler.

---

## 4. Which of the following is true? ##

(A) All files belonging to the same package cannot call each other's functions. 

(B) All files belonging to the same package can call each other's functions.

---

**Solution**: (B)

Functions that are defined in a Go package are exportable and readable by other functions in the same directory as long as they are defined within the package scope.

---

## 5. Which of the following is is the command-line argument used to fun multiple Go files. ##

(A) `go run *go`

(B) `go build *go`

(C) `go run go`

(D) `go run *.go`

---

**Solution**: (D)

The `*` represents a wild-card that will execute any file that ends in `.go`. 

A more verbose way to run all `.go` files in a directory is to explicitly call each file as an argument following `go run <file1> <file2> <file3>`