# Questions: Packages and Scopes (Part A) #

---

## (1) Where should a package's source code files be stored? ##

1. Each file should be saved in a different directory.
2. In a single directory.

---

**Solution: 2**

A package's source code should be in the same directory so that features/functionality can be shared between the various source files.

---

## (2) How are package clauses used in a `.go` files? ##
1. They are used to import packages.
2. They are used to inform Go whether or not a file belongs to a package.
3. They are used for declaring new functions.

---

**Solution: 2**

> 1. The `import` statement is used to import packages.
> 3. The keyword `func` is used for declaring new functions.  

---

## (3) Where should `package clause` be declared in a `.go` file? ##

1. It should be declared in the first line of code.
2. It should always be the finals line of code.
3. The `package clause` can be placed anywhere.

---

**Solution: 1**

The `package clause` must always be the first line of readable code interpreted by the Go compiler.

---

## (4) Which of the following is true? ##

1. All files belonging to the same package cannot call each other's function. 
2. All files belonging to the same package can call each other's functions. 

---

**Solution: 2**

Functions that are defined in a Go package are exportable and readable by other functions in the same directory as long as they are defined within the package scope.

---

## (5) Which of the following is is the command-line argument used to fun multiple Go files. ##

1. `go run *go`
2. `go build *go`
3. `go run go`
4. `go run *.go`

---

**Solution: 4**

The `*` represents a wild-card that will execute any file that ends in `.go`. 

A more verbose way to run all `.go` files in a directory is to explicitly call each file as an argument following `go run <file1> <file2> <file3>`