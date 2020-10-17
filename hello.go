package main

import "fmt"

func main() {
		fmt.Println("Hello, World!")
		fmt.Println("Sum of 2 + 2 is", addNums(2, 2))
}

func addNums(num1 int, num2 int) int {
	return num1 + num2
}