package main

import "fmt"

func taskOne() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred: ", err)
		}
	}()

	var numArr = []int{1, 2, 3}
	fmt.Print(numArr[3])
}
