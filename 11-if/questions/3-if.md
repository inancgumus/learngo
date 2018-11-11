## What does "control flow" mean?
1. Changing the top-to-bottom execution of a program
2. Changing the left-to-right execution of a program
3. Deciding which statements are executed *CORRECT*

> **1, 2:** You can't change that.
> 
> **3:** That's right. Control flow allows us to decide which parts of a program is to be run depending on condition values such as true or false.


## How can you simplify the condition expression of this if statement?
```go
if (mood == "perfect") {
    // this code is not important
}
```

1. `if {mood == "perfect"}`
2. `if [mood == "perfect"]`
3. `if mood = "perfect"`
4. `if mood == "perfect"` *CORRECT*

> **1, 2:** That's a syntax error. Try again.
> 
> **3:** `=` is the assignment operator. It cannot be used as a condition.
> 
> **4:** That's right. In Go, you don't need to use the parentheses.


## The following code doesn't work. How you can fix it?
```go
package main

import "fmt"

func main() {
    // this program prints "cool"
    // when the mood is "happy"

    mood := "happy"

    if "happy" {
        fmt.Println("cool")
    }
}
```
1. Just wrap the "happy" inside parentheses.
2. You need to compare mood with "happy". Like this: `if mood == "happy" { ... }` *CORRECT*
3. Just directly use `mood` instead of `happy`. Like this: `if mood { ... }`
4. This should work! This is a tricky question.

> **1:** That won't change anything. Go adds the parentheses automatically behind the scenes for every if statement.
>
> **2:** Yep. In Go, condition expressions always yield a bool value. Using a comparison operator will yield a bool value. So, it will work.
>
> **4:** No, it's not :)


## How can you simplify the following code? You only need to change the condition expression, but how?
```go
package main

import "fmt"

func main() {
    happy := true

    if happy == true {
        fmt.Println("cool!")
    }
}
```
1. `happy != false`
2. `!happy == false`
3. `happy` *CORRECT*
4. `!happy == true`

> **1, 2:** Right! But you can do it better.
>
> **3:** Perfect! You don't need to compare the value to `true`. `happy` is already true, so it'll print "cool!".
>
> **4:** That won't print anything. `!happy` yields false.


## How can you simplify the following code? You only need to change the condition expression, but how?
```go
package main

import "fmt"

func main() {
    happy := false

    if happy == !true {
        fmt.Println("why not?")
    }
}
```
1. Easy! Like this: `happy != true`
2. `!happy` *CORRECT*
3. `happy == false`
4. `!happy == false`

> **1, 3:** Right! But you can do it better.
>
> **2:** Perfect! You don't need to compare the value to `false` or to `!true` (which is false). `!happy` already returns true, because it's false at the beginning.
>
> **4:** That won't print anything. `happy` will be true.


## This code contains an error. How to fix it?
```go
package main

import "fmt"

func main() {
    happy := false

    if happy {
        fmt.Println("cool!")
    } else if !happy {
        fmt.Println("why not?")
    } else {
        fmt.Println("why not?")
    } else {
        fmt.Println("why not?")
    }
}
```
1. Remove one of the else branches. *CORRECT*
2. Move the else if as the last branch.
3. It repeats "why not?" several times.
4. Remove the `else if` branch.

> **1:** Right. There can be only one else branch.
> 
> **2:** If there's an else branch, you can't move else if branch as the last branch.
>
> **3, 4:** So? :) That's not the cause of the problem.
>


## What's the problem with the following code?
```go
package main

import "fmt"

func main() {
    happy := true
    energic := happy

    if happy {
        fmt.Println("cool!")
    } else if !happy {
        fmt.Println("why not?")
    } else if energic {
        fmt.Println("working out?")
    }
}
```
1. It declares the energic variable unnecessarily.
2. You can't use more than one else if branch.
3. It will never run the last else if branch. *CORRECT*
4. There's no else branch.

> **2:** Well, actually you can.
> 
> **3:** Right. `happy` can only be either true or false. That means, it will always execute the first two branches, but it will never execute the else if branch.
> 
> **4:** It doesn't have to be. Else branch is optional.


## How can you simplify the following code?
```go
package main

import "fmt"

func main() {
    happy := false

    if happy {
        fmt.Println("cool!")
    } else if happy != true {
        fmt.Println("why not?")
    } else {
        fmt.Println("why not?")
    }
}
```
1. Change `else if`'s condition to: `!happy`.
2. Move the else branch before else if.
3. Remove the else branch.
4. Remove the else if branch. *CORRECT*

> **1, 3:** Close! But, you can do it even better.
>
> **2:** You can't: `else` branch should be the last branch.
>
> **4:** Cool. That's not necessary because `else` branch already handless "unhappy" situation. It's simpler because it doesn't have a condition.