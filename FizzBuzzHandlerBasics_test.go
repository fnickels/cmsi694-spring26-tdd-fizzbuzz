package main

import "testing"

func TestFizzBuzzHandlerFunctionBasics(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "zero", input: 0, expected: "0"},
		{name: "positive", input: 42, expected: "42"},
		{name: "negative", input: -17, expected: "-17"},
		{name: "large", input: 1000000, expected: "1000000"},
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
