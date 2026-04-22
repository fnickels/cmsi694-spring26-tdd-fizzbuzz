package main

import "testing"

func TestFizzBuzzHandlerPanicsForOutOfRangeValues(t *testing.T) {
	testCases := []struct {
		name  string
		input int
	}{
		{name: "less_than_one", input: 0},
		{name: "greater_than_hundred", input: 101},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatalf("FizzBuzzHandler(%d) did not panic", tc.input)
				}
			}()

			FizzBuzzHandler(tc.input)
		})
	}
}
