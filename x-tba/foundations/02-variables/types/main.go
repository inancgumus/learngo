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
	"os"
	"runtime"
	"time"
)

func main() {
	// int, float64, bool, string
	var (
		cpus      int
		path, dir string
		max       float64
	)

	cpus = runtime.NumCPU()
	path = os.Getenv("PATH")
	max = math.Max(1.5, 2.5)
	dir, _ = os.Getwd()
	now := time.Now()

	// cpus := runtime.NumCPU()
	// path := os.Getenv("PATH")
	// max := math.Max(1.5, 2.5)
	// dir, _ := os.Getwd()

	// dir = `"` + dir + `"`
	// dir = strconv.Quote(dir)
	// cpus++
	// cpus *= 2.5
	// max++
	// max /= 2.5
	// paths = strings.Split(path, ":")
	// path = paths[0]

	fmt.Printf("# of CPUS        : %d\n", cpus)
	fmt.Printf("Path             : %s\n", path)
	fmt.Printf("max(1.5, 2.5)    : %g\n", max)
	fmt.Printf("Current Directory: %s\n", dir)
	fmt.Printf("Current Time     : %s\n", now)
}
