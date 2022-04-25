package main

import "fmt"

func bubble_sort(a []int) []int {
	len_a := len(a)
	for i := 0; i < len_a; i++ {
		for j := 0; j < len_a-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func main() {
	a := []int{1, 2, 7, 5, 6, 3, 4}
	fmt.Println("Raw:    ", a)
	fmt.Println("Sorted: ", bubble_sort(a))
}
