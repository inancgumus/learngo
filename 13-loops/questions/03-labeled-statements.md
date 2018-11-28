## Which scope a label belongs to?
1. To the scope of the statement that it is in
2. To the body of the function that it is in *CORRECT*
3. To the package scope that it is in

> **1:** Labels do not belong to statement scopes unlike other names like variable or constant names.
> 
> **2:** They can be used throughout the function, even before their declaration inside the function. This also what makes goto statement jump to any label within a function.


## There two for loop statements below and a label called "words". Which statement does the "words label" label?"
```go
for range words {
words:
    for range letters {
        // ...
    }
}
```
1. The first loop
2. The second, nested loop *CORRECT*
3. All the loops

> **2:** A label can only label one statement at a time.


## Will this loop terminate after the break?
**BTW:** _"terminate" is the same thing with "quit". Remember, statements can also terminate. This means that the statement will end. It doesn't mean that the program will end._

```go
package main

func main() {
main:
	for {
		switch "A" {
        case "A":
            break // <- here!
        case "B":
            continue main
        }
	}
}
```
1. No, the break will only terminate the switch but the loop will continue *CORRECT*
2. Yes, the break will terminate the loop
3. Yes, the break will terminate the switch

> **1:** Yep. This is an unlabeled break. So, it breaks the closest statement, which in here, it's that switch statement. And, since it only breaks the switch, the loop will keep continue.
> 
> **3:** Yep. However, why would that kill the loop as well?


## Will this loop ever terminate?
```go
package main

func main() {
	flag := "A"

main:
	for {
		switch flag {
		case "A":
			flag = "B"
			break
		case "B":
			break main
		}
	}
}

```
1. No, this loop will loop to infinity
2. Yes, the first break will terminate the loop
3. Yes, the second break will terminate the loop *CORRECT*

> **2:** No it does not but it helps.
> 
> **3:** Yep. Do you know why? Because, at first, first case will match, and it will set the flag to "B". So, in the next loop step, the 2nd case will be hit, then, it will break from the loop, because the loop is labeled with the main label.


## What the first break below does?

Note that, in this program, there's an infinite loop.

```go
package main

func main() {
	for {
	switcher:
		switch 1 {
		case 1:
			switch 2 {
			case 2:
				break switcher
			}
		}
		break
	}
}
```
1. It breaks from the 2nd switch causing the program will loop indefinitely
2. It breaks from the 2nd switch and then the 2nd break will terminate the loop
3. It breaks from the 1st switch and then the 2nd break will terminate the loop *CORRECT*

> **1:** There's another break after the switch, so the loop will end immediately.
>
> **2:** It doesn't break the 2nd switch. The label labels the 1st switch.


## What's wrong with this program?

```go
package main

func main() {
	for {
	switcher:
		switch {
		case true:
			switch {
			case false:
				continue switcher
			}
		}
	}
}
```
1. continue statement can only continue a loop *CORRECT*
2. continue statement cannot be used within a switch statement
3. It will loop to infinity

> **1:** Yes! Here, the switcher label labels the first switch statement. So, it's a switch label. And, then it tries to jump to that label using a continue. However, a continue statement can only be used to jump to a loop label.
>
> **2:** Yes, it can be used in a switch statement to continue for the next loop step. But, that's not the problem with this program.
>
> **3:** Yes, but that's not the real problem.


## Which one of these programs will terminate?

I mean: When you run it, which one will quit. Some of the codes here will indefinitely run.

1. ```go
    func main() {
        start: goto exit
        exit : fmt.Println("exiting")
        goto start
    }
   ```
2. ```go
    func main() {
        exit: fmt.Println("exiting")
        goto exit
    }
   ```

3. ```go
    func main() {
        goto exit
        start : goto getout
        exit  : goto start
        getout: fmt.Println("exiting")
    }
   ```
   *CORRECT*

> **1:** In the start label: "goto exit" sends the execution to the exit label. In the exit label: "goto start" sends the execution back to the start label. So, it's an infinite loop. The program will never terminate.
>
> **2:** In the exit label: "goto exit" sends the execution back to the exit label. So, it's an infinite loop. The program will never terminate.
>
> **3:** "goto exit" sends the execution to the exit label. In the exit label: "goto start" sends the execution to the start label. In the start label: "goto getout" sends the execution to the getout label. And, since the getout label is the last statement of the main function, the program will terminate there.