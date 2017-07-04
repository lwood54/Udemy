package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	// this file works with his code
	// file used by instructor: http://www.gutenberg.org/files/2701/old/moby10b.txt

	// not sure why this link wouldn't also work...it throws an out of index error message
	// alternate file: http://www.gutenberg.org/files/2701/2701-0.txt

	// Get and store response from url
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	// if there is an error, print error
	if err != nil {
		log.Fatalln(err)
	}
	// we will close the body at the end of the execution of the file
	defer res.Body.Close()

	// create a variable that stores the *Scanner (pointer to scanner)
	// the bufio.NewScanner takes in type io.Reader, returns *Scanner
	bodyContent := bufio.NewScanner(res.Body)
	fmt.Println("content: ", bodyContent) // prints -->  content: &{0xc4200c0700 0x106a510 65536 [] [] 0 0 <nil> 0 false false}
	// .Split takes in a split fun like ScanLines()
	// ScanLines() takes in takes in a slice of bytes ([]byte) and returns a toekn []byte
	// it says it returns each line of text. (but still not in string format)
	bodyContent.Split(bufio.ScanWords)        // splits bodyContent by each word
	bodyContent.Split(bufio.ScanLines)        // this is the default, but can be used to specify the action as well
	fmt.Println("after split: ", bodyContent) // still not string text. prints --> &{0xc4200be7c0 0x106a510 65536 [] [] 0 0 <nil> 0 false false}
	for i := 1; i <= 50; i++ {
		bodyContent.Scan() // scans the next token, which will be available throu Bytes or Text method
		// .Text() takes in a *Scanner and returns a string
		fmt.Println("Line ", i, ": ", bodyContent.Text()) // this is where the *Scanner is actuall converted to Text
	}

	///// using iouti.ReadAll and converting to a string
	reader := bufio.NewReader(res.Body)
	readerContent, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		fmt.Println("Read error: ", readErr)
	}
	fmt.Println(string(readerContent))

	// You could do ReadLine, but per the documentation:
	// ReadLine is a low-level line-reading primitive. Most callers should use
	// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.

	// NOTE: Scanner is more efficient and allows you to determine the number of
	// lines you read. ReadAll just gets all of it and ReadLine is apparently not
	// as efficient.
}
