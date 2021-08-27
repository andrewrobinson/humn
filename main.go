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

		22h10 - am wondering if chanOwnerChanConsumer or produceReceive() can help me

		22h20 - back to reading https://gobyexample.com/channels



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

func produceReceive() {
	//from Concurrency in Go
	//ch4

	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")

}

func chanOwnerChanConsumer() {
	//from Concurrency in Go
	//ch4 - confinement - 1
	//this format keeps the responsibilities of the 2 roles

	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)

}
