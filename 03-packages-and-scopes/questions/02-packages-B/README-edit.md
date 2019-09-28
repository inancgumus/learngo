# Questions: Packages (B)

---

## 1. Which of the choices is a valid package type in Go?

(A) Empty package

(B) Executable package

(C) Transferable package

(D) Librarian package

---

**Solution**: (B)

The `package main` declaration tells Go to execute the code in that file as the main executable package.

---

## 2. Which of the following package types will execute when `go run` is called?

(A) Empty package

(B) Executable package

(C) Transferable package

(D) Library package

---

**Solution**: (B)

The executable package is the only package that runs.

Library packages are compiled for the package to use but not run.

---

## 3. Which package type that `go build` can compile?

(A) Empty package

(B) Temporary package

(C) Both of executable and library packages 

(D) Transferable package

---

**Solution**: (C)

The key term is `build`. Only executable functions are called during runtime. 
Library functions are only compiled.

---

## 4. Which one is an executable package?

(A) `package main` with `func main()`

(B) `package Main` with `func Main()`

(C) `package exec` with `func exec()`

---

**Solution**: (A)
The `main()` function is a special function that Go will begin to execute at runtime.

---

## 5. Which one is a library package?

(A) `main package`

(B) `package lib` 

(C) `func package`

(D) `package main` with `func main()`

---

**Solution**: (B)

Except for (D), all other types are not valid package types.
Because the question specifies library packages, (D) is incorrect, since `package main` and a `main()` makes in executable.

---

## 6. Which package is used for an executable Go program?

(A) Empty package

(B) Executable package

(C) Transferable package

(D) Library package

---

**Solution**: (B)

Again, if there is a `main()` function, the package going to be the executable package that gets executed at runtime.

---

## 7. Which package is used for reusability and can be imported? 

(A) Empty package

(B) Executable package 

(C) Transferable package

(D) Library package

---

**Solution**: (D)

Just like the `fmt.Println()` function can be called multiple times throughout multiple functions in a single Go package, other library functions import reusable code.