// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bytes"
	"fmt"
	"testing"
	"unicode"
	"unicode/utf8"
)

/*
Let's run this program to see the speed differences.

benchDecoder  30000000          46.2 ns/op --> BEST
benchForRange 30000000          53.0 ns/op --> MEDIOCRE
benchConcater 20000000          93.7 ns/op --> WORST
*/

var word = []byte("öykü")

func decoder(w []byte) {
	_, size := utf8.DecodeRune(word)
	copy(w[:size], bytes.ToUpper(w[:size]))
}

func forRange(w []byte) {
	var size int
	for i := range string(w) {
		if i > 0 {
			size = i
			break
		}
	}
	copy(w[:size], bytes.ToUpper(w[:size]))
}

var globalString string

func concater(w []byte) {
	runes := []rune(string(w))
	runes[0] = unicode.ToUpper(runes[0])
	globalString = string(runes)
}

func bench(technique func([]byte)) testing.BenchmarkResult {
	return testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			technique(word)
		}
	})
}

func main() {
	fmt.Println("benchDecoder", bench(decoder))
	fmt.Println("benchForRange", bench(forRange))
	fmt.Println("benchConcater", bench(concater))
}
