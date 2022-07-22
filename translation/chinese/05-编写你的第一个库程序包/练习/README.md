[Check out the exercise and its solution here.](https://github.com/inancgumus/learngo/tree/master/05-write-your-first-library-package/exercise)

---

# 练习

1. 新建一个库
2. 在库中，创建一个返回 Go 版本的函数。
3. 创建一个可执行程序并导入你的库
4. 调用你创建的返回 Go 版本的函数
5. 执行你的程序

## 提示

**像这样创建你的包函数:**

```go
func Version() string {
    return runtime.Version()
}
```

## 期望输出

他应该打印出你系统安装的 Go 的版本

## 警告

你应该在你自己的文件夹下创建这个包，而不是在 github.com/inancgumus/learngo 目录下。同时，注意 VS Code 可能会自动导入我位于 github.com/inancgumus/learngo 的库而非你自己的库。

所以，如果你想要 VS Code 自动在保存时导入你自己的包，只需要将 github.com/inancgumus/learngo 移动到非 GOPATH 的某处，比如，你的桌面（记得在之后移动回来）

See [this question](https://www.udemy.com/learn-go-the-complete-bootcamp-course-golang/learn/v4/questions/5518190) in Q&A for more information.