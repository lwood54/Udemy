package main

import (
	"fmt"

	"github.com/lwood54/Udemy/stringutil"
)

func main() {
	original := "Hello from an alternate location."
	reversedMessage := stringutil.Reverse(original)
	fmt.Println("Original Message: ", original)
	fmt.Println("Reversed message: ", reversedMessage)
}
