package main

import "fmt"

func multiply(a int, b int) (result int) {
	result = a * b
	return result
}

func main() {
	result := multiply(2, 4)
	fmt.Println(result)
}
