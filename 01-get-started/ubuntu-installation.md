# Linux Installation Cheatsheet

## 1. Update the local packages

  ```bash
  sudo apt-get update
  ```

## 2. Install git

  ```bash
  sudo apt-get install git
  ```

## 3. Install Go

Select Linux and the download will begin.

  ```bash
  firefox https://golang.org/dl
  ```

## 4. Copy Go into the proper directory

1. Find out the name of the downloaded file
2. Use that filename to uncompress it

    ```bash
    gofile="DELETE_THIS_AND_TYPE_THE_NAME_OF_THE_DOWNLOADED_FILE_HERE (without its extension)"

    tar -C /usr/local -xzf ~/Downloads/$gofile.tar.gz
    ```

3. Add `go/bin` directory to `$PATH` to be able to run the fundamental Go commands.

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    ```

4. Add "$HOME/go/bin" directory to $PATH

    ```bash
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.profile
    ```

## Install Go Tools:

* These are very handy tools to ease the development (like goimports)

* `go get` cannot be used without installing a code versioning program like Git which we already have got it above.

* This will create `~/go` directory and will download go tools into there.

    * This directory is also a place where you should put your code into.
    (If you're not going to use Go Modules)

    ```bash
    go get -v -u golang.org/x/tools/...
    ```

## Install VSCode

1. Open "Ubuntu Software" application
2. Search for VSCode then click "Install"
3. Install [Vim-Go](https://github.com/fatih/vim-go#install)

## OPTIONAL STEP:

1. Create a hello.go inside `$GOPATH/src`

    ```bash
    cat <<EOF > $GOPATH/src/hello.go
    package main

    import "fmt"

    func main() {
        fmt.Println("hello gopher!")
    }
    EOF
    ```

2. Go to Go source directory and run our sample program

    ```bash
    cd $GOPATH/src
    go run hello.go
    It should print: hello gopher!
    ```

<div style="page-break-after: always;"></div>

> For more tutorials: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Click here to read the license.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
