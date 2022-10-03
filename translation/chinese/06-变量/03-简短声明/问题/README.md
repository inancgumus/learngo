## 哪个是正确的声明?
* var int safe := 3
* var safe bool := 3
* safe := true *正确*

## 哪个是正确的声明?
* var short := true
* int num := 1
* speed := 50 *正确*
* num int := 2

## 哪个是正确的声明?
* x, y, z := 10, 20
* x = 10,
* y, x, p := 5, "hi", 1.5 *正确*
* y, x = "hello", 10

## 哪个声明和下面的声明等价?
```go
var s string = "hi"
```

* var s int = "hi"
* s := "hi" *正确*
* s, p := 2, 3

## 哪个声明和下面的声明等价?
```go
var n = 10
```

* n := 10.0
* m, n := 1, 0
* var n int = 10 *正确*

## 变量 `s` 的类型是什么?
```go
s := "hmm..."
```

* bool
* string *正确*
* int
* float64

## 变量 `b` 的类型是什么?
```go
b := true
```

* bool *正确*
* string
* int
* float64

## 变量 `i` 的类型是什么?
```go
i := 42
```

* bool
* string
* int *正确*
* float64

## 变量 `f` 的类型是什么?
```go
f := 6.28
```

* bool
* string
* int
* float64 *正确*

## 变量 `x` 的类型是什么?

```go
y, x := false, 20
```

* 10
* 20 *正确*
* false

## 变量 `x` 的类型是什么?

```go
y, x := false, 20
x, z := 10, "hi"
```

* 10 *正确*
* 20
* false

## 下面哪个声明能在包作用域上使用?

* x := 10
* y, x := 10, 5
* var x, y = 5, 10 *正确*
