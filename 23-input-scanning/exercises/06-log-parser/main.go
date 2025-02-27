// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
var lines int
	var total int
	slice := make([]string, 0)
	sum := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		lines++
		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Println("ERROR")
			return
		}
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("ERROR in %d", lines)
			return
		}
		total += visits
		if _, ok := sum[domain]; !ok {
			slice = append(slice, domain)
		}
		sum[domain] = sum[domain] + visits

	}
	fmt.Printf("%-30s%10s", "DOMAIN", "VISITS")
	fmt.Print(strings.Repeat("-", 45))

	for value := range slice {
		visits := sum[value]
		fmt.Printf("%-20s%5d", value, visits)
		fmt.Println()
	}
	if err := in.Err(); err != nil {
		fmt.Println("ERROR")
		return
	}






















  
}
