## 当你把一个变量赋值给另一个变量时会发生什么？

* 变量们变成同一个
* 被赋值的变量被删除
* 变量的值被改变为赋值变量的值 *正确*

## 哪个是正确的赋值语句?

```go
opened := true
```

* `closed := true`
* `opened = false` *正确*
* `var x = 2`

## 哪个是正确的赋值语句?

* `a, b = 3; 5`
* `c, d = true, false` *正确*
* `a, b, c = 5, 3`

## 哪个是正确的赋值语句?

```go
var (
  n = 3
  m int
)
```

* `m = "4"`
* `n = 10` *正确*
* `n = true`
* `m = false`

## 哪个是正确的赋值语句?

```go
var (
  n = 3
  m int
  f float64
)

// one of the assignments below will be here

fmt.Println(n, m, f)
```

* `n, m = 4, f`
* `n = false`
* `n, m, f = m + n, n + 5, 0.5` *正确*
* `n, m = 3.82, "hi"`

## 哪个是正确的赋值语句?

```go
var (
  a int
  c bool
)
```

* `a = _`
* `c, _ = _, false`
* `_, _ = a, c` *正确*

## 哪个是正确的赋值语句?

**记住:** `path.Split` 返回两个 `string` 值

```go
  var c, f string
```

* `_ = path.Split("assets/profile.png")`
* `_, _, c = path.Split("assets/profile.png")`
* `f = path.Split("assets/profile.png")`
* `_, f = path.Split("assets/profile.png")` *正确*

