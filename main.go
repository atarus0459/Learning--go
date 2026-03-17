package main

import "fmt"

func main() {
	score := 75

	if score >= 80 {
		fmt.Println("Grade A")
	} else if score >= 70 {
		fmt.Println("Grade A-")
	} else if score >= 60 {
		fmt.Println("Grade B+")
	}else {
		fmt.Println("Grade C")
	}
}
	