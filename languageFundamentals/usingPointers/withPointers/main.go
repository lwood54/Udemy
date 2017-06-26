package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func zeroCorrect(y *int) int { // if you take away the return value here and the "int" on this line
	*y = 0    // then it will still change the value of what is stored in that
	return *y // memory location, which is why you can still print that var below
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println("First x: ", x)                  // the value of x is passed as an argument to the func zero
	zeroCorrect(&x)                              // ??
	fmt.Println("After zero with pointers: ", x) // x is still 5
}
