## 为什么有时需要用到注释？
1. 为了将不同的表达式组合起来
2. 为了给代码提供说明或者自动生成文档 *正确*
3. 为了让代码看起来更漂亮


## 下面的代码哪个是正确的？
1.
```go
package main

/ main function is an entry point /
func main() {
    fmt.Println("Hi")
}
```

2. *正确*
```go
package main

// main function is an entry point /*
func main() {
    fmt.Println(/* this will print Hi! */ "Hi")
}
```

3.
```go
package main

/*
main function is an entry point

It allows Go to find where to start executing an executable program.
*/
func main() {
    fmt.Println(// "this will print Hi!")
}
```

> **1:** `/` 不是注释，单行注释是 `//`.
>
>
> **2:** 多行注释可以放在任何位置。但是，当有一个以 `/*`开头的注释，需要以 `*/` 结尾。 这里，Go 不关心 `/* ... */` 里面的内容， 它直接跳过。同时，Go 看到以 `//` 开头的代码，直接跳过整行代码。
>
>
> **3:** `//` 阻止 Go 解析后面的内容，这就是它不能运行的原因， Go 不能解析 `"this will print Hi!")` 因为它们是注释。
>
>

## 应该如何命名代码，以便 Go 可以从代码自动生成文档？
1. 通过注释的每一行代码，这样就会自动生成文档
2. 通过在声明变量后开始进行注释 *正确*
3. 通过使用多行注释

> **1:** 对不起，这没用
>
>
> **3:** 它不关心使用单行注释还是多行注释
>
>


## 在命令行打印文档需要用到下面哪个？
1. go build
2. go run
3. go doctor
4. go doc *正确*


## `godoc` 和 `go doc` 有什么不同?
1. `go doc` 是 `godoc` 背后真实的工具
2. `godoc` 是 `go doc` 背后真实的工具 *正确*
3. `go` 是 `go doc` 背后真实的工具
4. `go` 是 `godoc` 背后真实的工具

> **2:** 是的， go doc 背后使用 godoc， go doc 只是简单版本的 godoc
>
>
