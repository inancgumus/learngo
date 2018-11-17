## Which one below is correct?
**NOTE** _There are explanations inside the answers. Even if you know the answer please try to select all of them one by one, so you can read the explanations._

1. You can run a library package.
2. In a library package, there should be a function named main (func main).
3. You can compile a library package. *CORRECT*
4. You have to compile a library package.

> **1:** You can't, but you can import it from other packages.
>
> **2:** In a library package, you don't have to include the main function. Because it isn't an executable package. Only in executable packages you need a main func.
>
> **4:** You don't have to compile it. When you import it, it will automatically be built by the other program or library when it gets compiled or ran.


## What do you need to export a name?
1. You need to type it in all capital letters
2. You need to type its first letter as a capital letter *CORRECT*
3. You need to put it inside a function scope
4. You need to create a new file for that name

> **1:** When you do so, it will be exported, that's true, but don't do that; so I assume that this answer is not correct :)
>
> **2:** That's right. Then the other packages can access it.
>
> **3:** It should be in a package scope, not function scope.
>
> **4:** You don't have to do that.


## How can you use a function from your library from an executable program?
1. You need to export your library package first; then you can access its imported names
2. You need to import your library package first; then you can access its exported names *CORRECT*
3. You can access your library package as if it's in your executable program
4. You can import it just by using its name

> **1:** You can't export packages. All packages are already exported. Unless you put them in a directory called: "internal". But, that's an advanced topic for now.
>
> **2:** That's right.
>
> **3:** You can't access a package from another package without importing it.
>
> **4:** No, you can't. You need to import it using its full directory path after GOPATH. BTW, in the near future, this may change with the Go modules support.


## In the following program, which names are exported?
```go
package wizard

import "fmt"

func doMagic() {
    fmt.Println("enchanted!")
}

func Fireball() {
    fmt.Println("fireball!!!")
}
```

1. fmt
2. doMagic
3. Fireball *CORRECT*
4. Println

> **1:** That's just an imported package name.
>
> **2:** It starts with a lowercase letter; so, it's not exported.
>
> **3:** That's right. It starts with a capital letter.
>
> **4:** This isn't your function. It's already been exported in the fmt package. But, it isn't exported here.


## In the following program, which names are exported?
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

1. doMagic and Fireball
2. Fireball and Two *CORRECT*
3. Fireball, greenTrees and Two
4. Fireball, greenTrees, one and Two

> **1:** doMagic starts with a lowercase letter; so, it's not exported.
>
> **2:** That's right. Fireball and Two, they both start with a capital letter.
>
> **3:** greenTrees starts with a lowercase letter; so, it's not exported.
>
> **4:** one and greenTrees do not start with capital letters; so, they're not exported.
