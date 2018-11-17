package main

// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
// ~GET SOCIAL~
//
// Tweet when you complete:
// http://bit.ly/GOTWEET-CLOCK
//
// Discuss it with other gophers:
// http://bit.ly/LEARNGOSLACK
// ➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

func main() {
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// Notes
	// + I've created the solutions step by step
	// + So, if you get stuck, you can check out the next step
	//   without having to look at the entire final solution
	// + For drawing the clock, you may use the artifacts below
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// ★ Usable Artifacts
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

	// Clock Characters:
	//
	//   You can put these in constants if you like.
	//
	//   Use this for the digits       : "█"
	//   Use this for the separators   : "░"

	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// ★ GOAL 1  : Printing the Digits
	// ★ Solution: 02-printing-the-digits
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

	// - [ ] Define a new placeholder type
	// ...

	// - [ ] Create the digits
	// zero := ...
	// one := ...
	// ...

	// - [ ] Put them into the "digits" array
	// digits := ...

	// - [ ] Print them side-by-side
	//       - [ ] Loop for all the lines in a digit
	//       - [ ] Print each digit line by line
	//
	//       - [ ] Don't forget printing a space after each digit
	//       - [ ] Don't forget printing a newline after each line
	//
	// EXAMPLE: Let's say you want to print 10.
	//
	//       ██   ███<--- Print a new line after printing a single line from
	//        █   █ █     all the digits.
	//        █   █ █
	//        █   █ █
	//       ███  ███
	//          ^^
	//          ||
	//          ++----> Add space between the digits
	//

	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// ★ GOAL 2  : Printing the Clock
	// ★ Solution: 03-printing-the-clock
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

	// - [ ] Get the current time

	// - [ ] Get the current hour, minute and second from the current time

	// - [ ] Create the clock array
	//
	//       - [ ] Get the individual digits from the digits array
	//
	// clock := ...

	// - [ ] Print the clock
	//
	//       - [ ] In the loops, use the clocks array instead

	// - [ ] Create a separator array (it's also a placeholder)

	// - [ ] Add the separators into the correct positions of
	//       the clock array

	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// ★ GOAL 3  : Updating the Clock
	// ★ Solution: 04-updating-the-clock
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

	// - [ ] Create an infinite loop to update the clock

	// - [ ] Update the clock every second
	//
	//       time.Sleep(time.Second) will stop the world for 1 second
	//
	//       See this for more info:
	//       https://golang.org/pkg/time/#Sleep

	// - [ ] Clear the screen before the infinite loop
	//
	//       This string value clears the screen when printed:
	//
	//         "\033[2J"
	//
	//       + For Go Playground, use this instead: "\f"
	//
	//       + This only works on bash command prompts.

	// - [ ] Move the cursor to the top-left corner of the screen before each step
	//       of the infinite loop
	//
	//       This string value moves the cursor like so when printed:
	//
	//         "\033[H"
	//
	//       + For Go Playground, use this instead: "\f"
	//
	//       + This only works on bash command prompts.

	//       + If you're curious:
	//         \033 is a special control code:
	//         [2J clears the screen and the cursor
	//         [H  moves the cursor to 0, 0 screen position
	//
	//         See for more info:
	//         https://bluesock.org/~willkg/dev/ansi.html

	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// ★ GOAL 4: Blinking the Separators
	// ★ Solution: 05-blink-the-separators
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★

	// - [ ] Blink the separators
	//
	//       They should be visible per two seconds.
	//
	//       Example: 1st second invisible
	//                2nd second visible
	//                3rd second invisible
	//                4th second visible
	//
	// HINT: There are two ways to do this.
	//
	//       1- Manipulating the clock array directly
	//          (by adding/removing the separators)
	//
	//       2- Deciding what to print when printing the clock

	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
	// YOU CAN FIND THE FINAL SOLUTION WITH ANNOTATIONS:
	// 06-full-annotated-code
	// ★★★★★★★★★★★★★★★★★★★★★★★★★★★
}
