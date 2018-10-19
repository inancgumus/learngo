## What's the result of this expression?
```go
"\"Hello\\"" + ` \"World\"`
```

1. "Hello" "World"
2. "Hello" \"World\" *CORRECT*
3. "Hello" `"World"`
4. "\"Hello\" `\"World\"`"

> **1:** Go doesn't interpret the escape sequences in raw string literals.
>
> **2:** That's right. Go interprets `\"` as `"` but it doesn't do so for ` \"World\"`.
>


## What's the best way to represent the following text in the code?
```xml
<xml>
  <items>
    <item>"Teddy Bear"</item>
  </items>
</xml>
```

1. *CORRECT*
```go
`<xml>
    <items>
        <item>"Teddy Bear"</item>
    </items>
</xml>`
```

2. 
```go
"<xml>
    <items>
        <item>"Teddy Bear"</item>
    </items>
</xml>"
```

3. 
```go
"<xml>
    <items>
        <item>\"Teddy Bear\"</item>
    </items>
</xml>"
```

4. 
```go
`<xml>
    <items>
        <item>\"Teddy Bear\"</item>
    </items>
</xml>`
```

> **2-3:** You can't write a string literal like that. It can't be multiple-lines.
>
> **4:** You don't need to use escape sequences inside raw string literals.
>


## What's the result of the following expression?
```go
len("lovely")
```

1. 7
2. 8
3. 6 *CORRECT*
4. 0

> **2:** Remember! "a" is 1 char. `a` is also 1 char.
>


## What's the result of the following expression?
```go
len("very") + len(`\"cool\"`)
```

1. 8
2. 12 *CORRECT*
3. 16
4. 10

> **1:** There are also double-quotes, count them as well.
>
> **2:** That's right. Go doesn't interpreted \" in raw string literals.
>
> **3:** Remember! "very" is 4 characters. `very` is also 4 characters.
>
> **4:** Remember! Go doesn't interpreted \" in raw string literals.
>


## What's the result of the following expression?
```go
len("very") + len("\"cool\"")
```

1. 8
2. 12
3. 16
4. 10 *CORRECT*

> **1:** There are also double-quotes, count them as well.
>
> **2:** Remember! Go interprets escape sequences in string literals.
>
> **4:** That's right. Go does interpret \" in a string literal. So, "\"" means ", which is 1 character.
>


## What's the result of the following expression?
```go
len("péripatéticien")
```

**HINT:** é is 2 bytes long. And, the len function counts the bytes not the letters.

**USELESS INFORMATION:** "péripatéticien" means "wanderer".

1. 14
2. 16 *CORRECT*
3. 18
4. 20

> **1:** Remember! é is 2 bytes long.
>
> **2:** An english letter is 1 byte long. However, é is 2 bytes long. So, that makes up 16 bytes. Cool.
>
> **3:** You didn't count the double-quotes, did you?
>


## How can you find the correct length of the characters in this string literal?
```go
"péripatéticien"
```

1. `len(péripatéticien)`
2. `len("péripatéticien")`
3. `utf8.RuneCountInString("péripatéticien")` *CORRECT*
4. `unicode/utf8.RuneCountInString("péripatéticien")`

> **1:** Where are the double-quotes?
>
> **2:** This only finds the bytes in a string value.
>
> **4:** You're close. But, the package's name is utf8 not unicode/utf8.
>


## What's the result of the following expression?
```go
utf8.RuneCountInString("péripatéticien")
```

1. 16
2. 14 *CORRECT*
3. 18
4. 20

> **1:** This is its byte count. `RuneCountInString` counts the runes (codepoints) not the bytes.
>
> **2:** That's right. `RuneCountInString` returns the number of runes (codepoints) in a string value.
>


## Which package contains string manipulation functions?
1. string
2. unicode/utf8
3. strings *CORRECT*
4. unicode/strings


## What's the result of this expression?
```go
strings.Repeat("*x", 3) + "*"
```

**HINT:** Repeat function repeats the given string.

1. `*x*x*x`
2. `x*x*x`
3. `*x3`
4. `*x*x*x*` *CORRECT*

> **1:** You're close but you missed the concatenation at the end.
>
> **2:** Look closely.
>
> **3:** Wow! You should really watch the lectures again. Sorry.
>
> **4:** That's right. Repeat function repeats the given string. And, the concatenation operator combines the strings.
>


## What's the result of this expression?
```go
strings.ToUpper("bye bye ") + "see you!"
```

1. `bye bye see you!`
2. `BYE BYE SEE YOU!`
3. `bye bye + see you!`
4. `BYE BYE see you!` *CORRECT*

> **1:** You missed the ToUpper?
>
> **2:** You're close but look closely. ToUpper only changes the first part of the string there.
>
> **3:** Not even close. Sorry.
>
> **4:** Perfect! Good catch! ToUpper only changes the first part of the string there.
>