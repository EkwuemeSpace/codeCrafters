package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	if a < 0 && b < 0 {
		return a + b
	}
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func modulus(a, b int) int {
	if a == 0 {
		fmt.Println("error cannot divide by 0")
		fmt.Println("Please enter a valid digit")
		_, err1 := fmt.Scan(&a)
		if err1 != nil {
			fmt.Println(err1)
			fmt.Scan(a)
		}
	}
	if b == 0 {
		fmt.Println("error cannot divide by 0")
		fmt.Println("please enter a valid digit")
		_, err2 := fmt.Scan(&b)
		if err2 != nil {
			fmt.Println(err2)
			fmt.Scan(b)
		}
	}
	return a / b
}
