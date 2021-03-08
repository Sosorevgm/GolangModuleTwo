package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	name  string
	input int
	want  int
}{
	{
		name:  "Case №1",
		input: -5,
		want:  995,
	},
	{
		name:  "Case №2",
		input: 0,
		want:  1000,
	},
	{
		name:  "Case №3",
		input: 100,
		want:  1100,
	},
	{
		name:  "Case №4",
		input: 1000,
		want:  2000,
	},
}

func TestWorkers(t *testing.T) {
	for _, tc := range cases {
		getNumber := Workers(tc.input)
		if getNumber != tc.want {
			t.Errorf("%s excpected: %v, got: %v", tc.name, tc.want, getNumber)
		}
	}
}

func ExampleWorkers() {
	fmt.Printf("%v", Workers(0))
	// Output: 1000
}
