# Questions: Scopes #

---

## 1. Which of the following best describes a scope? ##

(A) A scope is an executable block of code.

(B) A scope describes the visibility of declared names in a program.

(C) A scope determines what the program will run.

---

**Solution**: (B)

---

## 2. Which of the following choices is in the package scope? ##

```go
package awesome

import "fmt"

var enabled bool

func block() {
	var counter int
	fmt.Println(counter)
}
```

(A) `awesome`

(B) `fmt`

(C) `enabled`

(D) `counter`

--- 

**Solution**: (C)

That's right. `enabled` is declared outside of a function so it's a package scoped name. 

The`block()` function is also package scoped because it's outside of any function as well.

---

## 3. Which name below is file scoped? ##

(A) `awesome`

(B) `fmt`

(C) `enabled`

(D) `block`

(E) `counter`

---

**Solution**: (B)

Correct. Imported package names file scoped.

_Note: They can only be used within the same file._

---

## 4. Which name below is declared in the `block()` function's scope?

(A) `awesome`

(B) `fmt`

(C) `enabled`

(D) `block`

(E) `counter`

---

**Solution**: (E)

You got it. `counter` is declared within the `block()` function, so it's in the scope of the `block()` function. Outside of the `block()` function, other code segments cannot see it.

---

## 5. Is `enabled` visible to the `block()` function? ##

(A) Yes: It's in the package scope.

(B) No: It's in the file scope.

(C) No: It's in the `block()` function's scope.

---

**Solution**: (A)

That's right. All other parts of the code can see names declared at the package level.

---

## 6. Can other files in the `awesome` package see `counter`? ##

(A) Yes.

(B) No: It's in the package scope.

(C) No: It's in the file scope.

(D) No: It's in the `block()` function's scope.

---

**Solution**: (C)

Imported packages are only visible to packages in the file that they are imported, even though they are in the same package.

---

## 7. Which of the following statements is true about the declared name `enabled`?

```go
package awesome

import "fmt"

// declared at the package scope
var enabled bool
var enabled bool

func block() {
	// also declared in the block scope
	var enabled bool
	
	var counter int
	fmt.Println(counter)
}
```

(A) The newly declared name `enabled` will override the previous one.

(B) The program has an error because `enabled` is already declared at the package scope.

(C) The program has an error because `enabled` is already declared at the file scope.