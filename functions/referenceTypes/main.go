package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMeSlice(m)
	fmt.Println(m) // [Logan]

	myMap := make(map[string]int)
	changeMeMap(myMap)
	fmt.Println("Logan is: ", myMap["Logan"]) // 35
}

func changeMeSlice(z []string) {
	z[0] = "Logan"
	fmt.Println(z) // [Logan]
}

func changeMeMap(z map[string]int) {
	z["Logan"] = 35
}
