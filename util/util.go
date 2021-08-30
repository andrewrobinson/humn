package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/andrewrobinson/humn/model"
)

func GetPostcode(coord model.Coord, apiTokenFlag string, poolSizeFlag int) string {

	url := buildMapboxUrl(coord, apiTokenFlag)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// printBody(body)

	mb := model.MapbookResponse{}
	err = json.Unmarshal(body, &mb)
	if err != nil {
		log.Fatalln(err)
	}

	postcode := strings.Join(mb.Features, "-")

	//TODO - I am getting back an empty 'features' from mapbox

	// From the response https://docs.mapbox.com/api/search/geocoding/#geocoding-response-object
	// the relevant field you should obtain is the `text` field from the single returned Feature.

	// return "SE14 9XB"
	return "TODO - find out why features is []:" + postcode
}

func buildMapboxUrl(coord model.Coord, token string) string {

	urlTemplate := "https://api.mapbox.com/geocoding/v5/mapbox.places/<%f,%f>.json?types=postcode&limit=1&access_token=%s"
	return fmt.Sprintf(urlTemplate, coord.Lat, coord.Lng, token)

}

// func printBody(body []byte) {
// 	sb := string(body)
// 	log.Printf(sb)
// }
