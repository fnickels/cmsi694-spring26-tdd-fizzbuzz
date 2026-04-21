package main

import "strconv"

func FizzBuzzHandler(value int) string {

	if value%3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(value)
}
