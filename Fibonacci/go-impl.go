// Fibonacci with multiple implementations
package main

import (
	"fmt"
	"time"
)

func basicFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return basicFibonacci(n-1) + basicFibonacci(n-2)
}

func memoizedFibonacci(n int) int {
	memo := make([]int, n+1)
	return memoizedFibonacciHelper(n, memo)
}

func memoizedFibonacciHelper(n int, memo []int) int {
	if n <= 1 {
		return n
	}
	if memo[n] == 0 {
		memo[n] = memoizedFibonacciHelper(n-1, memo) + memoizedFibonacciHelper(n-2, memo)
	}
	return memo[n]
}

func bottomUpFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func main() {
	// start1 := time.Now()
	// fmt.Println(basicFibonacci(20))
	// elapsed1 := time.Since(start1).Seconds() * 1000
	// fmt.Printf("Execution time: %v ms", elapsed1)
	// fmt.Println()

	start2 := time.Now()
	fmt.Println(memoizedFibonacci(10000))
	elapsed2 := time.Since(start2).Seconds() * 1000
	fmt.Printf("Execution time: %v ms", elapsed2)
	fmt.Println()

	start3 := time.Now()
	fmt.Println(bottomUpFibonacci(10000))
	elapsed3 := time.Since(start3).Seconds() * 1000
	fmt.Printf("Execution time: %v ms", elapsed3)
	fmt.Println()

}
