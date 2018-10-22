1. Print your name and your best friend's name using Println twice. [Check out this challenge here]().

2. Print your GOPATH using `go env` tool. [Check out this challenge here]().

3. **Say hello to yourself.**

    1. Build your program using `go build`
   
    2. **Send it to your friend**

       S/he should use be using the same operating system.

       For example, if you're using Windows, then hers/his should be Windows as well.

    3. **Send your program to a friend with a different operating system**

       So, you should compile your program for her operating system.

       **For OSX, type:**
       `GOOS=darwin GOARCH=386 go build`

       **For Windows:**
       `GOOS=windows GOARCH=386 go build`

       **For Linux:**
       `GOOS=linux GOARCH=arm GOARM=7 go build`

       **You can find the full list in here:**
       https://golang.org/doc/install/source#environment
   
4. **Call [Print](https://golang.org/pkg/fmt/#Print) instead of [Println](https://golang.org/pkg/fmt/#Println)** to see what happens.

5. **Call [Println](https://golang.org/pkg/fmt/#Println) or [Print](https://golang.org/pkg/fmt/#Print) with multiple values** by separating them using commas.

6. **Remove the double quotes from a string literal** and see what happens.

7. **Move the package and import statement** to the bottom of the file and see what happens.

8. **[Read Go online documentation](https://golang.org/pkg)**.

    1. Take a quick look at the packages and read what they do.

    2. Look at their source-code by clicking on their titles.

    3. You don't have to understand everything, just do it.This will warm you up for the upcoming lectures.

9. Also, **take a tour on**: https://tour.golang.org/

    1. Have a quick look. Check out the language features.
    
    2. We're going to talk all about them soon.