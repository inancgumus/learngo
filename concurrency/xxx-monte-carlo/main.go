// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		exit("Please provide the sample rate")
	}

	samples, err := strconv.Atoi(os.Args[1])
	if err != nil {
		exit("Incorrect sample rate")
	}

	elapse := measure(time.Now())
	pi := spread(samples, runtime.NumCPU())

	fmt.Println("PI  : ", pi)
	fmt.Println("Time: ", elapse())
}

func exit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func measure(start time.Time) func() time.Duration {
	return func() time.Duration {
		return time.Since(start)
	}
}

// spread the work to P number of estimate go funcs.
// returns the estimation.
func spread(samples int, P int) (estimated float64) {
	counts := make(chan float64)

	for i := 0; i < P; i++ {
		go func() { counts <- estimate(samples / P) }()
	}

	for i := 0; i < P; i++ {
		estimated += <-counts
	}
	return estimated / float64(P)
}

func estimate(N int) float64 {
	const radius = 1.0

	var (
		seed   = rand.NewSource(time.Now().UnixNano())
		random = rand.New(seed)
		inside int
	)

	for i := 0; i < N; i++ {
		x, y := random.Float64(), random.Float64()

		if num := math.Sqrt(x*x + y*y); num < radius {
			inside++
		}
	}
	return 4 * float64(inside) / float64(N)
}
