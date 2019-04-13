# OS X INSTALLATION CHEATSHEET

## NOTE

If you have [homebrew](https://brew.sh) installed, you can easily install Go like so:

```
# if you don't have git install it like so:
brew install git

# then install go
brew install go

# add GOBIN path to your PATH in ~/.bash_profile
export PATH=${HOME}/go/bin:$PATH
```

## 1- Install Visual Studio Code Editor

1. Install it but don't open it yet.
2. Go to: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Select OS X (Mac) and start downloading
4. Uncompress the downloaded file and move it to your `~/Applications` directory.

## 2- Install Git

1. Grab and run the installer. Go to: [https://git-scm.com/downloads](https://git-scm.com/downloads)
2. Select VSCode as the default editor

## 3- Install Go

1. Go to [https://golang.org/dl](https://golang.org/dl)
2. Select OS X (Mac)
3. Start the installer

## 4- Configure VS Code

1. Open VS Code; from the extensions tab at the left, search for "go" and install it
2. Close VS Code completely and open it up again

3. Go to View menu; select **Command Palette**
    1. Or just press `cmd+shift+p`
    2. Type: `go install`
    3. Select _"Go: Install/Update Tools"_
    4. Check all the checkboxes

4. After it's done, open the **Command Palette** again
    1. Type: `shell`
    2. Select: _"Install 'code' command in PATH"_
        1. **NOTE:** You don't have to do this if you're on Windows.

## That's all! Enjoy! 🤩

<div style="page-break-after: always;"></div>

> For more tutorials: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> Learn Go Programming Course
> 
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
