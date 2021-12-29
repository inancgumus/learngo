## 属于同一个包的源代码存储在哪里？
1. 每个文件应该存储在不同的目录
2. 在同一个目录下 *正确*


## 为什么 Go 源代码需要使用 package 语句？
1. 用来导入包
2. 用来让 Go 知道该文件属于哪个包 *正确*
3. 用来声明新函数

> **1:** `import` 语句的作用.
>
>
> **3:** `func` 语句的作用
>
>


## `package clause` 需要写在 Go 源文件的哪里？
1. 写在 Go 源文件的开头 *正确*
2. 写在 Go 源文件的末尾
3. 写哪里都可以


## 一个 Go 源文件使用 `package clause` 几次
1. 一次 *正确*
2. 零次
3. 多次


## 下面哪一个正确使用 `package clause`?
1. `my package`
2. `package main`
3. `pkg main`


## 哪一个说法是正确?
1. 在同一个包下的文件不能调用其他文件定义的函数
2. 在同一个包下的文件可以调用其他文件定义的函数 *正确*


## 如何运行多个 Go 源代码
1. go run *.*go
2. go build *go
3. go run go
4. go run *.go *正确*

> **4:** 也可以是 (假设目录下有 file1.go file2.go 和 file3.go): go run file1.go file2.go file3.go
>
>
