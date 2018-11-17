[Check out the exercise and its solution here.](https://github.com/inancgumus/learngo/tree/master/05-write-your-first-library-package/exercise)

---

# EXERCISE
1. Create a new library
2. In it, create a function that returns the Go version
3. Create a command and import your library
4. Call your function that returns Go version
5. Run your program

## HINTS
**Create your package function like this:**

```go
func Version() string {
    return runtime.Version()
}
```

## EXPECTED OUTPUT
It should print the current Go version on your system.

## WARNING

You should create this package under your own folder, not in github.com/inancgumus/learngo folder. Also, please note that VS Code may automatically import my library which is in github.com/inancgumus/learngo instead of your own library.

So, if you want VS Code automatically import your own package when you save, just move github.com/inancgumus/learngo out of GOPATH to somewhere else, for example, to your Desktop (of course move it back afterward).

See [this question](https://www.udemy.com/learn-go-the-complete-bootcamp-course-golang/learn/v4/questions/5518190) in Q&A for more information.
