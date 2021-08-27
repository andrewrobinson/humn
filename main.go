package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// TODO
// cat coordinates.txt | ./your-program "api token" "pool size flag" > output.txt

type Coord struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Postcode string  `json:"postcode"`
}

/*
	{
	"type": "FeatureCollection",
	"query": [
	"51582935",
	"0465231"
	],
	"features": [],
	"attribution": "NOTICE: © 2021 Mapbox and its suppliers. All rights reserved. Use of this data is subject to the Mapbox Terms of Service (https://www.mapbox.com/about/maps/). This response and the information it contains may not be retained. POI(s) provided by Foursquare."
	}
*/

type MapbookResponse struct {
	features []string `json:"features"`
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

	for scanner.Scan() {

		lineNumber++
		line := scanner.Text()

		coord := Coord{}
		err := json.Unmarshal([]byte(line), &coord)
		if err != nil {
			log.Fatalln(err)
		}

		url := getUrl(coord)
		postcode := getPostcode(url)
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

func getUrl(coord Coord) string {

	token := "pk.eyJ1IjoiYW5kcmV3bWNyb2JpbnNvbiIsImEiOiJja3N1bjlubG4wbnRrMnZsc3pwbnVscXJ1In0.9IqlyGEbz7lfcRGcHZdJPQ"

	// https://docs.mapbox.com/api/search/#geocoding

	urlTemplate := "https://api.mapbox.com/geocoding/v5/mapbox.places/<%f,%f>.json?types=postcode&limit=1&access_token=%s"
	return fmt.Sprintf(urlTemplate, coord.Lat, coord.Lng, token)

}

func printBody(body []byte) {
	sb := string(body)
	log.Printf(sb)
}

func getPostcode(url string) string {
	// fmt.Printf("getPostcode with url:%s", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// printBody(body)

	mb := MapbookResponse{}
	err = json.Unmarshal(body, &mb)
	if err != nil {
		log.Fatalln(err)
	}

	postcode := strings.Join(mb.features, "-")

	//TODO - I am getting back an empty 'features' from mapbox

	// From the response https://docs.mapbox.com/api/search/geocoding/#geocoding-response-object
	// the relevant field you should obtain is the `text` field from the single returned Feature.

	// return "SE14 9XB"
	return "TODO - find out why features is []:" + postcode
}
