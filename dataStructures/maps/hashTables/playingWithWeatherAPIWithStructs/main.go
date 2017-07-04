package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// WeatherData struct to collect data from the API call
type WeatherData struct {
	Wind    Wind
	Sys     Sys
	Weather []Weather
	Name    string `json:"name"`
}

//Weather provides basic weather info
type Weather struct {
	Descrip string `json:"description"`
	Icon    string `json:"icon"`
	ID      int    `json:"id"`
}

// Sys includes sunrise, sunset, country, etc.
type Sys struct {
	Country string `json:"country"`
}

// Wind struct to get specific wind characteristics
type Wind struct {
	Speed  float64 `json:"speed"`
	Degree float64 `json:"deg"`
	Gust   float64 `json:"gust"`
}

func main() {
	res, getErr := http.Get("http://api.openweathermap.org/data/2.5/weather?zip=76177,us&appid=31fca86b431449f34e2decc907601056")
	if getErr != nil {
		log.Fatalln("http.Get error: ", getErr)
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalln("Read Error: ", readErr)
	}
	var data WeatherData
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	fmt.Println("Wind gusts: ", data.Wind.Gust)
	fmt.Println("Wind speed: ", data.Wind.Speed)
	fmt.Println("Wind degrees: ", data.Wind.Degree)

	fmt.Println("Country is: ", data.Sys.Country)
	fmt.Println("City is: ", data.Name)
	fmt.Println("General Weather Description: ", data.Weather[0].Descrip)
	fmt.Println("Weather ID: ", data.Weather[0].ID)
	// NOTE: it was tricky here because I had to
	// define Weather as slice of Weather, then
	// access it with index, THEN dot notation
}
