package main

import "fmt"

func main() {
	var myNums = make(map[int]int)
	fmt.Println(myNums)
	var myGreeting = make(map[string]string)
	fmt.Println(myGreeting)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

	// NOTE: If you just make a map and leg Go set to the default value of nil,
	// then you won't be able to use the map later.
	// So you can either do:
	// a) var myGreeting = make(map[string]string)  OR
	// b) var otherGreeting = map[string]string{}
	// *** Both ways seems to result in an empyt map *** //
	var otherNums = map[int]int{}
	fmt.Println(otherNums)
	var otherGreeting = map[string]string{}
	fmt.Println(otherGreeting)
	otherGreeting["Logan"] = "Hello."
	otherGreeting["Tiffany"] = "Hi there!."

	fmt.Println(otherGreeting)

	otherGreeting["Eisley"] = "Umm..hello?"
	otherGreeting["Amrynn"] = "HIIIIII! My name is Amryn, I just finished gymanstics and I am about to go to the park."

	fmt.Println(otherGreeting)

	// SHORTHAND
	apology := make(map[string]string)
	apology["Tiffany"] = "What do you think about..."
	apology["Logan"] = "I'm sorry, I think I did that because..."
	apology["Eisley"] = "I'm sorry, I will never do that ever again!"
	apology["Amrynn"] = "No I DIDN'T!!!!!"
	fmt.Println(apology)

	// making the map with values in the composite literals
	animalGreetings := map[string]string{
		"dog": "Woof!",
		"cat": "Meow!",
	}
	animalGreetings["fish"] = "blurp"
	fmt.Println(animalGreetings["dog"])
	fmt.Println(animalGreetings)
}
