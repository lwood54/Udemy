package main

import "fmt"

func main() {
	age := 44
	fmt.Println("1st memory address of &age: ", &age) // 0xc82000a300

	changeMe(&age)

	fmt.Println("2nd memory address of &age: ", &age) // 0xc82000a300
	fmt.Println("1st value of age: ", age)            // 24

	age2 := 123
	fmt.Println("1st memory address of &age2: ", &age2) // 0xc82000a3c0
	changeMe(&age2)

	fmt.Println("2nd memory address of &age2: ", &age2) // 0xc82000a3c0
	fmt.Println("1st value of age2: ", age2)            // 24
}

// this func takes a parameter of type pointer to an int
func changeMe(z *int) {
	fmt.Println("1st value of z: ", z) // 0xc82000a300
	fmt.Println("1st *z: ", *z)        // 44
	*z = 24
	fmt.Println("2nd z: ", z)   // 0xc82000a300
	fmt.Println("2nd *z: ", *z) // 24
}
