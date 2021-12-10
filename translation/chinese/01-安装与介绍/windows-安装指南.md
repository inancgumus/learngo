# WINDOWS 安装指南

## 笔记

如果已安装 [chocolatey.org](https://chocolatey.org/) 包管理工具, 可以通过如下命令安装:

```
choco install golang
```

## 1- 安装 Visual Studio Code Editor

1. 安装但是不要直接打开.
2. 打开 : [https://code.visualstudio.com](https://code.visualstudio.com)
3. 选择 Windows并进行安装
4. 执行installer

## 2- 安装 Git

1. 通过官网下载和安装. 官网地址: [https://gitforwindows.org](https://gitforwindows.org)
2. 选择 VSCode为默认编辑器
3. 勾选所以复选框,启用所有可选功能
4. 选择: _"Use Git from the Windows Command Prompt"_
5. 编码: 选择: _"Checkout as is..." option._

## 3- 安装 Go

1. 登陆Go官网 [https://golang.org/dl](https://golang.org/dl)
2. 选择 Windows版本下载
3. 执行Go安装器

## 4- 配置 VS Code

1. 运行 VS Code;在插件市场搜索"go"并安装
2. 关闭 VS Code并重新打开

3. 前往视图菜单; 选择 **Command Palette**
    1. 执行 `ctrl+shift+p`
    2. 键入: `go install`
    3. 选择 _"Go: Install/Update Tools"_
    4. 勾选所有复选框

## 5- 使用 Git-Bash

* 在本课程我们将使用bash作为命令行. Bash是在OS X 和 Linux系统上非常流行的命令行工具. 
  如果你也希望使用bash 而不是Windows提供的默认命令行, 可以使用已安装的bash命令行, 好似使用OS X 和 Linux系统.

* 如果您不打算使用git bash也是可以的. 但是, 本课程将使用git bash. git bash可以提供本课程需要的许多高级特性.

* 您也可以选择更加强大的工具来替代 git bash, 参考: [Linux Subsystem for Windows](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

* **因此, 如果使用 git bash, 请遵守如下步骤:**
    1. 在开始栏搜索git bash
    
    2. 或者, 如果已经下载可以直接启用

    3. 设置 VS Code 默认使用 git-bash:
        1. 启动 VS Code
        2. 转到命令面板
            1. 输入: `terminal`
            2. 选择: _"Terminal: Select Default Shell"_
            3. 选择: _"Git Bash"_

    4. **提示:** 一般, 可以在 `c:\`文件夹找到自定义文件. 但是, 档我们使用git bash, 自定义问价会在文件夹 `/c/`. 事实上，这两个文件目录是同一个目录, 只是一个软链接.

## 以上步骤完成，就可以使用了! 🤩

<div style="page-break-after: always;"></div>

> 更多资料, 参考: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
