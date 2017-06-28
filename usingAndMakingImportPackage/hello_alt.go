// several examples going on here
// 1: we named the main file something other than main.go, and it obviously runs fine
//    because this is identified as "package main"
// 2: we get to use a function we created in another package because we imported the
//		package "github.com/lwood54/Udemy_golang/stringutil"

package main

import (
	"fmt"

	"github.com/lwood54/Udemy_golang/stringutil"
)

func main() {
	original := "Hello from an alternate location."
	reversedMessage := stringutil.Reverse(original)
	fmt.Println("Original Message: ", original)
	fmt.Println("Reversed message: ", reversedMessage)
}
