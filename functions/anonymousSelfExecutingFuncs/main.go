// this is an anonymous self executing function
package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm anonymous...and I executed myself...wait what?")
	}() // parentheses tell the function to perform its action
}
