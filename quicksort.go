package main

import "fmt"

func sum_arrays(a, b []int) []int {
	for _, elem := range b {
		a = append(a, elem)
	}
	return a
}

func split_array(a []int, b int) ([]int, []int) {
	var result_less, result_more []int
	for i, _ := range a {
		if a[i] < b {
			result_less = append(result_less, a[i])
		}
		if a[i] > b {
			result_more = append(result_more, a[i])
		}
	}
	return result_less, result_more
}

func quicksort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	pivot := a[0]
	res_less, res_more := split_array(a, pivot)
	res_less = append(res_less, pivot)
	return sum_arrays(quicksort(res_less), quicksort(res_more))
}

func main() {
	a := []int{5, 4, 3, 2, 1, 6, 7, 10}
	fmt.Println("Raw:    ", a)
	fmt.Println("Sorted: ", quicksort(a))
}
