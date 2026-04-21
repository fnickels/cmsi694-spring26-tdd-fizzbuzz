package main

import "fmt"

func main() {

	const start, end = 1, 100

	// main loop
	for i := start; i <= end; i++ {
		if i != start {
			fmt.Print(",")
		}

		fmt.Print(FizzBuzzHandler(i))
	}
	fmt.Println()

	// edge cases
	fmt.Println(FizzBuzzHandler(start - 2))
	fmt.Println(FizzBuzzHandler(start - 1))
	fmt.Println(FizzBuzzHandler(end + 1))
}
