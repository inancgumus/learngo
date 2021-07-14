# Cheatsheet: Write your First Go Program

Hi!

For reference, you can store this cheatsheet after you take the lectures in this section.

You can also print this cheatsheet and then follow the video lectures in this section along with it.

Enjoy!

---

## COMMAND LINE COMMANDS:

* Enter into a directory: `cd directoryPath`

* **WINDOWS:**

    * List files in a directory: `dir`

* **OS X & LINUXes:**

    * List files in a directory: `ls`

## BUILDING & RUNNING GO PROGRAMS:

* **Build a Go program:**

    * While inside a program directory, type:
        * `go build main.go`

* **Run a Go program:**

    * While inside a program directory, type:
        * `go run main.go`

## WHERE YOU SHOULD PUT YOUR SOURCE FILES?

* In any directory you like.

## FIRST PROGRAM

### Create a directory
* Create a new directory:
  * `mkdir myDirectoryName`
* Go to that directory:
  * `cd myDirectoryName`

### Add the source code files
* Create a new `code main.go` file.
  * This is going to the create the file in the folder, and open up it in the Visual Studio Code.
* And add the following code to it and save it.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

### Run the program
* Finally, return back to the command-line.
  * Run it like this: `go run main.go`
* If you create other files and run them all, you can use this command:
  * `go run .`

That's all! Enjoy!

> For more tutorials: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
