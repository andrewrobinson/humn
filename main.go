package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/andrewrobinson/humn/model"
	"github.com/andrewrobinson/humn/util"
)

func main() {
	/*

		Usage:

		go build
		cat coordinates.txt | ./humn --apiToken=x --poolSize=100 > output.txt

		my token is: pk.eyJ1IjoiYW5kcmV3bWNyb2JpbnNvbiIsImEiOiJja3N1bjlubG4wbnRrMnZsc3pwbnVscXJ1In0.9IqlyGEbz7lfcRGcHZdJPQ

	*/

	apiTokenFlag := flag.String("apiToken", "", "no default")
	poolSizeFlag := flag.Int("poolSize", 5, "The number of goroutine for the worker pool")

	flag.Parse()

	poolSize := *poolSizeFlag
	apiToken := *apiTokenFlag

	if flag.Lookup("apiToken").Value.String() == "" {
		fmt.Println("--apiToken flag is required")
		os.Exit(1)
	}

	coordJobsFromStdin := util.GetJobsFromStdin(apiToken)

	runJobsConcurrently(coordJobsFromStdin, poolSize, apiToken)

}

func runJobsConcurrently(coordJobsFromStdin []model.Coord, poolSize int, apiToken string) {
	// modified version of https://gobyexample.com/worker-pools

	numJobs := len(coordJobsFromStdin)
	// fmt.Printf("numJobs:%d, poolSize:%d\n", numJobs, poolSize)

	jobs := make(chan model.Coord, numJobs)
	results := make(chan model.Coord, numJobs)

	for w := 1; w <= poolSize; w++ {
		go worker(w, jobs, results, apiToken)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- coordJobsFromStdin[j]
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		coord := <-results
		outputLine, _ := json.Marshal(coord)
		fmt.Printf("%s\n", outputLine)
	}

}

func worker(id int, jobs <-chan model.Coord, results chan<- model.Coord, apiToken string) {
	// modified version of https://gobyexample.com/worker-pools

	for coord := range jobs {
		// fmt.Println("worker", id, "started  job", coord)
		postcode := util.GetPostcode(coord, apiToken)
		// fmt.Printf("postcode looked up:%s", postcode)
		coord.Postcode = postcode
		// fmt.Printf("worker:%d finished job:%+v\n", id, coord)
		results <- coord
	}
}
