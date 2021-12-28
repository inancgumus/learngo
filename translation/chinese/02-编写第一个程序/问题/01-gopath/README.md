## 应该保存 Go 源代码到哪里？
* 我的电脑任何地方
* $GOPATH 下
* $GOPATH/src 下 *正确*

## $GOPATH 是什么?
* 它是 Go 运行时的文件
* 存储 Go 源代码文件和编译的包的目录
* 它是 gophers 需要关注的地址

## 需不需要设置 $GOPATH?
* 需要
* 不需要: 它保存在我的电脑里面
* 不需要: 它保存在用户目录下 *正确*

## 如何打印 $GOPATH 的值？
* 使用 `ls` 命令
* 使用 `go env GOPATH` 命令 *正确*
* 使用 `go environment` 命令