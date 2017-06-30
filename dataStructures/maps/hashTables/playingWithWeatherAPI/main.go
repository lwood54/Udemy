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

// need to make structs in order to read the JSON data, haven't learned structs yet

// Wind to read wind JSON
type Wind struct {
	Speed string `json:"speed"` // finish this later!!!
}

func main() {
	// get the book moby dick
	// this file works with his code
	// file used by instructor: http://www.gutenberg.org/files/2701/old/moby10b.txt

	// not sure why this link wouldn't also work...it throws an out of index error message
	// alternate file: http://www.gutenberg.org/files/2701/2701-0.txt

	// res, err := http.Get("http://samples.openweathermap.org/data/2.5/weather?zip=76177,us&appid=31fca86b431449f34e2decc907601056")
	// if err != nil {
	// 	log.Fatalln("This is a fatal error: ", err)
	// }
	// body := res.Body

	res, getErr := http.Get("http://api.openweathermap.org/data/2.5/weather?zip=76177,us&appid=31fca86b431449f34e2decc907601056")
	if getErr != nil {
		log.Fatalln("http.Get error: ", getErr)
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatalln("Read Error: ", readErr)
	}

	byt := []byte(body)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	// fmt.Println(dat)
	fmt.Printf("Wind data type: %T \n", dat["wind"])
	fmt.Println(dat["wind"])
	windData := dat["wind"]
	fmt.Printf("windData type: %T", windData)
	// works well to scan a page and print each line!
	// scanner := bufio.NewScanner(res.Body)
	// defer res.Body.Close()
	// scanner.Split(bufio.ScanLines)
	// i := 0
	// for scanner.Scan() {
	// 	if i < 2000 { // stops after the first 2,000 lines
	// 		fmt.Println(scanner.Text())
	// 		i++
	// 	}
	// }

}
