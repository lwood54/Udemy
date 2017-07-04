// This is a modification at the mobyDickHashTable lesson. I'm trying to
// get the same book, but different link to work in a similar fashion.
// I wasn't entirely sure why the other code was unique to the other link
// he used.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

	// unmarshaling and using comma, ok idiom to do type assertion
	// to access data
	byt := []byte(body) // conversion to slice of bytes
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Printf("Wind data type: %T \n", dat["wind"])
	fmt.Println(dat["wind"])
	windData := dat["wind"]
	fmt.Printf("windData type: %T \n", windData)
	fmt.Println(dat["wind"])
	wind, ok := dat["wind"].(map[string]interface{}) // used type assertion
	if ok {                                          // after type assertion,
		fmt.Println("wind speed: ", wind["speed"]) // check to make sure the type matches
		fmt.Println("degrees: ", wind["deg"])      // using the comma, ok idiom
		fmt.Println("gusts: ", wind["gust"])
	}
}
