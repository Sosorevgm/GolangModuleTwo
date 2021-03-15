package main

import "testing"

type InputStruct struct {
	Name    string
	Age     int
	Address string
	State   float64
}

var cases = []struct {
	name   string
	input  InputStruct
	want   InputStruct
	values map[string]interface{}
}{
	{
		name:  "Case №1",
		input: InputStruct{Name: "Alex", Age: 25, Address: "Pokrovka", State: 199.12},
		want:  InputStruct{Name: "Joe", Age: 78, Address: "White House", State: 100000.58},
		values: map[string]interface{}{
			"Name":    "Joe",
			"Surname": "Biden",
			"Age":     78,
			"Address": "White House",
			"Dog":     "Jack",
			"State":   100000.58,
		},
	},
	{
		name:  "Case №2",
		input: InputStruct{Name: "John", Age: 40, Address: "Pokrovka", State: 199.12},
		want:  InputStruct{Name: "Donald", Age: 74, Address: "Pokrovka", State: 199.12},
		values: map[string]interface{}{
			"Name":    "Donald",
			"Surname": "Tramp",
			"Age":     74,
		},
	},
	{
		name:  "Case №3",
		input: InputStruct{Name: "Victor", Age: 30, Address: "Pokrovka", State: 1110.00},
		want:  InputStruct{Name: "Barak", Age: 59, Address: "Pokrovka", State: 50220.00},
		values: map[string]interface{}{
			"Name":    "Barak",
			"Surname": "Obama",
			"Age":     59,
			"Dog":     "Bob",
			"State":   50220.00,
			"Car":     "BMW",
		},
	},
}

func TestFoo(t *testing.T) {
	for _, tc := range cases {
		inputStruct := tc.input
		err := Foo(&inputStruct, tc.values)
		if err != nil {
			t.Errorf("%v got error: %v", tc.name, err)
		}
		if inputStruct != tc.want {
			t.Errorf("%s excpected: %v, got: %v", tc.name, tc.want, inputStruct)
		}
	}
}
