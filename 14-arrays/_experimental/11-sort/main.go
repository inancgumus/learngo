package main

import (
	"fmt"
	"os"
	"strconv"
)

const usage = `Sorry. Go arrays are fixed.
So, for now, I'm only supporting sorting %d numbers...
`

func main() {
	args := os.Args[1:]

	var nums [5]float64

	if len(args) > 5 {
		fmt.Printf(usage, len(nums))
		return
	}

	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		nums[i] = n
	}

	for range nums {
		for i, v := range nums {
			if i < len(nums)-1 && v > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}

	fmt.Println(nums)
}
