## `os.Args` 变量在其第一项中存储了什么?

* 传递给程序的第一个参数
* 传递给程序的第二个参数
* 运行程序的路径 *正确*

## `Args` 变量的类型是什么？

```go
var Args []string
```

* string
* string 数组
* strings 的切片 *正确*

## `Args` 变量中每个值的类型是什么？

```go
var Args []string
```

* string *正确*
* string 数组
* strings 的切片

## 如何获得 `Args` 变量的第一项？

```go
var Args []string
```

* Args.0
* Args{1}
* Args[0] *正确*
* Args(1)

## 如何获得 `Args` 变量的第二项？

```go
var Args []string
```

* Args.2
* Args[1] *正确*
* Args{1}
* Args(2)

## 如何获得 `Args` 变量的长度？

```go
var Args []string
```

* length(Args)
* Args.len
* len(Args) *正确*
* Args.Length

## 如何从命令行获得第一个 "argument"？

* os.Args[0]
* os.Args[1] *正确*
* os.Args[2]
* os.Args[3]

