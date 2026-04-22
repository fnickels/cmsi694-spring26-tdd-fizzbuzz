package main

import "testing"

func TestFizzBuzzHandlerFunctionFizz(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "one", input: 1, expected: "1"},
		{name: "two", input: 2, expected: "2"},
		{name: "three", input: 3, expected: "Fizz"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FizzBuzzHandler(tc.input)
			if actual != tc.expected {
				t.Fatalf("FizzBuzzHandler(%d) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}
