// 4. What's the value of the expression: (true&&false) || (false&&true) || !(false&&false)?

package main

import "fmt"

func main() {
	expVal := (true && false) || (false && true) || !(false && false)
	fmt.Println(expVal)
}
