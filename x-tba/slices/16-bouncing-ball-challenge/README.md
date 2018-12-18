# Bouncing Ball Challenge

This document contains what you need to do to create the bouncing ball animation.You're going to learn a great deal of knowledge about slices and you'll gain a good experience while doing this.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## Declare the constants
* The width and the height of the board.

* It can be anything, just experiment.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## Declare the ball positions and velocity:
The ball should be in a known position on your board, and it should have a velocity.

* **Velocity means: Speed and Direction**

  * X velocity =  1 -> _ball moves right_
  * X velocity = -1 -> _ball moves left_
  * Y velocity =  1 -> _ball moves down_
  * Y velocity = -1 -> _ball moves up_

* Declare variables for the ball's positions: `X` and `Y`

* Declare variables for the ball's velocity: `xVelocity` and `yVelocity`

* On each step: Add velocities to ball's position. This will make the ball move.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## 🎾 CREATE THE BOARD

You can use a multi-dimensional bool slice to create the board. Each element in the slice corresponds to whether the ball is in that element (or position) or not.

* Use the `make` function to initialize your board.

  * Remember: Zero value for a slice is `nil`.
  * So, you need to initialize each sub-slice of the board slice.

* You can use `[][]bool` for your board.

  * It's because, when you set one of the elements to true, then you'd know that the ball is in that position.

  * *EXAMPLE:*
  ```
    false    false    false   false
    false    true -+  false   false
    false    false |  false   false
                   v
           the ball is here
           board[1][1] is true
   ```

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## 🎾 CLEAR THE SCREEN

* Before the loop, clear the screen once by using the screen package.

* It's [here](https://github.com/inancgumus/screen).
* Its documentation is [here](https://godoc.org/github.com/inancgumus/screen).

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖
 
## 🎾 DRAWING LOOP
This is the main loop that will draw the board and the ball to the screen continuously.

* Create the loop

* On each step of the loop, you're going to:
  * Clear the board
  * Calculate the next ball position
  * Draw the board and the balls

Explanations for these steps follow...

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## 🎾 CLEAR THE BOARD
At the beginning your board should not contain the ball.

* So, set all the inner slices of the board slice to `false`.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## 🎾 CALCULATE THE NEXT BALL POSITION
You need to calculate the ball's next position on the board.

* Add the velocities to the ball's current position:

  * `X += xVelocity`
  * `Y += yVelocity`

* When the ball hits the borders of the board change its direction.

  * **HINT:** You can multiply any velocity by `-1` to change its direction.

* Set the ball's position on the board.

  * You will use this information when drawing the board.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## 🎾 DRAW THE BOARD
Instead of drawing the board and the ball to the screen everytime, you will fill a buffer, and when you complete, you will draw the whole board and the ball once by printing the buffer that you filled.

* Make a large enough rune slice named `buf` using the `make` function.

* **HINT:** `width * height` will give you a large enough buffer.

* **TIP:** You could also use `string` concatenation but it would be inefficient.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## FILL YOUR BUFFER:

* Filling your buffer:

  * Loop for the height of the board (_described below_).

  * Then in a nested loop (_described below_), loop for the width of the board.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## NESTEP LOOP: WIDTH LOOP
* On each step of the nested loop, do this:

  * Check whether the ball is in the x, y positions.
    * You need to check for it using your board slice.

  * If so, `append` this tennis ball '🎾' to the buffer.
  * If not, `append` this pool ball '🎱' to the buffer.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## NESTEP LOOP: HEIGHT LOOP

* After the nested loop (but in the height loop):

* `append` the newline character to the buffer: `'\n'`
  * This will allow you to print the next row to the next line.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## 🎾 MOVE THE CURSOR

* After the loop, move the cursor to the top-left position by using the screen package.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## PRINT THE BOARD (USING THE BUFFER):

* Do not forget converting it to `string`. Because your buffer is `[]rune`, like so:

  `fmt.Print(string(buf))`

* You'll learn the details about rune and strings later.

➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖

## SLOW DOWN THE SPEED

And lastly, call the `time.Sleep` function to slow down the speed of the loop, so you can see the ball :) Like so:

  `time.Sleep(time.Millisecond * 60)`
