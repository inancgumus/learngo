## What's a scope?
* Executable block of code
* The visibility of the declared names **CORRECT**
* Determines what to run

```go
package awesome

import "fmt"

var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```

## Which name below is package scoped?
1. awesome
2. fmt
3. enabled **CORRECT**
4. counter

> **3:** That's right. `enabled` is out of any functions, so it's a package scoped name. `block()` function is also package scoped; it's out of any function too.
>
>


## Which name below is file scoped?
1. awesome
2. fmt **CORRECT**
3. enabled
4. block()
5. counter

> **2:** That's right. Imported package names are file scoped. And they can only be used within the same file.
>
>


## Which name below is in the scope of the block() func?
1. awesome
2. fmt
3. enabled
4. block()
5. counter **CORRECT**

> **5:** That's right. `counter` is declared within the `block()` func, so it's in the scope of the block func. Out of the `block()` func, other code can't see it.
>
>


## Can `block()` see `enabled` name?
1. Yes: It's in the package scope **CORRECT**
2. No: It's in the file scope
3. No: It's in the block scope of block()

> **1:** All code inside the same package can see all the other package level declared names.
>
>


## Can other files in `awesome` package see `counter` name?
1. Yes
2. No: It's in the package scope
3. No: It's in the file scope
4. No: It's in the block scope of block() **CORRECT**

> **4:** That's right. None of the other code can see the names inside the `block()` function. Only the code inside the `block()` function can see them (Only to some extent. For example: Inside the block, the code can only see the variables declared before it.)
>
>


## Can other files in `awesome` package see `fmt` name?
1. Yes
2. No: It's in the package scope
3. No: It's in the file scope **CORRECT**
4. No: It's in the block scope of block()

> **3:** Only the same file can see the imported packages, not the other files whether they're in the same package or not.
>
>


## What happens if you declare the same name in the same scope as this:
```go
package awesome

import "fmt"

// declared twice in the package scope
var enabled bool
var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```
1. The newly declared name will override the previous one.
2. I can't do that. It's already been declared at the package scope. *CORRECT*
3. I can't do that. It's already been declared at the file scope.

> **2:** That's right. You can't declare the same name in the same scope. If you could do so, then how would you access it again? Or to which one?
>
>


## What happens if you declare the same name in another scope like this:
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
1. The newly declared name will override the previous one. *CORRECT*
2. I can't do that. It's already been declared at the package scope.
3. I can't do that. It's already been declared at the file scope.

> **1:** Actually, you can declare the same name in the inner scopes like this. `block()`'s scope is inside its package. This means that it can access to its package's scope (but not vice versa). So, `block()`'s scope is under its package's scope. This means that you can declare the same name again. It will override the parent scope's name. They both can be exist together. Check out the example in the course repository to find out.
>
>
