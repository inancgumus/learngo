# ШПАРГАЛКА ПО УСТАНОВКЕ НА OS X

## ЗАМЕЧАНИЕ

Если у вас установлен [homebrew](https://brew.sh), Go устанавливается так:

```
# если у вас нет git, установите его так:
brew install git

# затем установите go
brew install go

# добавьте путь GOBIN в PATH в ~/.bash_profile
export PATH=${HOME}/go/bin:$PATH
```

## 1- Установите редактор Visual Studio Code

1. Установите его, но пока не открывайте.
2. Перейдите на: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Выберите OS X (Mac) и начните загрузку
4. Распакуйте скачанный файл и переместите его в каталог `~/Applications`.

## 2- Установите Git

1. Скачайте и запустите установщик. Перейдите на: [https://git-scm.com/downloads](https://git-scm.com/downloads)
2. Выберите VSCode в качестве редактора по умолчанию

## 3- Установите Go

1. Перейдите на [https://golang.org/dl](https://golang.org/dl)
2. Выберите OS X (Mac)
3. Запустите установщик

## 4- Настройте VS Code

1. Откройте VS Code; на вкладке расширений слева найдите "go" и установите его
2. Полностью закройте VS Code и откройте снова

3. Откройте меню View; выберите **Command Palette**
    1. Или просто нажмите `cmd+shift+p`
    2. Введите: `go install`
    3. Выберите _"Go: Install/Update Tools"_
    4. Отметьте все флажки

4. Когда всё готово, снова откройте **Command Palette**
    1. Введите: `shell`
    2. Выберите: _"Install 'code' command in PATH"_
        1. **ЗАМЕЧАНИЕ:** На Windows этого делать не нужно.

## Вот и всё! Приятного обучения! 🤩

<div style="page-break-after: always;"></div>

> Больше руководств: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> Learn Go Programming Course
> 
> [Нажмите здесь, чтобы прочитать лицензию.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
