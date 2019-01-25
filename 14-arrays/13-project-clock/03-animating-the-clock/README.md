# GOAL 3: Animate the Clock

## Notes

* Note main.go file contains the solution of the previous step.
* "solution" folder contains the solution for this step.

## Challenge Steps

1. Create an infinite loop to update the clock

2. Update the clock every second

   [time.Sleep(time.Second)](https://golang.org/pkg/time/#Sleep) will stop the world for 1 second

3. Clear the screen before the infinite loop

   1. Get my library for clearing the screen:

        `go get -u github.com/inancgumus/screen`

   2. Then, import it and call it in your code like this:

        `screen.Clear()`

   3. If you're using Go Playground instead, do this:

        `fmt.Println("\f")`

 4. Move the cursor to the top-left corner of the screen, before each step
    of the infinite loop

    * Call this in your code like this:

        `screen.MoveTopLeft()`

    * If you're using Go Playground instead, do this again:

        `fmt.Println("\f")`

---

## SIDE NOTE FOR THE CURIOUS

If you're curious about how my screen clearing package works, read on.

**On bash**, it uses special commands, if you open the code, you can see that.

* `\033` is a special control code:
* `[2J` clears the screen and the cursor
* `[H`  moves the cursor to 0, 0 screen position
* [See for more information](https://bluesock.org/~willkg/dev/ansi.html).

**On Windows**, I'm directly calling the Windows API functions. This is way advanced at this stage of the course, however, I'll probably explain it afterward.

So, my package automatically adjusts itself depending on where it is compiled. On Windows, it uses the special Windows API calls; On other operating systems, it uses the bash special commands that I've explained above.
