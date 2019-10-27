# Questions: Packages and Scopes (Part B) #

---

## (1) Which of the choices is a valid package type in Go?

1. Empty package
2. Executable package
3. Transferable package 
4. Librarian package 

---

**Solution: 2**

The `package main` declaration tells Go to execute the code in that file as the main executable package.

---

## (2) Which of the following package types will execute when `go run` is called?

1. Empty package
2. Executable package
3. Transferable package 
4. Librarian package 

---

**Solution: 2**

The executable package is the only package that runs.

> 4. Library packages are compiled for the package to use but not run.

---

## (3) Which package type does `go build` compile?

1. Empty package
2. Temporary package
3. Both of executable and library packages 
4. Transferable package

---

**Solution: 3**

Only library packages are compiled. The key term is `build`. Only executable functions are called during runtime.

---

## (4) Which of the following is an executable package?

1. `package main` with `func main()`
2. `package Main` with `func Main()`
3. `package exec` with `func exec()`

---

**Solution: 1**

The `main()` function is a special function that Go will begin to execute at runtime.

---

## (5) Which one is a library package?

1. `main package`
2. `package lib` 
3. `func package`
4. `package main` with `func main()`

---

**Solution: 2**

Except for 4., all other types are not valid package types.
Because the question specifies library packages, 4. is incorrect, since `package main` and a `main()` makes in executable.

---

## 6. Which package is used for an executable Go program?

1. Empty package
2. Executable package
3. Transferable package
4. Library package

---

**Solution: 2**

Again, if there is a `main()` function, the package going to be the executable package that gets executed at runtime.

---

## 7. Which package is used for reusability and can be imported? 

1. Empty package
2. Executable package 
3. Transferable package
4. Library package

---

**Solution: 4**

Library functions import reusable code. For example, the `fmt.Println()` function can be called multiple times after being imported.

Other library functions import reusable code. Just like the `fmt.Println()` function can be called multiple times throughout multiple functions in a single Go package, 