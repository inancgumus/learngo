// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

func main() {
	// [LENGTH]TYPE {
	//   ELEMENTS
	// }

	// LENGTH=2 and TYPE=[2]int

	// nums := [2][2]int{
	// 	[2]int{2, 4},
	// 	[2]int{1, 3},
	// }

	// code below is the same as the code above
	nums := [2][2]int{
		{2, 4},
		{1, 3},
	}

	fmt.Println("nums         =", nums)
	fmt.Println("nums[0]      =", nums[0])
	fmt.Println("nums[1]      =", nums[1])

	fmt.Println("nums[0][0]   =", nums[0][0])
	fmt.Println("nums[0][1]   =", nums[0][1])
	fmt.Println("nums[1][0]   =", nums[1][0])
	fmt.Println("nums[1][1]   =", nums[1][1])

	fmt.Println("len(nums)    =", len(nums))
	fmt.Println("len(nums[0]) =", len(nums[0]))
	fmt.Println("len(nums[1]) =", len(nums[1]))

	for i, array := range nums {
		for j, n := range array {
			// nums[i][j] = number
			fmt.Printf("nums[%d][%d] = %d\n", i, j, n)
		}
	}
}
