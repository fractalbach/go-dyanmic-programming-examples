package main

import (
	"fmt"
	"time"
)


func powerset(set ...interface{}) []interface{} {
	var output []interface{}
	for i, item := range set {
		output = append(output, []interface{}{item})
		subset := []interface{}{item}
		for j := i + 1; j < len(set); j++ {
			subset = append(subset, set[j])
			output = append(output, subset)
		}
	}
	return output
}

func runTest(n int) time.Duration {
	args := make([]interface{}, n)
	for i:=0; i<n; i++ {
		args[i] = i
	}
	t0 := time.Now()
	powerset(args...)
	tf := time.Now()
	dt := tf.Sub(t0)
	fmt.Printf("n%d: %v\n", n, dt)
	return dt
}

func main() {
	fmt.Println("Results:")
	// for _, v := range s {
	// 	fmt.Printf("%v\n", v)
	// }
	var timeList []time.Duration
	for i:=0; i<1000; i += 100 {
		timeList = append( timeList, runTest(i)	)
	}

}
