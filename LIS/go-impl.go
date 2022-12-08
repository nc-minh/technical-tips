// Longest Increasing Subsequence
package main

import (
	"fmt"
	"time"
)

func lowerBound(a []int, low int, high int, element int) int {
	for low < high {
		mid := low + (high-low)/2
		if a[mid] < element {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func longestIncreasingSubsequenceLength(a []int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	tail := make([]int, n)
	tail[0] = a[0]
	length := 1
	for i := 1; i < n; i++ {
		if a[i] > tail[length-1] {
			tail[length] = a[i]
			length++
		} else {
			tail[lowerBound(tail, 0, length-1, a[i])] = a[i]
		}
	}
	return length
}

func main() {
	start := time.Now()

	a := []int{2, 5, 3, 7, 11, 8, 10, 13, 6}
	println(longestIncreasingSubsequenceLength(a))

	elapsed := time.Since(start).Seconds() * 1000
	fmt.Printf("Execution time: %v ms", elapsed)
}
