package main

import "strconv"

func FizzBuzzHandler(value int) string {

	if value%3 == 0 {
		return "Fizz"
	}

	if value%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(value)
}
