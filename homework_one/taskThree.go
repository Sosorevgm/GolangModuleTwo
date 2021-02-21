package main

import (
	"fmt"
	"os"
)

func taskThree() {

	file, err := os.Create("homework_one/taskThreeFile")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	bytesNumber, writeErr := file.WriteString("Some important string")
	if writeErr != nil {
		fmt.Println(writeErr)
	} else {
		fmt.Printf("Wrote %d bytes in file", bytesNumber)
	}
}
