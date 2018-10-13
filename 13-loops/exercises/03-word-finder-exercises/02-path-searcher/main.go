// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

// ---------------------------------------------------------
// EXERCISE: Path Searcher
//
//  Your program should search inside the path environment
//  variable.
//
//  Remove the corpus constant then get the corpus from the
//  environment variable "Path" or "PATH" which
//  constains paths to the executable programs on your
//  operating system.
//
// HINTS
//  1. Search the web for what is an environment
//     variable and how to use it, if you don't know
//     what it is.
//
//  2. Look up for the necessary function for getting
//     an environment variable. It's in the "os" package.
//
//     Search for it on the Go online documentation.
//
//  3. Look up for the necessary function for splitting
//     the path variable into directories. It's in
//     the "strings" package.
//
// EXAMPLE
//  For example, on my Mac, my PATH environment variable
//  looks like this:
//
//    "/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//
//    go run main.go /sbin
//
//  It should print this:
//
//    #2 : "/sbin"
// ---------------------------------------------------------

// ---------------------------------------------------------
// BONUS EXERCISE
//  Make your program cross platform. So, it can search
//  the path environment variable when you run it on
//  a Windows or on a Mac (OS X) or on a Linux.
//
// HINT
//  1. What you're looking for is the runtime.GOOS constant.
//  2. Get the operating system name using GOOS.
//  3. Adjust the path environment variable name and
//     the directory separator accordingly.
//
//  FOR EXAMPLE: On OS X, path environment variable's name
//    is PATH and the separator is a colon `:`.
//
//  Or, on Windows, its name is Path and the separator is
//    a semicolon `;`.
// ---------------------------------------------------------

func main() {
}
