# Questions: Scopes #

---

## (1) Which of the following best describes a scope? ##

1. A scope is an executable block of code.
2. A scope describes the visibility of declared names in a program.
3. A scope determines what the program will run.

---

**Solution**: (B)

---

## (2) Which of the following choices is in the package scope? ##

```go
package awesome

import "fmt"

var enabled bool

func block() {
	var counter int
	fmt.Println(counter)
}
```

1. `awesome`
2. `fmt`
3. `enabled`
4. `counter`

--- 

**Solution: 3**

That's right. `enabled` is declared outside of a function so it's a package scoped name. 

The`block()` function is also package scoped because it's outside of any function as well.

---

## (3) Which name below is file scoped? ##

1. `awesome`
2. `fmt`
3. `enabled`
4. `block`
5. `counter`

---

**Solution: 2**

Correct. Imported package names file scoped.

_Note: They can only be used within the same file._

---

## (4) Which name below is declared in the `block()` function's scope?

1. `awesome`
2. `fmt`
3. `enabled`
4. `block`
5. `counter`

---

**Solution: 5**

You got it. `counter` is declared within the `block()` function, so it's in the scope of the `block()` function. Outside of the `block()` function, other code segments cannot see it.

---

## (5) Is `enabled` visible to the `block()` function? ##

1. Yes: It's in the package scope.
2. No: It's in the file scope.
3. No: It's in the `block()` function's scope.

---

**Solution: 5**

That's right. All other parts of the code can see declared names at the package level.

---

## (6) Can other files in the `awesome` package see `counter`? ##

1. Yes.
2. No: It's in the package scope.
3. No: It's in the file scope.
4. No: It's in the `block()` function's scope.

---

**Solution: 3**

Imported packages are only visible to packages in the file that they are imported, even though they are in the same package.

---

## (7) Which of the following statements is true about the declared name `enabled`?

```go
package awesome

import "fmt"

// declared at the package scope
var enabled bool

func block() {
	// also declared in the block scope
	var enabled bool
	
	var counter int
	fmt.Println(counter)
}
```

1. The newly declared name `enabled` will override the previous one.
2. The program has an error because `enabled` is already declared at the package scope.
3. The program has an error because `enabled` is already declared at the file scope.

---

**Solution: 1**

That's right. You can declare the same name because names declared inside of the function `block()` are not visible to named declared outside of it. The second `enabled` overrides the `enabled` from the parent scope.

Check out the course repository to find out more.