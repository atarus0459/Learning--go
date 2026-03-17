package main

import "fmt"

func divide(a float64, b float64) (float64, float64)  {
	return a / b, a * b
}

func main() {
	quotient, product := divide(10, 2)

	fmt.Println("Quotient:", quotient)
	fmt.Println("Product:", product)
}
