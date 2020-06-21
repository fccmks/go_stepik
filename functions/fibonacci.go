package main

import "fmt"

func main() {
	fmt.Print(fibonacci(3))
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		fibslice := make([]int, n)
		fibslice[0] = 1
		fibslice[1] = 1
		for i := 2; i < n; i++ {
			fibslice[i] = fibslice[i-1] + fibslice[i-2]
		}
		return fibslice[n-1]
	}
}
