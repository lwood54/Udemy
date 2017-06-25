package main

import "fmt"

func add5(num int) int {
	num += 5
	return num
}

func add5ToValue(num *int) int {
	*num += 5
	return *num
}

func main() {
	// var x int = 10 // same meaning as below
	x := 10
	add5(x)
	fmt.Println("The value of x: ", x) // not 15, instead still 10

	// var y int = add5(x) // same meaning as below
	// assigned value of func to new variable, so y is 15
	y := add5(x)
	fmt.Println("The value of y: ", y)

	// passing the atual memory location to fun in order to change the
	// actual value of x using an *int (int pointer) as a required parameter
	add5ToValue(&x) // passing the actual memory location to func
	fmt.Println("The value of x is now: ", x)
}
