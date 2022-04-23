package main

import (
	"fmt"
	"time"
)

func fibonacci_recursive(a uint64) uint64 {
	if a < 0 {
		fmt.Println("Error, number is too big")
		return 0
	}
	if a <= 1 {
		return a
	}
	return fibonacci_recursive(a-1) + fibonacci_recursive(a-2)
}

func fibonacci_iterative(a uint64) uint64 {
	var x, y, z uint64
	y = 1
	if a <= 1 {
		return a
	}
	for ; a > 1; a-- {
		z = y
		y = x + y
		x = z
	}
	return y
}

func main() {
	var a uint64
	fmt.Scanf("%d", &a)
	start1 := time.Now()
	fmt.Println(fibonacci_recursive(a))
	end1 := time.Now()
	fmt.Println("Recursion took ", end1.Sub(start1))

	start2 := time.Now()
	fmt.Println(fibonacci_iterative(a))
	end2 := time.Now()
	fmt.Println("Iteration took ", end2.Sub(start2))
}
