// Running up the Stairs

// You can take either 1, 2, or 3 steps up the stairs at a time.  How
// many different ways can you go up a staircase with n steps?

// My Process

// - Create the base cases: 1,2,3
// - Create a map to store solutions.
// - produce combinations in the form n = a + b
// - solve for f(n) = f(a) + f(b) when n > 3
// - populate the map with newly discovered values

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
	ways := 0
	for i := 1; i < n; i++ {
		ways += nWays(i)
		ways += nWays(n - i)
	}
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
	fmt.Printf("f(%d) := 0\n", n)
	ways := 0
	for i := 1; i < n; i++ {
		a := i
		b := n - i
		aWays := nWaysVerbose(a)
		bWays := nWaysVerbose(b)
		ways += aWays + bWays
		fmt.Printf(
			"f(%d) += f(%d) + f(%d) = %d + %d\n",
			n, a, b, aWays, bWays)
	}
	memo[n] = ways
	fmt.Printf("f(%d) = %d\n", n, ways)
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
		n         = 20
		startTime = time.Now()
		answer    = nWaysVerbose(n)
		endTime   = time.Now()
		runTime   = endTime.Sub(startTime)
	)
	printMap()
	fmt.Println("Answer:", answer)
	fmt.Println("TimeTaken:", runTime)
}
