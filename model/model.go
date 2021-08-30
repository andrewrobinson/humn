package model

type Coord struct {
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Url      string
	Postcode string `json:"postcode"`
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
	Features []string `json:"features"`
}
