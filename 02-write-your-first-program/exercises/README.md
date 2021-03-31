1. **Print your name and your best friend's name** using Println twice. [Check out this exercise here](https://github.com/inancgumus/learngo/tree/master/02-write-your-first-program/exercises/01-print-names).

2. **Say hello to yourself.**

    1. Build your program using `go build`

    2. **Send it to your friend**

       S/he should use be using the same operating system.

       For example, if you're using Windows, then hers/his should be Windows as well.

    3. **Send your program to a friend with a different operating system**

       So, you should compile your program for her operating system.

       **Create an OSX executable:**
       `GOOS=darwin GOARCH=386 go build`

       **Create a Windows executable:**
       `GOOS=windows GOARCH=386 go build`

       **Create a Linux executable:**
       `GOOS=linux GOARCH=arm GOARM=7 go build`

       **You can find the full list in here:**
       https://golang.org/doc/install/source#environment

       **NOTE:** If you're using the command prompt or the PowerShell, you may need to use it like this:
       `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

3. **Call [Print](https://golang.org/pkg/fmt/#Print) instead of [Println](https://golang.org/pkg/fmt/#Println)** to see what happens.

4. **Call [Println](https://golang.org/pkg/fmt/#Println) or [Print](https://golang.org/pkg/fmt/#Print) with multiple values** by separating them using commas.

5. **Remove the double quotes from a string literal** and see what happens.

6. **Move the package and import statement** to the bottom of the file and see what happens.

7. **[Read Go online documentation](https://golang.org/pkg)**.

    1. Take a quick look at the packages and read what they do.

    2. Look at their source-code by clicking on their titles.

    3. You don't have to understand everything, just do it. This will warm you up for the upcoming lectures.

8. Also, **take a tour on**: https://tour.golang.org/

    1. Have a quick look. Check out the language features.

    2. We're going to talk all about them soon.

9. [Follow me on twitter to learn more](https://twitter.com/inancgumus).