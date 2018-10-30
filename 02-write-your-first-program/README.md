# Cheatsheet: Write your First Go Program

Hi! For reference, you can store this cheatsheet after you take the lectures in this section.

You can also print this cheatsheet and then follow the video lectures in this section along with it.

Enjoy!

---

## COMMAND LINE COMMANDS:

* Enter into a directory: `cd directoryPath`

* **WINDOWS:**

    * List files in a directory: `dir`

* **OS X & LINUXes:**

    * List files in a directory: `ls`

---

## BUILDING & RUNNING PROGRAMS:

* **Build a Go program:**

    * While inside a program directory, type:
        * `go build`

* **Run a Go program:**

    * While inside a program directory, type:
        * `go run main.go`

<br /><br /><br /><br />
_See the next page..._

<div style="page-break-after: always;"></div>

## WHAT IS $GOPATH?

* _$GOPATH_ is an environment variable which points to a directory where the downloaded and your own Go files are stored.

    * **On Windows**, it's in: `%USERPROFILE%\go`

    * **On OS X & Linux**, it's in: `~/go`

    * **NOTE:** Never set your GOPATH manually. It's by default under your users directory.

* **GOPATH has 3 directories:**

    * **src:** Contains the source files for your own or other downloaded packages. You can build and run programs while in a program directory under this directory.

    * **pkg:** Contains compiled package files. Go uses this to make the builds (compilation & linking) faster.

    * **bin:** Contains compiled and executable Go programs. When you call go install on a program directory, Go will put the executable under this folder.

        * _You might want to add this to your `PATH` environment variable if it's not there already._

---

## WHERE YOU SHOULD PUT YOUR SOURCE FILES?

* `$GOPATH/src/github.com/yourUsername/programDirectory`

* **Example:**

    * My github username is: inancgumus

    * So, I put all my programs under: `~/go/src/github.com/inancgumus/`

    * And, let's say that I've a program named hello, then I put it under this directory: `~/go/src/github.com/inancgumus/hello`

<br /><br /><br /><br />
_See the next page..._

<div style="page-break-after: always;"></div>

## FIRST PROGRAM

* Under, `~/go/src/github.com/inancgumus/hello`, create a `main.go` file.
* And add the following code to it.
* Then, run it like this: `go run main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

That's all! Enjoy!

---

## NOTE:

* BTW, There's a new *Go Modules* support which allows you to run your programs in any directory that you want. After the section ends, you'll also find an information about it.

> For more tutorials: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> Learn Go Programming Course
> 
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)