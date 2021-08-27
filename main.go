package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TODO
// cat coordinates.txt | ./your-program "api token" "pool size flag" > output.txt

type Coord struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Postcode string  `json:"postcode"`
}

func main() {

	file, err := os.Open("coordinates.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0

	// { "lat": <float64>, "lng": <float64>, "postcode": <string> }

	for scanner.Scan() {

		lineNumber++
		line := scanner.Text()

		coord := Coord{}
		json.Unmarshal([]byte(line), &coord)

		url := getUrl(coord)
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

func getUrl(coord Coord) string {

	token := "pk.eyJ1IjoiYW5kcmV3bWNyb2JpbnNvbiIsImEiOiJja3N1bjlubG4wbnRrMnZsc3pwbnVscXJ1In0.9IqlyGEbz7lfcRGcHZdJPQ"

	// https://docs.mapbox.com/api/search/#geocoding

	urlTemplate := "https://api.mapbox.com/geocoding/v5/mapbox.places/<%f,%f>.json?types=postcode&limit=1&access_token=%s"
	return fmt.Sprintf(urlTemplate, coord.Lat, coord.Lng, token)

}

func getPostcode(url string) string {
	// fmt.Printf("getPostcode with url:%s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	return "SE14 9AB"
}
