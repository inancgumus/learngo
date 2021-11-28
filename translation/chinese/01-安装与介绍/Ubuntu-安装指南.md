# Linux 安装指南

如果,只需要安装Go，请执行如下指令:

    sudo snap install go --classic

否则, 请按如下安装步骤安装:

## 1. 更新本地依赖包

  ```bash
  sudo apt-get update
  ```

## 2. 安装 git

  ```bash
  sudo apt-get install git
  ```

## 3. 安装 Go

安装Go环境有如下两种方法:

1- 网站: 选择 Linux 版本并安装.

  ```bash
  firefox https://golang.org/dl
  ```

2- 使用snap: 如果选择snap安装, 可以直接跳到步骤5.

  ```bash
  sudo snap install go --classic
  ```

## 4. 将Go拷贝到适当的工作目录

1. 找到下载好的文件

2. 解压文件

    ```bash
    gofile="DELETE_THIS_AND_TYPE_THE_NAME_OF_THE_DOWNLOADED_FILE_HERE (替换为下载文件，省去后缀)"
    tar -C /usr/local -xzf ~/Downloads/$gofile
    ```

## 5. 添加Go的可执行路径

1. 添加`go/bin`目录到 `$PATH`, 以便执行Go的基础命令.

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    ```

2. 添加 "$HOME/go/bin" 目录到 $PATH

    ```bash
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.profile
    ```

## 安装Go工具:

* 有很多开发工具，可以让简化Go的开发 (如 goimports)

* 如果系统中没有安装代码版本工具(如 Git), 将无法使用`go get`.

* 上述命令会创建`~/go` 目录并下载到相应目录中.

    * 此目录也将是我们存放代码的目录.(如果不实用Go Modules)

    ```bash
    go get -v -u golang.org/x/tools/...
    ```

## 安装 VSCode (可选)

提示: 本课程将使用Visual Studio Code (VSCode). 但是, 可以使用任何代码编辑工具.

1. 打开 "Ubuntu Software" 应用

2. 搜索并安装VSCode


## 可选步骤:

1. 在`$GOPATH`之外的任何路径创建文件 hello.go

    ```bash
    cat <<EOF > hello.go
    package main
    import "fmt"
    func main() {
        fmt.Println("hello gopher!")
    }
    EOF
    ```
2. 执行程序
    ```bash
    go run hello.go
    It should print: hello gopher!
    ```
<div style="page-break-after: always;"></div>
> 更多内容: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)