1. **Run your own program? Say hello to yourself.**

    1. Build your program using `go build`
   
    2. And, send it to your friend
       (s/he should use be using the same operating system)
       (if you're using windows, then hers/his should be
       windows too)

    3. And then send your program to a friend with a different
       operating system.

       (So, you should compile your program for her operating system).

       **For OSX, type:**
       GOOS=darwin GOARCH=386 go build

       **For Windows:**
       GOOS=windows GOARCH=386 go build

       **For Linux:**
       GOOS=linux GOARCH=arm GOARM=7 go build

       **You can find the full list in here:**
       https://golang.org/doc/install/source#environment
   
2. **Call Print instead of Println** to see what happens.

3. **Call Println or Print with multiple values** by separating them using commas.

4. **Remove double quotes from string literals** and see what happens.

5. **Move the package and import statement** to the bottom of the file and see what happens.

6. **Read Go online documentation**. Take a quick look at the packages and read what they do. Look at their source-code by clicking on their titles.

You don't have to understand anything, just do it. This will warm you up for the upcoming lectures. https://golang.org/pkg

7. Also, **take a tour**: https://tour.golang.org/ See the language features. We're going to talk all about them soon.