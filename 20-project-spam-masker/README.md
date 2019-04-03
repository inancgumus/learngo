# Spam Masker Challenge Tips

## Rules:

* You shouldn't use a standard library function.

* You should only solve the challenge by manipulating the bytes directly.

* Manipulate the bytes of a string using indexing, slicing, appending etc.

* Be efficient: Do not use string concat (+ operator).
	* Instead, create a new byte slice as a buffer from the given string argument.
	* Then, manipulate it during your program.
	* And, for once, print that buffer.

* Mask only links starting with `http://`

* Don't check for uppercase/lowercase letters

	* The goal is learning how to manipulate bytes in strings, it's not about creating the perfect masker.

	* For example: A spammer can prevent the masker like this (for now this is OK):

	  ```
      "Here's my spammy page: hTTp://youth-elixir.com"
                               ^^
      ```

* But, you should catch this:

  ```
  INPUT:
  Here's my spammy page: http://hehefouls.netHAHAHA see you.

  OUTPUT:
  Here's my spammy page: http://******************* see you.
  ```

## Steps:

1. Check whether there's a command line argument or not. If not, quit from the program with a message.

2. Create a byte buffer as big as the argument.

3. Loop and detect the `http://` patterns

4. Copy the input character by character to the buffer

5. If you detect `http://` pattern, copy the `http://` first, then copy the `*`s instead of the original link until you see whitespace character.

    For example: 
    ```
	INPUT:
    Here: http://www.mylink.com Click!

    OUTPUT:
	Here: http://************** Click!
    ```

6. Print the buffer as a string