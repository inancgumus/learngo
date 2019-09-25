# Questions: Packages and Scopes

---
## 1. Which of the choices is a valid package type in Go?
(A) Empty package <br>
(B) Executable package <br>
(C) Transferable package <br>
(D) Librarian package <br>

---
**Solution**: (B) <br>
The `package main` declaration tells Go to execute the code in that file as the main executable package.

---
## 2. Which of the follow package types will execute when `go run` is called?
(A) Empty package <br>
(B) Executable package <br>
(C) Transferable package <br>
(D) Library package <br>

---
**Solution**: (B) <br>
The executable package is the only package that runs. <br>
Library packages are compiled for the package to use but not run. <br>

---
## 3. Which package type that `go build` can compile?
(A) Empty package <br>
(B) Temporary package <br>
(C) Both of executable and library packages  <br>
(D) Transferable package <br>

---
**Solution**: (C) <br>
The key term is `build`. Only executable functions are called during runtime. 
<br>
Library functions are only compiled.

---
## 4. Which one is an executable package?
(A) `package main` with `func main()` <br>
(B) `package Main` with `func Main()` <br>
(C) `package exec` with `func exec()` <br>

---
**Solution**: (A) <br>
The `main()` function is a special function that Go will begin to execute at runtime.

---
## 5. Which one is a library package?
(A) `main package` <br>
(B) `package lib`  <br>
(C) `func package` <br>
(D) `package main` with `func main()` <br>

---
**Solution**: (B) <br>
Except for (D), all other types are not valid package types. <br>
Because the question specifies library packages, (D) is incorrect, since `package main` and a `main()` makes in executable.

---
## 6. Which package is used for an executable Go program?
(A) Empty package <br>
(B) Executable package <br>
(C) Transferable package <br>
(D) Library package <br>

---
**Solution**: (B) <br>
Again, if there is a `main()` function, the package going to be the executable package that gets executed at runtime.

---
## 7. Which package is used for reusability and can be imported? 
(A) Empty package <br>
(B) Executable package  <br>
(C) Transferable package <br>
(D) Library package <br>

---
**Solution**: (D) <br>
Just like the `fmt.Println()` function can be called multiple times throughout multiple functions in a single Go package, other library functions import reusable code.