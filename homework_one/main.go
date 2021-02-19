package main

import (
	"GolangModuleTwo/homework_one/custom_errors"
	"fmt"
)

func main() {
	taskOne()

	taskTwo := func() {
		customErr := custom_errors.New("Some error occurred")
		fmt.Print(customErr)
	}
	taskTwo()

	taskThree()
}
