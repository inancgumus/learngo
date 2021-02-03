## Which one is **not** a predeclared data type of Go?
1. int
2. float64
3. uint64
4. uint
5. duration *CORRECT*
6. int8
7. rune
8. byte
9. float32
10. complex128


## What's a predeclared data type?
1. A data type used only in the compiler
2. A built-in data type that comes with Go that you can use it from anywhere without importing any package *CORRECT*
3. The data type of a variable


## By using only 8 bits, how many different numbers or states can you represent?
1. 8
2. 16
3. 256 *CORRECT*
4. 65536

> **3:** 2^8 is 256, so you can represent 256 different states.
>


## How many bits 2 bytes contains?
1. 2
2. 8
3. 16 *CORRECT*
4. 32
5. 64


## What's the output of the following code?
```go
fmt.Printf("%08b = %d", 2, 2)
```
1. 00000001 = 2
2. 00000010 = 2 *CORRECT*
3. 00000100 = 2
4. 00001000 = 2

> EXPLANATION = From right to left, each bit goes from 2^0 to 2^(n - 1).

> **1:** EXPLANATION. Here: 1 is the first digit from the right. So, it is 2^(1 - 1) = 2^0 = 1.
>
> **2:** EXPLANATION. Here: 1 is the second digit from the right. So, it is 2^(2 - 1) = 2^1 = 2.
>
> **3:** EXPLANATION. Here: 1 is the third digit from the right. So, it is 2^(3 - 1) = 2^2 = 4.
>
> **4:** EXPLANATION. Here: 1 is the fourth digit from the right. So, it is 2^(4 - 1) = 2^3 = 8.
>


## How many bytes of memory does an int64 value use?
1. 4
2. 8 *CORRECT*
3. 32
4. 64

> **2:** 1 byte is 8 bits and int64 is 64 bits. So, 64/8=8 bytes.
>


## How many bytes are needed to store a value of uint32 type?
1. 4 *CORRECT*
2. 8
3. 32
4. 64

> **2:** 1 byte is 8 bits and uint32 is 32 bits. So, 32/8=4 bytes.
>


## What's the size of int data type?
1. Depends: 32 bits or 64 bits. *CORRECT*
2. 32 bits
3. 64 bits

> **1:** That's right. Go can change its size at the compile-time depending on which target machine you're compiling your program into.
>


## English letters can be represented by the numbers within the range of: 0-255. For example, 'A' can be 65. Or, 'B' can be 66. So, what's the best data type for storing an English letter?

1. byte *CORRECT*
2. rune
3. int64
4. float64

> **1:** That's right. A byte can represent 0-255 different values. So, it's a great fit for representing English letters, and numbers.
>
> **2:** In practice, you can do it with a rune value. However, rune is 32-bits long and it can store almost every letter in existince. I'm asking for the optimal data type. Try again.
>
> **3:** That would be too large for only 255 different numbers.
>
> **4:** Float is not the best data type for numbers without fractional parts.
>



## What does the following code print?
```go
var letter uint8 = 255
fmt.Print(letter + 5)
```
1. 0
2. 4 *CORRECT*
3. 5
4. 260

> **2:** Unsigned integers wrap around after their maximum capacity. Uint8's max is 255, so, if 255 + 1 is 0, then 255 + 5 is 4.
>
> **3:** You're very close.
>
> **4:** Uint8's max capacity is 255. It can't be 260.
>


## What does the following code print?
```go
var num int8 = -128
fmt.Print(num - 3)
```
1. -131
2. 125 *CORRECT*
3. -125

> **1:** int8's min capacity is -128. It can't be -131.
>
> **2:** Signed integers wrap around after their minimum capacity. int8's min is -128, so, if -128 - 1 is 127, then -128 - 3 is 125.
>
