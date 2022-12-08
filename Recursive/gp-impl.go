package main

import (
	"fmt"
	"time"
)

func fib(n int) int {

	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)

}

func main() {
	start := time.Now()

	fmt.Println(fib(48))

	elapsed := time.Since(start).Seconds() * 1000
	fmt.Printf("Execution time: %v ms", elapsed)
}
