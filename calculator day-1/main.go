package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		a       int
		b       int
		name    string
		command int
	)
	fmt.Println("Please enter your first name.")
	fmt.Scan(&name)
	fmt.Printf("Welcome %v!\n", strings.ToUpper(strings.TrimSpace(name)))

	for {

		fmt.Println("Enter a digt")
		_, err1 := fmt.Scan(&a)
		if err1 != nil {
			fmt.Println(err1)
			continue
		}

		fmt.Println("Enter a second digit")
		_, err2 := fmt.Scan(&b)
		if err2 != nil {
			fmt.Println(err2)
			continue
		}

		fmt.Println("choose an operation: 1 to add | 2 to subtract | 3 to divide | 4 to multiply | 5 to quit programme.")
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println(err)
			fmt.Scan(&command)
			continue
		}
		if command > 5 && command < 1 {
			fmt.Println("Error: enter a valid operation!")
			continue
		}
		if command == 1 {
			fmt.Println("result:", add(a, b))
			continue
		}
		if command == 2 {
			fmt.Println("result:", sub(a, b))
			continue
		}
		if command == 3 {
			_, err3 := fmt.Println("result:", modulus(a, b))
			if err3 != nil {

			}
			continue
		}
		if command == 4 {
			fmt.Println("result:", mult(a, b))
			continue
		}
		if command == 5 {
			fmt.Printf("goodbye %v", name)
			break
		}
	}
}
