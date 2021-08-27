package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/andrewrobinson/humn/model"
)

// cat coordinates.txt | ./humn "api token" "pool size flag" > output.txt

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()

		coord := model.Coord{}
		err := json.Unmarshal([]byte(line), &coord)
		if err != nil {
			//TODO - stderr
			log.Fatalln(err)
		}

		// postcode := util.GetPostcode(coord)
		postcode := "code commented out"

		coord.Postcode = postcode
		outputLine, _ := json.Marshal(coord)
		//TODO - stdout via a channel
		fmt.Println(string(outputLine))

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
