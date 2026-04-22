package main

import "strconv"

func FizzBuzzHandler(value int) string {

	if value < 1 || value > 100 {
		panic(1)
	}

	switch value % 15 {
	case 0:
		return "FizzBuzz"

	case 3, 6, 9, 12:
		return "Fizz"

	case 5, 10:
		return "Buzz"
	}

	return strconv.Itoa(value)
}
