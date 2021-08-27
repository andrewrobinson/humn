package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// TODO
// cat coordinates.txt | ./your-program "api token" "pool size flag" > output.txt

func getPostcode(url string) string {
	return "SE14 9AB"
}

func main() {

	token := "pk.eyJ1IjoiYW5kcmV3bWNyb2JpbnNvbiIsImEiOiJja3N1bjlubG4wbnRrMnZsc3pwbnVscXJ1In0.9IqlyGEbz7lfcRGcHZdJPQ"

	urlTemplate := "https://api.mapbox.com/geocoding/v5/mapbox.places/<%f,%f>.json?types=postcode&limit=1&access_token=%s"

	file, err := os.Open("coordinates.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0

	// { "lat": <float64>, "lng": <float64>, "postcode": <string> }
	type coords struct {
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
		Postcode string  `json:"postcode"`
	}

	for scanner.Scan() {

		lineNumber++
		line := scanner.Text()

		coord := coords{}

		json.Unmarshal([]byte(line), &coord)

		url := fmt.Sprintf(urlTemplate, coord.Lat, coord.Lng, token)

		postcode := getPostcode(url)
		coord.Postcode = postcode

		// fmt.Printf("coord:%+v\n", coord)

		coordStr, _ := json.Marshal(coord)
		fmt.Println(string(coordStr))

		// if lineNumber == 4 {
		// 	break
		// }

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error on line %v: %v", lineNumber, err)
	}

}
