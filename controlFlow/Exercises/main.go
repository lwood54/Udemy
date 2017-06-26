package main

import "fmt"

func main() {

	// Number 1
	fmt.Println("E - 1:")
	fmt.Println("Create a program which prints 'Hello World' to the terminal.")
	fmt.Println("Hello world.")

	// Number 2
	fmt.Println("E - 2:")
	fmt.Println("Modify the previous program so that instead of printing Hello World it prints Hello, my name is followed by your name.")
	var myName string
	var yourName string
	fmt.Print("My name is:")
	fmt.Scan(&myName)
	fmt.Print("What is your name?")
	fmt.Scan(&yourName)
	fmt.Println("Hello, ", myName, " is followed by ", yourName)

	// Number 3, blocked from view.

	// Number 4
	fmt.Println("E - 4:")
	fmt.Println("Create a program that prints to the terminal asking the user for 2 numbers.Print the remainder of the larger number divided by the smaller number.")
	var num1 int
	var num2 int
	var remainder int
	fmt.Print("Pick a number...")
	fmt.Scan(&num1)
	fmt.Print("Pick another number...")
	fmt.Scan(&num2)
	if num1 >= num2 {
		remainder = num1 % num2
		fmt.Println("The remainder of ", num1, " / ", num2, " is: ", remainder)
	} else {
		remainder = num2 % num1
		fmt.Println("The remainder of ", num2, " / ", num1, " is: ", remainder)
	}

	// Number 5
	fmt.Println("E - 5: ")
	fmt.Println("Print all of the even numbers between 0 and 100.")
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Number 6
	fmt.Println("E - 6: ")
	fmt.Println(
		"Write a program that prints the numbers from 1 to 100.\n",
		"But for multiples of three print 'Fizz' instead of the number and\n",
		"for the multiples of five print 'Buzz'. For numbers which are multiples\n",
		"of both three and five print 'FizzBuzz'.", // If I put paren here, then I don't need comma.
	)
	var fizzNum1 int
	var fizzNum2 int
	fmt.Print("Pick a starting number to fizz buzz...")
	fmt.Scan(&fizzNum1)
	fmt.Print("Pick an ending number to fizz buzz...")
	fmt.Scan(&fizzNum2)
	fizzBuzzIt(fizzNum1, fizzNum2)

	// Number 7
	fmt.Println("E - 7: ")
	fmt.Println(
		"If we list all the natural numbers below 10 that are multiples of 3 or 5,\n",
		"we get 3, 5, 6 and 9. The sum of these multiples is 23.\n",
		"Find the sum of all the multiples of 3 or 5 below 1000.")
	var sumNum1 int
	var sumNum2 int
	fmt.Print("Pick a number to start from: ")
	fmt.Scan(&sumNum1)
	fmt.Print("Pick a number to stop before: ")
	fmt.Scan(&sumNum2)
	var sum int = sumNums3_5(sumNum1, sumNum2)
	fmt.Println("The sum of all numbers divisible by 5 or 3\n",
		"starting from", sumNum1, " and stopping before ", sumNum2, " = ", sum)
}

func fizzBuzzIt(firstNum int, lastNum int) { // I could also have this return a string with the list.
	for i := firstNum; i <= lastNum; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("--FizzBuzz--")
		} else if i%3 == 0 {
			fmt.Println("--Fizz--")
		} else if i%5 == 0 {
			fmt.Println("--Buzz--")
		} else {
			fmt.Println(i)
		}
	}
}

func sumNums3_5(firstNum int, lastNum int) int {
	total := 0
	for i := firstNum; i < lastNum; i++ {
		if i%5 == 0 || i%3 == 0 {
			total += i
		}
	}
	return total
}
