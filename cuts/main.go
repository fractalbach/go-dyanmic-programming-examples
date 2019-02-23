/*Dynamic Programming Rod Cutting Example

Rod-Cutting Problem from Introduction to Algorithms, 3rd edition.


Given a rod of length n and a table of prices p_i for i = 1,2,...,n
determine the maximum revenue r_n obtainable by cutting up the rod and
selling the pieces.


This problem is one which can be solved using dynamic programming.
It's a problem that arises in different contexts, but is common enough
that the similarities can be seen.

Some of these similarities are:

1. Breaking problems apart into smaller versions of the same problem

2. Seeing the same problem show up multiple times.


*/
package main

import (
	"fmt"
)

// PriceTable is a lookup table that tells you the known prices of rods
// given their length.
type PriceTable struct {
	values map[int]int
}

// Lookup retrieves a value from the price table
func (pt *PriceTable) Lookup(n int) int {
	v, ok := pt.values[n]
	if ok {
		return v
	}
	panic(fmt.Sprint("cannot lookup value:", n))
	// return 0
}

func (pt *PriceTable) ShowPrices(lengthList []int) int {
	sum := 0
	for _, v := range(lengthList) {
		fmt.Print()
		sum += pt.Lookup(v)
	}
	return sum
}

func (pt *PriceTable) AddValueIfNeeded(k, v int) {
	if _, ok := pt.values[k] ; ok {
		return
	}
	pt.values[k] = v
}

// MyPriceTable Sample Price Table: (length of rod) -> (price) length
// -> value
var MyPriceTable = &PriceTable{
	values: map[int]int{
		1:  1,
		2:  5,
		3:  8,
		4:  9,
		5:  10,
		6:  17,
		7:  7,
		8:  20,
		9:  24,
		10: 30,
	},
}

func addSpaces(depth int) {
	s := ""
	for i := 0; i < depth; i++ {
		s += "  "
	}
	fmt.Print(s)
}

var depth = 0

func topdown(prices *PriceTable, n int) int {
	addSpaces(depth)
	fmt.Printf("call f(%d)\n", n)
	if n <= 0 {
		return 0
	}
	best := -1
	for i := 1; i < n+1; i++ {
		depth++
		test := prices.Lookup(i)+topdown(prices, n-i)
		depth--
		if test > best {
			best = test
			addSpaces(depth)
			fmt.Printf("f(%d) = %d\n", n, best)
		}
	}
	return best
}

// return the value and a list of the cut lengths.
func naive(prices *PriceTable, n int) (int, []int) {
	// fmt.Printf("call  f(%d)\n", n)
	best := -1
	bestCuts := []int{}
	if n <= 0 {
		return 0, bestCuts
	}	
	for i := 1; i < n+1; i++ {
		
		fmt.Printf("cut   f(%d) = f(%d) + f(%d)\n", n, i, n-i)
		prev, prevCuts := naive(prices, n-i)
		current := prices.Lookup(i) + prev
		cuts := append([]int{i}, prevCuts...)
		fmt.Println(eqString(n, cuts))
		
		if current > best {
			bestCuts = cuts
			best = current
		}
	}
	return best, bestCuts
}

func eqString(n int, cuts []int) string {
	s := fmt.Sprint(n, " = ")
	for i, v := range cuts {
		s += fmt.Sprint(v)
		if i != len(cuts)-1 {
			s += " + "
		}
	}
	return s
}

func main() {
	// best := topdown(MyPriceTable, 2)
	best, bestCuts := naive(MyPriceTable, 7)
	fmt.Println("\n__________[ Results ]__________")
	fmt.Println("\nfinal cuts:", bestCuts)
	fmt.Println("\nfinal result:", best)
}
