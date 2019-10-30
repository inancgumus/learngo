// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Infinite Kill
//
//  1. Create an infinite loop
//  2. On each step of the loop print a random character.
//  3. And, sleep for 250 milliseconds.
//  4. Run the program and wait a couple of seconds
//     then kill it using CTRL+C keys
//
// RESTRICTIONS
//  1. Print one of those characters randomly: \ / | -
//  2. Before printing a character print also this
//     escape sequence: \r
//
//     Like this: "\r/", or this: "\r|", and so on...
//
//  3. NOTE : If you're using Go Playground, use "\f" instead of "\r"
//
// HINTS
//  1. Use `time.Sleep` to sleep.
//  2. You should pass a `time.Duration` value to it.
//  3. Check out the Go online documentation for the
//     `Millisecond` constant.
//  4. When printing, do not use a newline! Or a Println!
//     Use Print or Printf instead.
//
// NOTE
//  If this exercise is hard for you now, wait until the
//  lucky number lecture. Even then so, then just skip it.
//
//  You can return back to it afterwards.
//
// EXPECTED OUTPUT
//  - The program should display the following messages in any order.
//  - And, the first character (\, -, /, or |) should be randomly
//    displayed.
//  - \r or \f sequence will clear the line.
//
//  \ Please Wait. Processing....
//  - Please Wait. Processing....
//  / Please Wait. Processing....
//  | Please Wait. Processing....
// ---------------------------------------------------------

func main() {
}
