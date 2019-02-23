// Running up the Stairs

// You can take either 1, 2, or 3 steps up the stairs at a time.  How
// many different ways can you go up a staircase with n steps?

// Overview:

/*

- Identify the base cases:  f(1), f(2), f(3)

- Create data storage for base cases and future solutions. In this
  problem, an array could be used, but I will use a hash map because
  it makes it a bit easier to confirm if a value is absent.

- Find a recursive formula:  f(n) = f(n-1) + f(n-2) + f(n-3)

- After calling f(n) check to see if you already have the solution. If
  you don't have it yet, then execute the recursive formula.

- Save the solution after calculating a new f(n) value.

- Let recursion slowly diminish n down to 0 until the program finishes.

*/

package main

import (
	"fmt"
	"time"
)

var memo = map[int]int{
	1: 1, // f(1) = 1 because 1 = 1
	2: 2, // f(2) = 2 because 2 = 2 = 1+1
	3: 4, // f(3) = 4 because 3 = 3 = 2+1 = 1+2 = 1+1+1
}

// nWays recursively solves for the number of ways you can go up the
// staircase, using memoization to store solutions as they are
// discoverd.  A hashmap is used as storage, although in this problem:
// a simple array would work just as well.
func nWays(n int) int {
	if n <= 0 {
		return 0
	}
	v, ok := memo[n]
	if ok {
		return v
	}
	// ways := 0
	// for i := 1; i < n; i++ {
	// 	ways += nWays(i)
	// 	ways += nWays(n - i)
	// }
	ways := nWays(n-1) + nWays(n-2) + nWays(n-3)
	memo[n] = ways
	return ways
}

// nWaysVerbose does the same thing as nWays, but will additionally
// show the process of how nWays(n) is solved, revealing its tree-like
// recursive structure.
func nWaysVerbose(n int) int {
	if n <= 0 {
		return 0
	}
	v, ok := memo[n]
	if ok {
		return v
	}
	ways1 := nWaysVerbose(n - 1)
	ways2 := nWaysVerbose(n - 2)
	ways3 := nWaysVerbose(n - 3)
	ways := ways1 + ways2 + ways3
	fmt.Printf(
		"f(%d) = f(%d) + f(%d) + f(%d)  =  %d + %d + %d  =  %d\n",
		n,
		n-1, n-2, n-3,
		ways1, ways2, ways3, ways,
	)
	memo[n] = ways
	return ways
}

// printMap displays the saved solutions that have been discovered
// using nWays.
func printMap() {
	fmt.Println("SavedSolutions:")
	for i := 0; i <= len(memo); i++ {
		v := memo[i]
		fmt.Printf("f(%d) = %d\n", i, v)
	}
}

func main() {
	fmt.Println("Staircase:")
	var (
		n         = 40
		startTime = time.Now()
		answer    = nWaysVerbose(n)
		endTime   = time.Now()
		runTime   = endTime.Sub(startTime)
	)
	printMap()
	fmt.Println("Answer:", answer)
	fmt.Println("TimeTaken:", runTime)
}
