## Which expression increases `n` by 1?
```go
var n float64
```
1. `n = +1`
2. `n = n++`
3. `n = n + 1` *CORRECT*
4. `++n`

> **1:** This just assigns 1 to n.
>
> **2:** IncDec statement can't be used as an operator.
>
> **4:** Go doesn't support prefix incdec notation.
>


## Which expression decreases `n` by 1?
```go
var n int
```
1. `n = -1`
2. `n = n--`
3. `n = n - 1` *CORRECT*
4. `--n`

> **1:** This just assigns -1 to n.
>
> **2:** IncDec statement can't be used as an operator.
>
> **4:** Go doesn't support prefix incdec notation.
>


## Which code below equals to `n = n + 1`?
1. `n++` *CORRECT*
2. `n = n++`
3. `++n`
4. `n = n ++ 1`

> **2:** IncDec statement can't be used as an operator.
>
> **3:** Go doesn't support prefix incdec notation.
>
> **4:** What's that? ++?
>


## Which code below equals to `n = n + 1`?
1. `n = n++`
2. `n += 1` *CORRECT*
3. `++n`
4. `n = n ++ 1`

> **1:** IncDec statement can't be used as an operator.
>
> **3:** Go doesn't support prefix incdec notation.
>
> **4:** What's that? ++?
>


## Which code below equals to `n -= 1`?
1. `n = n--`
2. `n += 1--`
3. `n--` *CORRECT*
4. `--n`

> **1:** IncDec statement can't be used as an operator.
>
> **2:** IncDec statement can't be used as an operator. And also, you can't use it with `1--`. The value should be addressable. You're going to learn what that means soon.
>
> **4:** Go doesn't support prefix incdec notation.
>


## Which code below divides the `length` by 10?
1. `length = length // 10`
2. `length /= 10` *CORRECT*
3. `length //= 10`

> **1:** What's that? `//`?
>
> **2:** That's right. This equals to: `length = length / 10`
>
> **3:** What's that? `//=`?
>


## Which code below equals to `x = x % 2`?
1. `x = x / 2`
2. `x =% 2`
3. `x %= 2` *CORRECT*

> **1:** This is a division. You need to use the remainder operator.
>
> **2:** Close... But, the `%` operator is on the wrong side of the assignment.
>


## Which function below converts a string value into a float value?
1. `fmtconv.ToFloat`
2. `conv.ParseFloat`
3. `strconv.ParseFloat` *CORRECT*
4. `strconv.ToFloat`


## Which code is correct?
If you don't remember it, this its function signature:
```go
func ParseFloat(s string, bitSize int) (float64, error)
```
1. `strconv.ParseFloat("10", 128)`
2. `strconv.ParseFloat("10", 64)` *CORRECT*
3. `strconv.ParseFloat("10", "64")`
4. `strconv.ParseFloat(10, 64)`

> **1:** There are no 128-bit floating point values in Go (Actually there are, but they only belong to the compile-time).
>