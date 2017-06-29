package main

import "fmt"

func main() {
	// shorthand to make an empty slice
	// actually creates slices with empty strings
	student := []string{}
	students := [][]string{}
	fmt.Println("shorthand make empty slice: ", student)
	fmt.Println("shorthand make multiDim empty slice: ", students)
	fmt.Println("Is 'student' nil? ", student == nil)
	fmt.Println("---------")

	// var way of making a slice, creates the default value
	// for slice and map this will be nil because they are
	// reference types
	var person []string
	var people [][]string
	fmt.Println("var make default val slice: ", person)
	fmt.Println("var make default val multiDim slice: ", people)
	fmt.Println("Is 'person' nil? ", person == nil)
	fmt.Println("---------")

	// make a slice using make, actually defining its length
	// NOTE since we don't specify the capacity, it automatically
	// makes the capacity of the underlying array the same as
	// the length of the slice

	// STANDARD WAY TO DO IT, all the golang doc appears to use this technique
	dog := make([]string, 35)
	dogs := make([][]string, 35)
	fmt.Println("make empty slice w defined len: ", dog)
	fmt.Println("make empty slice of slices w defined len: ", dogs)
	fmt.Println("Is 'dog' nil? ", dog == nil)

}
