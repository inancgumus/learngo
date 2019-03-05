package main

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got a requests log for a system in a slice: `reqs`.
//  The log contains the total request counts per 8 hours.
//
//  The reqs slice is a single-dimensional slice but you need
//  to group the requests daily into a slice named: `daily`.
//
//  Please follow the instructions inside the code to solve
//  the exercise.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	//
	// #1: DAILY REQUESTS DATA
	//
	// . The system collects and groups the requests per 8 hours
	//
	// . So there are 3 requests totals per day
	//
	// . Your code should be robust enough to work with
	//   insufficient data. For example, in the last day
	//   there are only two request totals: 100 and 150.
	//
	//   So, don't forget to handle that edge case as well.
	//
	// . Uncomment the code below and start
	//

	// reqs := []int{
	//	// 1st day
	//	500, 600, 250,
	//	// 2nd day
	//	200, 400, 50,
	//	// 3rd day
	//	900, 800, 600,
	//	// 4th day
	//	750, 250, 100,
	//	// last day
	//	100, 150,
	// }

	//
	// #2: Group the `reqs` per day into a slice named: `daily`
	//
	// . Create the daily slice using the `make` function
	//
	// . Anticipate the length argument to the make function using this data:
	//
	//   + The length of the reqs slice
	//   + There are 3 requests totals per day
	//   + There are residual elements (the last day)
	//
	//   ! So, do not blindly allocate a slice.
	//
	//   ! Allocate the slice efficiently with the exact size needed.
	//
	// . Then append to it the daily requests in a "loop"
	//

	//
	// #3: Print the header:
	//
	// Day       Requests
	// ====================
	//

	//
	// #4: Print the data per day along with the totals:
	//
	// 1         500
	// 1         600
	// 1         250
	// --------------------
	//           1350  --> Print the daily total requests
	//
	// 2         200
	// 2         400
	// 2         50
	// --------------------
	//           650
	//
	//           2000  --> Also print the grand total

	// ------------------------------------------------------------------------
	//
	// ❤️ NOTE ❤️
	//
	// If you could't solve this challenge, please do not get discouraged.
	//
	// Look at the solution, then try to solve it again. This is valuable too.
	//
	// Then ️change the request data, the number of requests per day (now it's 3),
	// etc., and then try to solve it again.
	//
	// ------------------------------------------------------------------------
}
