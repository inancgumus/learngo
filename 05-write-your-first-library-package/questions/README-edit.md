# Questions: Library Packages #

---

## 1. Which of the following choices is true? ##

---

**Note**: Each answer choice has an explanation so even if you know the answer, select all of them one by one to read the explanations.

(A) You can run a library packages.

(B) In a library package, there should be a function named `main()`

(C) You can compile library packages.

(D) You have to compile library packages.

---

**Solution**: (C)

(A) Library packages cannot be run; they can only be imported from other packages.

(B) In library packages, you don't have to include the `main()` function because it isn't an executable package. Only executable packages need a `main()` function.

(D) Library packages don't need to be compiled. When you run your program, library packages will automatically be built by the library or package that it came from.

---

## 2. Which of the following is true about exported names? ##

(A) Exported names must be typed in all capital letters.

(B) The first letter of an exported name must be a capital letter.

(C) An exported name must be inside the function scope.

(D) A new file must be created just for the exported name.

---

**Solution**: (B)

That's right. Only the first letter needs to be capitalized.

(A) If a name is declared in all capital letters, it will be exported but it is not a **must**. It is only exported because the first letter is capitalized.

(C) Exported names should be declared at the package scope, not the function scope.

(D) Creating a new file just for the exported name is not necessary.

---

## 3. How can you access your library package's functions in an executable program? ##

(A) The library package must be exported before you can access its imported names.

(B) You need to import your library package first, then you can access it's exported names.

(C) You can access your library package anywhere as long as it is in the same directory as your executable program.

(D) You can import it by using just by using the function's name.

---

**Solution**: (B)

(A) You can't access packages. All packages are already exported (_unless you put them in a directory called `"internal"`, but more on that later!_)

(C) That's right. The package must be imported first before its functions are accessible.

(D) You can't use functions from a library package just by calling the function's name. You need to import it by using the full directory path after the `$GOPATH`.

---


## 4. In the following program, which names are exported? ##

```go
package wizard

import "fmt"

func doMagic() {
    fmt.Println("enchanted!")
}

func Fireball() {
    fmt.Println("Fireball!!!")
}
```

(A) `fmt`

(B) `doMagic()`

(C) Fireball()

(D) `Println()`

---

**Solution**: (C)

(A) `fmt` is just the name of the imported package.

(B) `doMagic()` does not start with a lowercase letter, so it is not exported.

(C) Correct! It starts with an uppercase letter so it gets exported.

(D) `Println()` is not exported because it has already been exported from `fmt` package.

---


## 5. In the following programs, which names are exported? ##

```go
package wizard
import "fmt"

var one string
var Two string
var greenTrees string

func doMagic() {
    fmt.Println("enchanted!")
}

func Fireball() {
    fmt.Println("fireball!!!")
}
```

(A) `doMagic()` and `Fireball()`

(B) `Fireball()` and `Two`

(C) `Fireball()`, `greenTrees` and `Two`

(D) `Fireball()`, `greenTrees`, `one` and `Two`

---

**Solution**: (B)

That's right. `Fireball()` and `Two` both start with capital letters so the will be exported.

(A) `doMagic()` starts with a lowercase letter, so it is not exported.

(C) `greenTrees` starts with a lowercase letter as well.

(D) `one` and `greenTrees` both do not start with capital letters, so they will not be exported.