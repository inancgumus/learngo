# Strings, Runes and Bytes Quiz

## Which byte slice below equals to the "keeper" string?
```go
// Here are the corresponding code points for the runes of "keeper":
// k => 107
// e => 101
// p => 112
// r => 114
```
1. []byte{107, 101, 101, 112, 101, 114} *CORRECT*
2. []byte{112, 101, 101, 112, 114, 101}
3. []byte{114, 101, 112, 101, 101, 112}
4. []byte{112, 101, 101, 114, 107, 101}


## What does this code print?
```go
// Code points:
// g => 103
// o => 111
fmt.Println(string(103), string(111))
```
1. 103 111
2. g o *CORRECT*
3. n o
4. "103 111"


## What does this code print?
```go
const word = "gökyüzü"
bword := []byte(word)

// ö => 2 bytes
// ü => 2 bytes
fmt.Println(utf8.RuneCount(bword), len(word), len(string(word[1])))
```
1. 7 10 2 *CORRECT*
2. 10 7 1
3. 10 7 2
4. 7 7 1


## Which one below is true?
1. for range loops over the bytes of a string
2. for range loops over the runes of a string *CORRECT*


## For a utf-8 encoded string value, which one below is true?
1. runes always start and end in the same indexes
2. runes may start and end in different indexes *CORRECT*
3. bytes may start and end in different indexes



## Why can't you change the bytes of a string value?
1. Strings values are immutable byte slices
2. Strings are used a lot so they are being shared behind the scenes
3. All of above *CORRECT*