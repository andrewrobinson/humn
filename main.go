package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("coordinates.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0

	type coords struct {
		Lat float32
		Lng float32
	}

	for scanner.Scan() {

		lineNumber++
		line := scanner.Text()

		coord := coords{}

		json.Unmarshal([]byte(line), &coord)
		fmt.Printf("coord:%+v\n", coord)

		// if lineNumber == 4 {
		// 	break
		// }

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error on line %v: %v", lineNumber, err)
	}

}
