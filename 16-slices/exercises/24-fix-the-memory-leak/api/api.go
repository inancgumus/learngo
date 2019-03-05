package api

import (
	"fmt"
	"math/rand"
	"runtime"
)

// DO NOT TOUCH THE FOLLOWING CODE
// THIS IS THE API
// YOU CANNOT CONTROL IT! :)

// Read returns a huge slice (allocates ~65 MB of memory)
func Read() []int {
	return rand.Perm(2 << 22)
}

// Report cleans the memory and prints the current memory usage
func Report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
