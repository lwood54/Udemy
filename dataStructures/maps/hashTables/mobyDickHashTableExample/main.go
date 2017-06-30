package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	// this file works with his code
	// file used by instructor: http://www.gutenberg.org/files/2701/old/moby10b.txt

	// not sure why this link wouldn't also work...it throws an out of index error message
	// alternate file: http://www.gutenberg.org/files/2701/2701-0.txt

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	// fmt.Println("******************")
	// for i := 28; i <= 126; i++ {
	// 	fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	// }
}

// HashBucket is for making a hash bucket of words.
func HashBucket(word string) int {
	return int(word[0])
}
