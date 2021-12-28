1. **打印你和你最好朋友的名字** 使用 Println 两次。 [练习点击这里](https://github.com/inancgumus/learngo/tree/master/02-write-your-first-program/exercises/01-print-names).

2. **自娱自乐**

    1. 使用 `go build` 编译程序

    2. **发送程序给你的朋友**

       他/她应该和你使用相同的操作系统

       例如，如果你使用的是 Windows ，他/她应该也使用 Windows 

    3. **发送程序给你的朋友如果他使用不同的操作系统**

       这样，你应该为他/她的操作系统编译程序

       **编译 OSX 可执行程序:**
       `GOOS=darwin GOARCH=386 go build`

       **编译 Windows 可执行程序:**
       `GOOS=windows GOARCH=386 go build`

       **编译 Linux 可执行程序:**
       `GOOS=linux GOARCH=arm GOARM=7 go build`

       **通过这里可以获取全部列表:**
       https://golang.org/doc/install/source#environment

       **注意:** 如果你使用命令提示符或者 PowerShell , 你可能需要使用下面的命令:
       `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

3. **调用 [Print](https://golang.org/pkg/fmt/#Print) 而不是 [Println](https://golang.org/pkg/fmt/#Println)** 看看会发生什么

4. **调用 [Println](https://golang.org/pkg/fmt/#Println) 或者 [Print](https://golang.org/pkg/fmt/#Print) 输入多个数据** 通过逗号分隔

5. **移除字符串两边的双引号** 看看会发生什么

6. **将 package 和 import 声明** 移到文件末尾看看会发生什么

7. **[阅读 Go 在线文档](https://golang.org/pkg)**.

    1. 快速浏览 packages ，看看它们的用途

    2. 通过点击标题查看其源代码

    3. 你不必弄懂全部，只需要亲自写吧，这仅仅是为了后面的课程热身

8. 对了, **打开 tour**: https://tour.golang.org/

    1. 快速浏览一下，查看一下语言特点

    2. 我们很快会讨论它们

9. [关注我的 twitter 学习更多](https://twitter.com/inancgumus).