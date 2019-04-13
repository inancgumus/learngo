# WINDOWS INSTALLATION CHEATSHEET

## NOTE

If you have [chocolatey.org](https://chocolatey.org/) package manager, you can easily install Go like so:

```
choco install golang
```

## 1- Install Visual Studio Code Editor

1. Install it but don't open it yet.
2. Go to: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Select Windows and start downloading
4. Run the installer

## 2- Install Git

1. Grab and run the installer. Go to: [https://gitforwindows.org](https://gitforwindows.org)
2. Select VSCode as the default editor
3. Enable all the checkboxes
4. Select: _"Use Git from the Windows Command Prompt"_
5. Encodings: Select: _"Checkout as is..." option._

## 3- Install Go

1. Go to [https://golang.org/dl](https://golang.org/dl)
2. Select Windows and download
3. Start the installer

## 4- Configure VS Code

1. Open VS Code; from the extensions tab at the left, search for "go" and install it
2. Close VS Code completely and open it up again

3. Go to View menu; select **Command Palette**
    1. Or just press `ctrl+shift+p`
    2. Type: `go install`
    3. Select _"Go: Install/Update Tools"_
    4. Check all the checkboxes

## 5- Using Git-Bash

* In this course I'll be using bash commands. Bash is just a command-line interface used in OS X and Linux. It's one of the most popular command-line interfaces. So, if you want to use it too, instead of using the Windows default command-line, you can use git bash that you've installed. With git bash, you can type a command in command-line as you're on OS X or Linux.

* If you don't want to use git bash, it's ok too. It depends on you. But note that, I'll be using bash commands mostly. Because it allows more advanced commands as well.

* You can also prefer to use the more powerful alternative to git bash: [Linux Subsystem for Windows](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

* **So, to use git bash, follow these steps:**
    1. Just search for git bash from the start bar
    2. Or, if there's one, click on the icon on your desktop

    3. Also, setup VS Code to use git-bash by default:
        1. Open VS Code
        2. Go to Command Palette
            1. Type: `terminal`
            2. Select: _"Terminal: Select Default Shell"_
            3. And, Select: _"Git Bash"_

    4. **NOTE:** Normally, you can find your files under `c:\`, however, when you're using git bash, you'll find them under `/c/` directory. It's actually the very same directory, it's just a shortcut.

## That's all! Enjoy! 🤩

<div style="page-break-after: always;"></div>

> For more tutorials: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
