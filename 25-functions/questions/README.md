## Which one is a correct function declaration?
1. add func(a, b int) {}
2. function run(a int, b int) {}
3. func run(int a, b) {}
4. func run(a, b int) {} *CORRECT*


## Which one is the input and result names of the following func?
```go
func run(p Process, id1, id2 int) (pid int, err error) {}
```
1. Inputs: p, id1, and id2. Results: pid, err. *CORRECT*
2. Inputs: Process, int, int. Results: int, error.
3. Inputs: run, p, id1, id2. Results: pid, err.
4. The declaration syntax is wrong.


## What is a return statement?
1. Terminates a program.
2. Terminates a func by returning zero or more values to a caller func. *CORRECT*
3. Skips the next statement and runs the next.


## What is wrong with the following code?
```go
func add(a, b int) {
    return a + b
    return
}
```
1. The return statement should be called only once
2. The last return statement is wrong
3. It should declare an int result value and remove the last return statement *CORRECT*
4. It should declare any numeric result value

> **2:** Actually, it is correct because the func doesn't declare a result value.
> 
> **3:** Correct. It should be: `func add(a, b int) int { return a + b }`


## How to fix the following code?
```go
func incr(a int) {
    a++
    return
}

num := 10

// You want it to print 11 but it prints 10 instead.
fmt.Println( incr(num) )
```
1. Change the func: `func incr(a int) int { a++; return a }` *CORRECT*
2. Change the func: `func incr(a int, newA int) { a++; newA = a }`
3. Change the func: `func incr(a int) int { return a++ }`

> **1:** Go is a 100% pass-by-value language. So, the inputs to a func are local to that function: The changes are not visible outside of that func.


## Why should you not use package level variables?
1. Nothing is wrong with them
2. Funcs cannot use package level variables
3. They may increase code coupling and cause fragile code *CORRECT*

> **3:** It's because: Anyone can access and change them.


## Why should you return an error from the following func?
```go
// Why this?
func incr(n string) (int, error) {
	m, err := strconv.Atoi(n)
	return n + 1, err
}

// Instead of this?
func incr(n string) int {
	m, _ := strconv.Atoi(n)
	return m + 1
}
```
1. You want to let the caller know when something goes wrong *CORRECT*
2. When an error occurs, `Atoi` returns 0, so you don't need to return an error

> **2:** Sometimes, this is partly true however it is better to let the caller know when something goes wrong.


## How and why does the following return statement work?
```go
func spread(samples int, P int) (estimated float64) {
	for i := 0; i < P; i++ {
		estimated += estimate(i, P)
	}
	return
}
```
1. `estimated` is a named result value. So the naked return returns `estimated` automatically. *CORRECT*
2. return statement is not necessary there. Go automatically returns `estimated`.
3. Result values cannot have a name. This code is incorrect.


## Does the following code work? If so, why?
IT SHOULD PRINT: map[1:11 10:3]
```go
func main() {
    stats := map[int]int{1: 10, 10: 2}
    incrAll(stats)
    fmt.Print(stats)
}

func incrAll(stats map[int]int) {
    for k := range stats {
        stats[k]++
    }
}
```
1. No, it doesn't work: Go is a pass by value language. `incrAll` cannot update the map value.
2. Yes, it works: `incrAll` can update the map value. *CORRECT*

> **2:** Map values are pointers. So, `incrAll` can update the map value.


## Does the following code work? If so, why?
IT SHOULD PRINT: [10 5 2]
```go
func main() {
    stats := []int{10, 5}
    add(stats, 2)
    fmt.Print(stats)
}

func add(stats []int, n int) {
    stats = append(stats, n)
}
```
1. No, it doesn't work: `add()` cannot update the original slice header. *CORRECT*
2. Yes, it works: `add()` can add new element to the original slice header.

> **1:** Go is a pass-by-value programming language. add() creates a copy of the original slice header and adds the new element to the new slice header but it never returns the updated one. So, it cannot update the original slice header. It should have been returning the original slice header.