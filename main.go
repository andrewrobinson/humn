package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/andrewrobinson/humn/model"
)

// cat coordinates.txt | ./humn "api token" "pool size flag" > output.txt

func main() {

	/*

		Started 19h40
		21h04 - Now I started finding code on the internet in order to achieve:

		- Stdin should be read by a separate single goroutine and not directly by the worker pool.
		- Stdout should be written to from a separate single goroutine and not directly from the worker pool.

		//1st I wrote what you see in main2

		//Then I used this as a base. It doesn't actually use channels yet....
		// https://play.golang.org/p/OSS71nSpkV
	*/

	rdr := bufio.NewReader(os.Stdin)
	out := os.Stdout

	for {
		switch line, err := rdr.ReadString('\n'); err {

		case nil:

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

			lineWithEnd := fmt.Sprintf("%s\n", outputLine)

			if _, err = out.WriteString(lineWithEnd); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}

		case io.EOF:
			os.Exit(0)

		// Otherwise there's a problem
		default:
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
}

//this is what you see in the "read-stdin" branch, before I pasted
func main2() {

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
