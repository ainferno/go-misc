package main

import "fmt"

func merge(a, b []int) []int {
	i := 0
	j := 0
	len_a := len(a)
	len_b := len(b)
	var result []int
	for i+j < len_a+len_b {
		if i < len_a && j < len_b {
			if a[i] > b[j] {
				result = append(result, b[j])
				j++
			} else {
				result = append(result, a[i])
				i++
			}
		} else if i < len_a && j >= len_b {
			result = append(result, a[i])
			i++
		} else if i >= len_a && j < len_b {
			result = append(result, b[j])
			j++
		}
	}
	return result
}

func merge_sort(a []int) []int {
	len := len(a)
	if len <= 1 {
		return a
	}
	return merge(merge_sort(a[:len/2+len%2]), merge_sort(a[len/2+len%2:]))
}

func main() {
	a := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Raw:    ", a)
	fmt.Println("Sorted: ", merge_sort(a))
}
