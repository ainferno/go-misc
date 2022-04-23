package main

import "fmt"

func mod(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}

func sqrt2(x0, x1, x2, eps float64) float64 {
	if mod(x2-x1) < eps {
		return x2
	}
	return sqrt2(x0, x2, 0.5*(x2+x0/x2), eps)
}

func sqrt(x, eps float64) float64 {
	if x == 0 {
		return 1
	}
	return sqrt2(x, x, 0.5*(x+1), eps)
}

func main() {
	var x, eps float64
	fmt.Print("Enter epsilon: ")
	fmt.Scanf("%f", &eps)
	fmt.Print("Enter x: ")
	fmt.Scanf("%f", &x)
	fmt.Printf("X=%f, eps=%f\n", x, eps)
	fmt.Printf("Result=%f\n", sqrt(x, eps))
}
