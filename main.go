package main

import "fmt"

func main() {

	const start, end = 1, 100

	for i := 1; i <= 100; i++ {
		if i != start {
			fmt.Print(",")
		}

		fmt.Print(FizzBuzzHandler(i))
	}
	fmt.Println()
}
