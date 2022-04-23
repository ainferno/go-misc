package main

import "fmt"

func polynomial(x, a, b float64) float64 {
	return a + x*b
}

func main() {
	var x, a, b float64
	a = 0
	b = 0
	fmt.Print("Enter x: ")
	fmt.Scanf("%f", &x)
	for i := 0; ; i++ {
		fmt.Print("Enter a", i, ": ")
		_, err := fmt.Scanf("%f", &a)
		if err != nil {
			break
		}
		b = polynomial(x, a, b)
	}
	fmt.Println("\nResult =", b)
}
