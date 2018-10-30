# CHEATSHEET: INSTALLATION

👉 NOTE: For the missing steps you should continue clicking the Next buttons :)

# 1 Install Visual Studio Code Editor

1. Install it but don't open it yet.
2. Go to: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Select your operating system (OS) and start downloading
    1. **Windows**: Run the installer.
    2. **OS X**: Uncompress the downloaded file and move it to your `~/Applications` directory.

# 2 Install Git

1. Go to: https://git-scm.com/downloads
2. Select your OS and start downloading
3. Run the installer
4. Select VSCode as the default editor
    1. **Windows**:
        1. Enable all the checkboxes
        2. Select: _"Use Git from the Windows Command Prompt"_
        3. Encodings: Select: _"Checkout as is..." option._

# 3 Install Go

1. Go to [https://golang.org/dl](https://golang.org/dl)
2. Select your OS and download
3. Start the installer
    1. **Windows:**
        1. Installs Go to `C:\Go`
        2. Your `$GOPATH` will be `C:\Go\src`

    2. **OS X:**
        1. Installs Go to `~/go`
        2. Your `$GOPATH` will be `~/go/src`

<div style="page-break-after: always;"></div>

# 4 Configure VS Code

1. Open VS Code; from the extensions tab at the left, search for "go" and install it
2. Close VS Code completely and open it up again

3. Go to View menu; select **Command Palette**
    1. Or just press `cmd+shift+p` or `ctrl+shift+p`
    2. Type: `go install`
    3. Select _"Go: Install/Update Tools"_
    4. Check all the checkboxes

4. After it's done, open the **Command Palette** again
    1. Type: `shell`
    2. Select: _"Install 'code' command in PATH"_
        1. **NOTE:** You don't have to do this if you're on Windows.

5. **Additional Step for Windows Users: Git-Bash**
    * In this course I'll be using bash commands. Bash is just a command-line interface used in OS X and Linux. It's one of the most popular command-line interfaces. So, if you want to use it too, instead of using the Windows default command-line, you can use git bash that you've installed. With git bash, you can type a command in command-line as you're on OS X or Linux.

    * If you don't want to use git bash, it's ok too. It depends on you. But note that, I'll be using bash commands mostly. Because, it allows more advanced commands as well.

    * So, to use git bash, follow these steps:
        1. Just search for git bash from the start bar
        2. Or, if there's one, click on the icon on your desktop

        3. Also setup VS Code to use git-bash by default:
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