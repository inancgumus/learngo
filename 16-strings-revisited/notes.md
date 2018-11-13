# Strings Revisited

## Bytes
* ASCII
* Immutable

## Runes
* Unicode
* vs ASCII
* UTF-8
* Made up of bytes

## Slicing
* String: Read-Only Byte Slice
* Slicing -> String
* Index   -> Byte

---

## Read-only byte slice
* A string is a read-only slice
* You can't change its data
* Indexable: Returns you a byte
* Slicable: Returns you a string

## Slicing
* Strings can be sliced just like a slice
* After slicing Go returns you a new string slice
* WARNING: Indexing expression returns you a byte
    * s := "hey"
    * s[0] + s[1] + s[2] != "hey"
    * s[0:3] == "hey

## Underlying array
* Underlying array is a string array
* There's no capacity this time: Only length and pointer.

* Sliced string will refer to that array
* String slicing is cheap — They share the same array

## Unicode
* At the beginning there was only ASCII code standard
    * It was using 7-bits to represents 128 characters
    * Only English characters
    * Each code corresponding to a character

* After Internet nothing couldn't stay the same
* There was a need to introduce more languages
* 127 characters aren't enough for the entire world

* So: Unicode is born
* It collects all of the characters in world's languages
* Unicode can represent every character in every imaginable language system.

* Assigns each character to a codepoint or a rune (in Go)
* Unicode assigns each character a unique number, or code point.

* Codepoint is a numeric number which represents a character in general
* U+2700 -> hex
* Unicode defines codepoints for 1m+ characters 
* It includes the ASCII codes too

    A chinese character:      汉
    Its unicode value:        U+6C49
    convert 6C49 to binary:   01101100 01001001
    embed 6C49 as UTF-8:      11100110 10110001 10001001


## Unicode and Runes
* Rune is a 4-bytes type for storing unicode codepoints

* Rune data type and rune codepoints are different things!

* There's UTF-32 standard which assigns 4 bytes to each codepoint
* But, that's inefficient, so, instead Go uses a variable encoding standard called UTF-8. It assigns different number of bytes to codepoints.
* UTF-8 has been invented by Rob Pike and Ken Thompson (two of the creators of Go)

* So, a rune is 1-4 bytes. Uses 1 byte for ASCII (english).
* 2-3 bytes for most of the characters.

* A string can contain runes
* Each rune can span to multiple bytes
* WARNING: Getting one byte of a string may give you corrupt data
    * If you're getting one part of a rune inside the string!

* In a string with runes, you can't easily index the characters
    * You need to use unicode and utf8 packages
    * Or you need to convert the string into a rune slice
    * unicode: letters vs nums, to uppercase, ...
    * utf8   : working w/bytes and runes

* RuneCountInString(s) == len([]rune(s))
* DecodeRuneInString(s) returns the first rune

## Ranging over strings
* You can range over a string like a slice
* It will jump over the runes inside the string
* The index variable will be the starting position of each rune
    * And the value will be the rune itself

## Representing bytes
* Unicode characters can be hard to type in code
* So, we can use \x and \u in a string to represent bytes and runes

* A string literal is always utf-8 but a string value is not

## Convenience
* It's easy to work with runes in code: []rune
* However, it will consume more memory: Each char is 4 bytes

* "inanç"[4] = gibberish

* r := []rune("inanç") -> five elements rune slice
    * r[4] = 'ç'
    * string(r)
    * // inanç: automatically concatenates the runes to form a string

* string(105) // i -> interprets 105 as a rune value; 'i' not 105
* string(351) // ş -> ""

* printf: %q -> 'ç' %c -> ç %d -> 231

## Bytes
* major libs:
    * strings, bytes (have corresponding funcs)
    * strconv, unicode
    * bytes.Buffer

* []byte can be modified whereas string is immutable
    * if you do a lot of string manipulations you can use []byte

* []byte <-> string convertable
    * but, each conversion copies the data
    * compiler optimizes it mostly

    * however, do not blindly convert; use bytes pkg
    * it's like the string pkg

    * s := "inanc"
    * b := []byte(s)
    * s := string(b)

## Sprintf
* Just like printf but instead of printing it returns a string

## Builders
* bytes.Buffer
* strings.Builder

* Use WriteRune when adding rune


## Terminology:
Summary: Unicode is a large table mapping characters to numbers and the different UTF encodings specify how these numbers are encoded as bits. 

* **ASCII** First character set that maps characters to codepoints or character codes. In terms of alphabets, it only supports basic latin alphabet: English. 2^7=127

    * The center of the computer industry was in the USA at that time. As a consequence, they didn't need to support accents or other marks such as á, ü, ç, ñ, etc.

    * Once upon a time, computer memory and storage was very expensive. And all of the computers in the world (for practical purposes) were in the hands of English-speaking countries. 

    * Single byte encoding only using the bottom 7 bits. Basic Latin. (Unicode code points 0-127.) No accents etc.

* **Unicode** is a coded character set. A set of characters and a mapping between the characters and integer code points representing them. Unicode is a superset of ASCII.

    * You cannot save text to your hard drive as "Unicode". Unicode is an abstract representation of the text. You need to "encode" this abstract representation. That's where an encoding comes into play.

    * Unicode first and foremost defines a table of code points for characters. That's a fancy way of saying "65 stands for A, 66 stands for B and 9,731 stands for ☃" (seriously, it does). How these code points are actually encoded into bits is a different topic. 

* **UTF-8** is a character encoding - a way of converting from sequences of bytes to sequences of characters and vice versa. It covers the whole of the Unicode character set.

    * UTF-8 uses the ASCII set for the first 128 characters. That's handy because it means ASCII text is also valid in UTF-8.

* **Character Set:** A character set is a list of characters with unique numbers (these numbers are sometimes referred to as “code points”). For example, in the Unicode character set, the number for A is 41.

* **Codepoint:** Characters are referred to by their "Unicode code point". 

    * Written in hexadecimal (to keep the numbers shorter).

    * Preceded by a "U+" (that's just what they do, it has no other meaning than "this is a Unicode code point").

    * Unicode itself is a mapping, it defines codepoints and a codepoint is a number, associated with usually a character.

    * Code: a system of words, letters, figures, or other symbols substituted for other words, letters, etc.

* **Encoding:** Converting data into a coded form. An encoding on the other hand, is an algorithm that translates a list of numbers to binary so it can be stored on disk. For example UTF-8 would translate the number sequence 1, 2, 3, 4 like this: `00000001 00000010 00000011 00000100`. Our data is now translated into binary and can now be saved to disk.

    * To encode means to use something to represent something else. An encoding is the set of rules with which to convert something from one representation to another.

    * To represent 1,114,112 different values, two bytes aren't enough. Three bytes are, but three bytes are often awkward to work with, so four bytes would be the comfortable minimum. But, unless you're actually using Chinese or some of the other characters with big numbers that take a lot of bits to encode, you're never going to use a huge chunk of those four bytes.

    * If the letter "A" was always encoded to 00000000 00000000 00000000 01000001, "B" always to 00000000 00000000 00000000 01000010 and so on, any document would bloat to four times the necessary size.

    * To optimize this, there are several ways to encode Unicode code points into bits. UTF-8 is one of them.

    character	encoding	bits
    A	        UTF-8	    01000001
    A	        UTF-16	    00000000 01000001
    A	        UTF-32	    00000000 00000000 00000000 01000001

    U+0000 to U+007F are (correctly) encoded with one byte
    U+0080 to U+07FF are encoded with 2 bytes
    U+0800 to U+FFFF are encoded with 3 bytes
    U+010000 to U+10FFFF are encoded with 4 bytes

    * There is NO string or text, without an accompanying encoding standard.

## REFS:
https://unicode-table.com/en/

What's the difference between ASCII and Unicode?
https://stackoverflow.com/a/41198513/115363

https://stackoverflow.com/questions/643694/what-is-the-difference-between-utf-8-and-unicode

https://stackoverflow.com/questions/3951722/whats-the-difference-between-unicode-and-utf-8

https://stackoverflow.com/questions/1543613/how-does-utf-8-variable-width-encoding-work

http://kunststube.net/encoding/
(detailed and simple)

http://www.joelonsoftware.com/articles/Unicode.html

Unicode codepoint to UTF-8 encoding answer: https://stackoverflow.com/a/27939161/115363

http://www.polylab.dk/utf8-vs-unicode.html

Characters, Symbols and the Unicode Miracle - Computerphile
https://www.youtube.com/watch?v=MijmeoH9LT4

The history of UTF-8 as told by Rob Pike
http://doc.cat-v.org/bell_labs/utf-8_history