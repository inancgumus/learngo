# Шпаргалка по установке на Linux

Если вы хотите использовать snap, Go устанавливается так:

    sudo snap install go --classic
    
Иначе выполните шаги ниже:

## 1. Обновите локальные пакеты

  ```bash
  sudo apt-get update
  ```

## 2. Установите git

  ```bash
  sudo apt-get install git
  ```

## 3. Установите Go

Есть два способа:

1- Через браузер: выберите Linux, и загрузка начнётся.

  ```bash
  firefox https://golang.org/dl
  ```

2- Через snap: если вы выбрали этот способ, переходите сразу к шагу 5.

  ```bash
  sudo snap install go --classic
  ```

## 4. Скопируйте Go в нужный каталог

1. Узнайте имя скачанного файла

2. Используйте это имя, чтобы распаковать файл

    ```bash
    gofile="DELETE_THIS_AND_TYPE_THE_NAME_OF_THE_DOWNLOADED_FILE_HERE (without its extension)"

    tar -C /usr/local -xzf ~/Downloads/$gofile
    ```

## 5. Добавьте каталог исполняемых файлов Go в PATH

1. Добавьте каталог `go/bin` в `$PATH`, чтобы можно было запускать основные команды Go.

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    ```

2. Добавьте каталог "$HOME/go/bin" в $PATH

    ```bash
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.profile
    ```

## Установка инструментов Go:

* Это очень удобные инструменты, которые облегчают разработку (например, goimports)

* `go get` нельзя использовать без программы контроля версий, например Git, которую мы уже установили выше.

* Эта команда создаст каталог `~/go` и скачает туда инструменты Go.

    * В этот каталог также стоит класть свой код.
    (Если вы не собираетесь использовать Go Modules)

    ```bash
    go get -v -u golang.org/x/tools/...
    ```

## Установка VSCode (необязательно)

Замечание: при желании вы можете использовать другой редактор кода. Однако в курсе используется Visual Studio Code (VSCode).

1. Откройте приложение "Ubuntu Software"

2. Найдите VSCode и нажмите "Install"


## НЕОБЯЗАТЕЛЬНЫЙ ШАГ:

1. Создайте файл hello.go в новом каталоге, но в любом месте вне `$GOPATH`

    ```bash
    cat <<EOF > hello.go
    package main

    import "fmt"

    func main() {
        fmt.Println("hello gopher!")
    }
    EOF
    ```

2. Запустите программу

    ```bash
    go run hello.go
    It should print: hello gopher!
    ```

<div style="page-break-after: always;"></div>

> Больше руководств: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Нажмите здесь, чтобы прочитать лицензию.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
