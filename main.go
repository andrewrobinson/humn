package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/andrewrobinson/humn/model"
	"github.com/andrewrobinson/humn/util"
)

// TODO
// cat coordinates.txt | ./your-program "api token" "pool size flag" > output.txt

func main() {

	file, err := os.Open("coordinates.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0

	for scanner.Scan() {

		lineNumber++
		line := scanner.Text()

		coord := model.Coord{}
		err := json.Unmarshal([]byte(line), &coord)
		if err != nil {
			log.Fatalln(err)
		}

		postcode := util.GetPostcode(coord)
		coord.Postcode = postcode
		coordStr, _ := json.Marshal(coord)
		fmt.Println(string(coordStr))

		if lineNumber == 4 {
			break
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error on line %v: %v", lineNumber, err)
	}

}
