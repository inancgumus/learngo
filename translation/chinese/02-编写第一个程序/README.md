# 备忘录: 编写第一个 GO 程序

嗨！

作为参考，你可以在参加本节讲座后保存这个备忘录

你也可以打印下来这个备忘录，然后和本节的视频讲座一起操作

加油！

---

## 命令行命令:

* 进入目录: `cd directoryPath`

* **WINDOWS:**

    * 列出目录所有文件: `dir`

* **OS X & LINUXes:**

    * 列出目录所有文件: `ls`

## 编译 & 运行 GO 程序:

* **编译一个 Go 程序:**

    * 在程序所在目录，输入:
        * `go build main.go`

* **运行一个 Go 程序:**

    * 在程序所在目录，输入:
        * `go run main.go`

## 你应该将源代码放在哪里？

* 你想放到哪个目录都可以

## 第一个程序

### 新建一个目录
* 新建一个目录:
  * `mkdir myDirectoryName`
* 进入目录:
  * `cd myDirectoryName`

### 添加源代码文件
* 新建源代码文件 `code main.go` 
  * 这条命令会在当前目录新建一个文件并使用 Visual Studio Code 打开该文件。
* 然后复制下面的代码到文件并保存文件

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

### 运行程序
* 最后，返回到命令行
  * 执行命令: `go run main.go`
* 如果你创建了其他文件并想同时执行它们，可以使用这个命令:
  * `go run .`

这就是全部内容了!

> 更多资料, 参考: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
