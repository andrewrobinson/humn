package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/andrewrobinson/humn/model"
	"github.com/andrewrobinson/humn/util"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	// https://gobyexample.com/worker-pools

	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		ret := j + "_x"
		fmt.Printf("worker:%d finished job:%s return was:%s\n", id, j, ret)
		results <- ret
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- strconv.Itoa(j) + "_j"
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}

func main2() {

	/*

		Usage:

		go build
		cat coordinates.txt | ./humn --apiToken=asd --poolSize=4 > output.txt

		token is: pk.eyJ1IjoiYW5kcmV3bWNyb2JpbnNvbiIsImEiOiJja3N1bjlubG4wbnRrMnZsc3pwbnVscXJ1In0.9IqlyGEbz7lfcRGcHZdJPQ


	*/

	apiTokenFlag := flag.String("apiToken", "", "no default")
	poolSizeFlag := flag.Int("poolSize", 5, "The number of goroutine for the worker pool")

	flag.Parse()

	if flag.Lookup("apiToken").Value.String() == "" {
		fmt.Println("--apiToken flag is required")
		os.Exit(1)
	}

	// fmt.Printf("apiTokenFlag:%s, poolSizeFlag:%d\n\n", *apiTokenFlag, *poolSizeFlag)

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

			postcode := util.GetPostcode(coord, *apiTokenFlag, *poolSizeFlag)
			// postcode := "code commented out"

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

// func produceReceive() {
// 	//from Concurrency in Go
// 	//ch4

// 	chanOwner := func() <-chan int {
// 		resultStream := make(chan int, 5)
// 		go func() {
// 			defer close(resultStream)
// 			for i := 0; i <= 5; i++ {
// 				resultStream <- i
// 			}
// 		}()
// 		return resultStream
// 	}

// 	resultStream := chanOwner()
// 	for result := range resultStream {
// 		fmt.Printf("Received: %d\n", result)
// 	}
// 	fmt.Println("Done receiving!")

// }

// func chanOwnerChanConsumer() {
// 	//from Concurrency in Go
// 	//ch4 - confinement - 1
// 	//this format keeps the responsibilities of the 2 roles

// 	chanOwner := func() <-chan int {
// 		results := make(chan int, 5)
// 		go func() {
// 			defer close(results)
// 			for i := 0; i <= 5; i++ {
// 				results <- i
// 			}
// 		}()
// 		return results
// 	}

// 	consumer := func(results <-chan int) {
// 		for result := range results {
// 			fmt.Printf("Received: %d\n", result)
// 		}
// 		fmt.Println("Done receiving!")
// 	}

// 	results := chanOwner()
// 	consumer(results)

// }
