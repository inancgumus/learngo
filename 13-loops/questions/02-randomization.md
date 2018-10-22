## What's pseudorandom number generation?
1. Numbers appear to be randomly generated but in reality they are not *CORRECT*
2. Generating random numbers according to the physical laws
3. Generating pseudo even and odd numbers

> **1:** Computers are deterministic machines. They can't generate truly random numbers (unlike actual physical processes).


## What's a seed number?
1. Exchanging of random numbers between two computers
2. It's used to getting a random number between 0 and the seed number
3. It's used initialize a pseduorandom number generator *CORRECT*


## Which package is used to generate pseudorandom numbers in Go?
1. pseudorand
2. rand *CORRECT*
3. random
4. randomizer


## What does [0, 5) mean?
1. A range of numbers between 0 and 5 (excluding 5) *CORRECT*
2. A range of numbers between 0 and 5 (including 5)
3. Just 0 and 5
4. Just 0 and 4

> **1:** Right. The square-brace means: "inclusion". The parenthesis means: "exclusion". So, [0, 5] means: 0, 1, 2, 3, 4. It's called the "mathematical interval notation".


## Why this function call would not work?
```go
rand.Intn(0)
```
1. First you should seed it
2. It expects two arguments
3. Intn works within a range of [0, 0). So, it doesn't make sense to include 0 and not include 0 at the same time. *CORRECT*

> **1:** That's not the cause of this error. You don't always have to seed it.
> **2:** No, it does not.


## What does this program print?
Note that, each seed number below returns pseudorandom numbers as these:

```
Seed: 0
  3 3 6 8 4 1 9 3 6 6

Seed: 1
  1 1 9 3 2 4 7 6 6 6

Seed: 2
  10 1 2 2 0 6 4 1 0 5
```

Here's the program:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 3; i++ {
		rand.Seed(int64(i))
		fmt.Print(rand.Intn(11), " ")
		fmt.Print(rand.Intn(11), " ")
	}
}
```
1. 3 1 10 3 1 1
2. 3 6 1 6 10 5
3. 1 10 1 1 3 3
4.  3 3 1 1 10 1 *CORRECT*

> **4:** The numbers are determined depending on the seed number. So, this loop, seeds the pseudorandom generator with 0, 1, and 2 respectively.
> 
> And, after each seed, it calls Intn twice to generate two random numbers.
> 
> So, if you look at the result, 3 3 is the first two numbers of Seed: 0. 1 1 for Seed: 1. And, 10 1 for Seed: 2.


## What you should do if you want the pseudorandom generator to produce random numbers each time you run your program?
1. You need to seed it like this: rand.Seed(rand.Random)
2. You need to seed it like this: rand.Seed(time.Now().UnixNano()) *CORRECT*
3. You need to seed it like this: rand.Seed(time.Now())