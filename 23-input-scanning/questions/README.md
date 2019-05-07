## What is an input stream?
1. Any data source that generates contiguous data like Standard Input *CORRECT*
2. An input from a user
3. The contents of a file
4. The contents of a website

> **1:** That's right. The input stream can come from any data source. Standard Input is one of them, and you can redirect it to almost any data source like user input, files, website contents and so on.
>
> **2-4:** Yes, that may be an input stream, but it doesn't explain what an input stream is.
>


## What does the program print?
```go
in := bufio.Scanner(os.Stdin)
in.Scan() // user enters: "hi!"
in.Scan() // user enters: "how are you?"
fmt.Println(in.Text())
```
1. "hi" and "how are you?"
2. "hi"
3. "how are you?" *CORRECT*
4. Nothing

> **3:** The Text() method only returns the last scanned token. A token can be a line or a word and so on.
>


## Using bufio.Scanner, how can you detect errors in the input stream?
1. Using the Err() method *CORRECT*
2. Using the Error() method
3. By checking whether the Scanner is nil or not


## How can you configure bufio.Scanner to only scan for the words?
```go
in := bufio.Scanner(os.Stdin)
// ...
```
1. `in = bufio.NewScanner(in, bufio.ScanWords)`
2. `in.Split(bufio.ScanWords)` *CORRECT*
3. `in.ScanWords()`

> **2:** That's right. bufio has a few splitters like ScanWords such as ScanLines (the default), ScanRunes, and so on.
>


## The function uses the "Must" prefix, why?
```go
regexp.MustCompile("...")
```
1. It's only being used for readability purposes
2. It's a guarantee that the function will work, no matter what
3. The function may crash your program *CORRECT*

> **3:** "Must" prefix is a convention. If a function or method may panic (= crash a program), then it's usually being prefixed with a "must" prefix.
>