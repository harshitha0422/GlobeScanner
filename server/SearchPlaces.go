package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

// returns the api url string
func apiGet(method string, query string) string {

	//fmt.Println("inside apiGet")
	var apiKey = "5ae2e3f221c38a28845f05b6ddba4f8e8b04b5c6bbb6c180ea77c286"
	var otmAPI = "https://api.opentripmap.com/0.1/en/places/" +
		method +
		"?apikey=" +
		apiKey
	if query != "" {
		otmAPI += "&" + query
	}

	return otmAPI

}

func searchPlaces(name string) []SearchPlacesResponse {

	// fmt.Println("Inside search Places")
	// fmt.Println("Place" + name)
	//w.Header().Set("Content-Type", "application/json")

	// calling the geoname api
	//eg: https://api.opentripmap.com/0.1/en/places/geoname?apikey=5ae2e3f221c38a28845f05b6ddba4f8e8b04b5c6bbb6c180ea77c286&name=London

	geoApi := apiGet("geoname", "name="+name)
	//fmt.Println(geoApi)
	response, err := http.Get(geoApi)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Print(string(responseData))
	var geoResponse GeoResponse
	json.Unmarshal(responseData, &geoResponse)

	var longitude string
	longitude = strconv.FormatFloat(geoResponse.Lon, 'f', -1, 64)
	var latitude string
	latitude = strconv.FormatFloat(geoResponse.Lat, 'f', -1, 64)

	//calling the radius api
	//eg: https://api.opentripmap.com/0.1/en/places/radius?apikey=5ae2e3f221c38a28845f05b6ddba4f8e8b04b5c6bbb6c180ea77c286&radius=100000&lon=-82.32483&lat=29.65163&format=json

	radiusApi := apiGet("radius", "radius=1000&lon="+longitude+"&lat="+latitude+"&rate=2&format=json")
	//fmt.Println(radiusApi)
	radresponse, err := http.Get(radiusApi)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	radresponseData, err := ioutil.ReadAll(radresponse.Body)
	//print(radresponseData)
	if err != nil {
		log.Fatal(err)
	}

	var radiusResponse []RadiusResponse
	json.Unmarshal(radresponseData, &radiusResponse)
	placesResponseList := make([]SearchPlacesResponse, 0)
	l := len(radiusResponse)
	for i := 0; i < l; i++ {
		item := radiusResponse[i]
		xid := item.Xid
		objectApi := apiGet("xid/"+xid, "")
		//fmt.Println(objectApi)
		objresponse, err := http.Get(objectApi)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		objresponseData, err := ioutil.ReadAll(objresponse.Body)

		if err != nil {
			log.Fatal(err)
		}

		var objectResponse ObjectResponse
		json.Unmarshal(objresponseData, &objectResponse)
		//fmt.Print(objectResponse)
		var spResponse SearchPlacesResponse
		spResponse = SearchPlacesResponse{Xid: objectResponse.Xid, Name: objectResponse.Name, Kinds: objectResponse.Kinds, Prev: objectResponse.Prev, Wiki_ext: objectResponse.Wiki_ext, Inf: objectResponse.Inf}
		placesResponseList = append(placesResponseList, spResponse)

	}

	return placesResponseList

}
