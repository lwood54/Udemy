package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0:  "Good morning!",
		01: "Bonjour!",
		02: "Buenos dias!",
		03: "Bongiorno!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
