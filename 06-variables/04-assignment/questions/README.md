## What happens when you assign a variable to another variable?
* The variables become the same variable
* Assignee variable gets deleted
* Variable's value is changed to the assigned variable's value *CORRECT*

## Which one is a correct assignment statement?
```go
opened := true
```
* `closed := true`
* `opened = false` *CORRECT*
* `var x = 2`

## Which one is a correct assignment statement?
* `a, b = 3; 5`
* `c, d = true, false` *CORRECT*
* `a, b, c = 5, 3`

## Which one is a correct assignment?
```go
var (
  n = 3
  m int
)
```

* `m = "4"`
* `n = 10` *CORRECT*
* `n = true`
* `m = false`

## Which one is a correct assignment?
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
* `n, m, f = m + n, n + 5, 0.5` *CORRECT*
* `n, m = 3.82, "hi"`

## Which one is a correct assignment statement?
```go
var (
  a int
  c bool
)
```
* `a = _`
* `c, _ = _, false` 
* `_, _ = a, c` *CORRECT*

## Which one is a correct assignment?
**REMEMBER:** `path.Split` returns two `string` values

```go
  var c, f string
```
* `_ = path.Split("assets/profile.png")`
* `_, _, c = path.Split("assets/profile.png")`
* `f = path.Split("assets/profile.png")`
* `_, f = path.Split("assets/profile.png")` *CORRECT*
