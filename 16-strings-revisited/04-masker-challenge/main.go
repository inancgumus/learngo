// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

/*
------------------------------------------------------------
RULES
------------------------------------------------------------

* You shouldn't use a standard library function.

* You should only solve the challenge by manipulating the bytes directly.

* Manipulate the bytes of a string using indexing, slicing, appending etc.

* Be efficient: Do not use string concat (+ operator).

	* Instead, create a new byte slice as a buffer from the given string argument.

	* Then, manipulate it during your program.

	* And, for once, print that buffer.

------------------------------------------------------------
STEPS
------------------------------------------------------------

* Mask only links starting with http://

* Don't check for uppercase/lowercase letters
	* The goal is to learn manipulating bytes in strings
	* It's not about creating a perfect masker

	* For example: A spammer can prevent the masker like this (for now this is OK):

	  "Here's my spammy page: hTTp://youth-elixir.com"

* But, you should catch this:

	  "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	  "Here's my spammy page: http://******************* see you."
*/
package main

func main() {
	// ---------------------------------------------------------------
	// #1
	// ---------------------------------------------------------------
	// Check whether there's a command line argument or not
	// If not, quit from the program with a message

	// ---------------------------------------------------------------
	// #2
	// ---------------------------------------------------------------
	// Create a byte buffer as big as the argument

	// ---------------------------------------------------------------
	// #3
	// ---------------------------------------------------------------
	// 1- Loop and detect the http:// patterns
	// 2- Copy the input character by character to the buffer
	// 3- If you detect http:// pattern, copy the http:// first,
	//    then copy the *s instead of the original link until
	//    you see a whitespace character.
	//
	//    Here: http://www.mylink.com Click!
	// -> Here: http://************** Click!
	//

	// ---------------------------------------------------------------
	// #4
	// ---------------------------------------------------------------
	// Print the buffer as a string
}
