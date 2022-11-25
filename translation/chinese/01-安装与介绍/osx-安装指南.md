# OS X 安装

## 提示

如果已经安装好[homebrew](https://brew.sh), 可以按照如下步骤安装Go

```
# 安装git可以参照如下命令:
brew install git

# go 安装命令
brew install go

# 添加 GOBIN 路径 ~/.bash_profile
export PATH=${HOME}/go/bin:$PATH
```

## 1- 安装 Visual Studio Code Editor

1. 安装地址: [https://code.visualstudio.com](https://code.visualstudio.com)
2. 选择 OS X (Mac) 下载
3. 解压 添加到应用程序`~/Applications`路径.

## 2- 安装 Git

1. 下载Git. 官网: [https://git-scm.com/downloads](https://git-scm.com/downloads)
2. 选择VSCode为默认编辑器

## 3- 安装 Go

1. 官网 [https://golang.org/dl](https://golang.org/dl)
2. 选择 OS X (Mac)
3. 下载

## 4- 配置 VS Code

1. 打开 VS Code; 点击扩展插件功能, 搜索 "go" 下载
2. 下载完成后重启VS Code

3. 选择 View 菜单; 选择 **Command Palette**
    1. 输入 `cmd+shift+p`
    2. 输入: `go install`
    3. 选择 _"Go: Install/Update Tools"_
    4. 选择所有 checkboxes

4. 重启 **Command Palette** 
    1. 输入: `shell`
    2. 选择: _"Install 'code' command in PATH"_
        1. **提示:** Windows操作系统不需要上述操作.

## 工作完成! Enjoy! 🤩

<div style="page-break-after: always;"></div>

> 更多内容: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> Learn Go Programming Course
> 
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
