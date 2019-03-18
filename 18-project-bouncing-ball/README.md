# Bouncing Ball Challenge Tips

Use the following tips only when you get stuck. This document isn't in a particular order, please do not follow it like so.

## CALCULATING THE VELOCITY

You can use velocity to change the ball's speed and position. In my example, the speed is constant, so I always use unit value: 1.

* On each loop step: Add velocities to ball's position. This will make the ball move.

* **Velocity means: Speed and Direction**

  * X velocity =  1 -> _ball moves right_
  * X velocity = -1 -> _ball moves left_
  * Y velocity =  1 -> _ball moves down_
  * Y velocity = -1 -> _ball moves up_

* **For more information on graphics and velocity:**

  * [Youtube: Crash Course: 2D Graphics](https://www.youtube.com/watch?v=7Jr0SFMQ4Rs&t=529)

  * [Youtube: Crash Course: Velocity](https://www.youtube.com/watch?v=ZM8ECpBuQYE)

## CREATING THE BOARD

I use `[][]bool` for the board but you can use anything you like. For example, you can directly use `[][]rune` or `[]rune`. Experiment with them and decide which one is the best for you.

## CLEARING THE SCREEN

* Before the loop, clear the screen once by using my [screen package](https://github.com/inancgumus/screen), click on the link. You can find its [documentation here](https://godoc.org/github.com/inancgumus/screen).

* After each loop step, move the cursor to the top-left position by using the screen package. So that you can draw the animation frame all over again in the same position.

* You can find more information about the screen package and screen clearing in the [Retro Clock project section lectures](https://github.com/inancgumus/learngo/tree/master/15-project-retro-led-clock).

## DRAWING THE BOARD

Instead of drawing the board and the ball to the screen everytime, you will fill a buffer, and when you complete, you can draw the board and the ball once by printing the buffer. I use a `[]rune` slice as a buffer because `rune` can store an emoji character.

* Make a large enough rune slice named `buf` using the `make` function.

  * **HINT:** `width * height` will give you a large enough buffer.

  * **TIP:** You could also use `string` concatenation to draw into a `string` buffer but it would be inefficient.

  * You will find more information about bytes and runes in the strings section.

```go
// TIP for converting the buffer
var buffer []rune

// For printing, you can convert a rune slice to a string like so:
str := string(buffer)
```

## SLOWING DOWN THE SPEED

Call the `time.Sleep` function to slow down the speed of the loop a little bit, so you can see the ball :)

```go
time.Sleep(time.Second / 20)
```
