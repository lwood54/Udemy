package main

import "fmt"

const metersToYards float64 = 1.09361
const yardsToFBField = .01

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) // without the "&", the scanner just returns default value of "0" here
	yards := meters * metersToYards
	fbFields := yards * yardsToFBField
	fmt.Println(meters, " meters is ", yards, " yards, and ", fbFields, " football fields long.")
}
