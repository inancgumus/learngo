# Linux Installation Cheatsheet

If you want to use snap, you can easily install Go like so:

    sudo snap install go --classic
    
Otherwise, please follow the steps below:

## 1. Update the local packages

  ```bash
  sudo apt-get update
  ```

## 2. Install git

  ```bash
  sudo apt-get install git
  ```

## 3. Install Go

There are two ways:

1- From Web: Select Linux and the download will begin.

  ```bash
  firefox https://golang.org/dl
  ```

2- By using snap: If you use this option, skip to the 5th step.

  ```bash
  sudo snap install go --classic
  ```

## 4. Copy Go into the proper directory

1. Find out the name of the downloaded file

2. Use that filename to uncompress it

    ```bash
    gofile="DELETE_THIS_AND_TYPE_THE_NAME_OF_THE_DOWNLOADED_FILE_HERE (without its extension)"

    tar -C /usr/local -xzf ~/Downloads/$gofile
    ```

## 5. Add Go executables directory to your PATH

1. Add `go/bin` directory to `$PATH` to be able to run the fundamental Go commands.

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    ```

2. Add "$HOME/go/bin" directory to $PATH

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

## Install VSCode (Optional)

Note: You may use another coding editor if you like. However, the course uses Visual Studio Code (VSCode).

1. Open "Ubuntu Software" application

2. Search for VSCode then click "Install"


## OPTIONAL STEP:

1. Create a hello.go file in a new directory but anywhere outside of `$GOPATH`

    ```bash
    cat <<EOF > hello.go
    package main

    import "fmt"

    func main() {
        fmt.Println("hello gopher!")
    }
    EOF
    ```

2. Run the program

    ```bash
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
