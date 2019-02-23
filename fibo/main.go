/*

Todo List:

- clean up the old code and only leave the math/big stuff.

- save data into files.

- load data from files.

- do some benchmarking

*/

package main

import (
	"fmt"
	"math/big"
)

var memo = map[int]int{
	0: 0,
	1: 1,
}

var bigmemo = map[int]*big.Int{
	0: big.NewInt(0),
	1: big.NewInt(1),
}

func bigfib(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if v, ok := bigmemo[n]; ok {
		return v
	}
	result := big.NewInt(0)
	result.Add(bigfib(n-1), bigfib(n-2))
	bigmemo[n] = result
	return result
}

func fib(n int) int {
	if n < 0 {
		return 0
	}
	if v, ok := memo[n]; ok {
		return v
	}
	result := fib(n-1) + fib(n-2)
	memo[n] = result
	return result
}

func printMemo(n int) {
	fmt.Println("MEMO {")
	for i := 0; i <= n; i++ {
		if v, ok := memo[i]; ok {
			fmt.Printf("  %d: %d\n", i, v)
		}
	}
	fmt.Println("}")
}

func showFib(n int) {
	ans := fib(n)
	fmt.Printf("fib(%d) = %d\n", n, ans)
	printMemo(n)
}

func showBigFib(n int) {
	bigfib(n)
	for i:=0; i<=n; i++ {
		if v, ok := bigmemo[i] ; ok {
			fmt.Println(i, v)
		}
	}
}

func main() {
	showBigFib(200)
}
