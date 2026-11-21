# ШПАРГАЛКА ПО УСТАНОВКЕ НА WINDOWS

## ЗАМЕЧАНИЕ

Если у вас есть пакетный менеджер [chocolatey.org](https://chocolatey.org/), Go устанавливается так:

```
choco install golang
```

## 1- Установите редактор Visual Studio Code

1. Установите его, но пока не открывайте.
2. Перейдите на: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Выберите Windows и начните загрузку
4. Запустите установщик

## 2- Установите Git

1. Скачайте и запустите установщик. Перейдите на: [https://gitforwindows.org](https://gitforwindows.org)
2. Выберите VSCode в качестве редактора по умолчанию
3. Включите все флажки
4. Выберите: _"Use Git from the Windows Command Prompt"_
5. Кодировки: выберите вариант _"Checkout as is..."_.

## 3- Установите Go

1. Перейдите на [https://golang.org/dl](https://golang.org/dl)
2. Выберите Windows и скачайте
3. Запустите установщик

## 4- Настройте VS Code

1. Откройте VS Code; на вкладке расширений слева найдите "go" и установите его
2. Полностью закройте VS Code и откройте снова

3. Откройте меню View; выберите **Command Palette**
    1. Или просто нажмите `ctrl+shift+p`
    2. Введите: `go install`
    3. Выберите _"Go: Install/Update Tools"_
    4. Отметьте все флажки

## 5- Использование Git-Bash

* В этом курсе я буду использовать команды bash. Bash — это просто интерфейс командной строки, который применяется в OS X и Linux. Это один из самых популярных интерфейсов командной строки. Поэтому, если вы тоже хотите его использовать вместо стандартной командной строки Windows, вы можете взять git bash, который вы уже установили. В git bash можно вводить команды так же, как в OS X или Linux.

* Если вы не хотите использовать git bash, это тоже нормально. Решать вам. Но учтите, что я в основном буду использовать команды bash. Потому что они позволяют выполнять и более сложные команды.

* Вы также можете предпочесть более мощную альтернативу git bash: [Linux Subsystem for Windows](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

* **Итак, чтобы использовать git bash, выполните следующие шаги:**
    1. Просто найдите git bash в строке поиска меню Пуск
    2. Или, если он есть, нажмите на значок на рабочем столе

    3. Кроме того, настройте VS Code так, чтобы он по умолчанию использовал git-bash:
        1. Откройте VS Code
        2. Откройте Command Palette
            1. Введите: `terminal`
            2. Выберите: _"Terminal: Select Default Shell"_
            3. И выберите: _"Git Bash"_

    4. **ЗАМЕЧАНИЕ:** Обычно ваши файлы находятся в `c:\`, однако при использовании git bash вы найдёте их в каталоге `/c/`. Это фактически один и тот же каталог, просто ярлык.

## Вот и всё! Приятного обучения! 🤩

<div style="page-break-after: always;"></div>

> Больше руководств: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Learn Go Programming Course
>
> [Нажмите здесь, чтобы прочитать лицензию.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
