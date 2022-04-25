package main

import "fmt"

func select_sort(a []int) []int {
	len_a := len(a)
	for i := 0; i < len_a; i++ {
		b := i
		for j := i; j < len_a; j++ {
			if a[j] < a[b] {
				b = j
			}
		}
		a[i], a[b] = a[b], a[i]
	}

	return a
}

func main() {
	a := []int{1, 2, 7, 5, 6, 3, 4}
	fmt.Println("Raw:    ", a)
	fmt.Println("Sorted: ", select_sort(a))
}
