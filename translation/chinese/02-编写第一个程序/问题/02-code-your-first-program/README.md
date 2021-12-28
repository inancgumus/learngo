## 下面哪一个关键字用来定义一个包？
```go
package main

func main() {
}
```
1. func
2. package *正确*
3. fmt.Println
4. import

> **1:** func 关键字用来新定义一个函数
>
>
> **2:** 对，package 关键字用来在 Go 文件中定义包
>
>
> **3:** 这不是关键字，它是 fmt 包中的一个函数
>
>
> **4:** import 关键字用来导入一个包
>
>


## 在下面代码中使用 package main 的目的是什么？
```go
package main

func main() {
}
```
* 为了创建一个库包
* 为了正确退出程序
* 为了创建一个可执行的 Go 程序 *正确*


## 在下面代码中使用 func main 的目的是什么？
```go
package main

func main() {
}
```
1. 为了创建一个叫 main 的包
2. 为了让 Go 开始执行程序 *正确*
3. 为了向控制台打印一条消息

> **1:** main function 不创建一个包
>
>
> **2:** 对，Go 自动地从 main 函数开始执行
>
>
> **3:** 它不会打印任何东西(目前为止)
>
>


## 在下面代码中使用 import "fmt" 的目的是什么？
```go
package main
import "fmt"

func main() {
    fmt.Println("Hi!")
}
```
1. 它打印 "fmt" 到控制台
2. 它创建一个叫 "fmt" 的包
3. 它导入了 `fmt` 包; 这样你就可以使用其功能 *正确*

> **1:** `fmt.Println` 打印一条消息而不是 `import "fmt"`
>
>
> **2:** `package` 关键字创建新包，而不是 `import`
>
>
> **3:** 对，举个例子，导入 fmt 包后，你就可以使用 Println 打印一条消息到控制台
>
>


## 下面哪一个用来定义函数？
* func *正确*
* package
* Println
* import


## 函数是什么?
1. 它就像迷你程序，它是可复用可执行的代码块 *正确*
2. 它允许 Go 执行一个程序
3. 它允许 Go 导入一个叫 function 的包
4. 它打印一条消息到控制台

> **2:** Go 寻找 main 包的 func main 去执行，一个函数自己不能执行
>
>
> **3:** 这是 `import` 的作用
>
>
> **4:** 举个例子: 这是 `fmt.Println` 的功能
>
>


## 是否需要亲自调用 main 函数
1. 是的，只有这样才能运行自己的程序
2. 不必，Go 会自动地调用 main 函数 *正确*

> **1:** 不必, 不需要调用 main 函数， Go 自动执行它
>
>


## 是否需要调用一个函数去执行它?
_(main 函数除外)_
1. 需要，只有这样 Go 才能执行一个函数 *正确*
2. 需要，只有这样 Go 才执行程序
3. 不需要，函数会自动执行

> **1:** 对的，需要自己调用函数，Go 不会自动执行函数，只有 main 函数自动调用(还有其他一些函数，不过目前还未学到)
>
>

> **2:** 这是 `func main` 的唯一工作， 也只有 `func main` 才能执行程序
>
>

> **3:** Go 不会自动调用函数，除了 main 函数(还有其他一些函数，不过目前还未学到)，所以，除了 main 函数，你需要手动调用函数
>


## 下面的程序输出是什么？
```go
package main

func main() {
}
```
1. 它打印一条消息到控制台
2. 它是正确的程序，但是什么都不输出 *正确*
3. 它不是正确的程序

> **1:** 它不是输出任何消息，可以使用 fmt.Println 函数来输出内容
>
>

> **2:** 对，它是正确的程序，然后因为它不包含 fmt.Println 所以它什么也不输出
>
>

> **3:** 它是正确的程序，它使用 package 关键字，并且有 main 函数。因此，这是一个有效且可执行的 Go 程序
>
>


## 这个程序的输出是什么？
```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}
```
* Hi! I want to be a Gopher!
* 它什么都不打印
* 它不是正确的程序 *正确*

> **1:** 没有用双引号括起来，Println 不能获取值，应该是: fmt.Println("Hi! I want to be a Gopher")
>
>

> **3:** 它没有导入 "fmt" 包， 还有 #1.
>
>


## 这个程序的输出是什么？
```go
package main
import "fmt"

func main() {
    fmt.Println("Hi there!")
}
```
* Hi there! *正确*
* fmt
* 它不是正确的程序; 它导入的包里面没有函数叫 `Println`

> **2:** import "fmt" 导入 fmt 包，就可以使用其功能
>
>
> **3:** 实际上，这个程序是正确的
>
>