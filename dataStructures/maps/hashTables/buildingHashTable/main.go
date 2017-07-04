package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 12) // changed to 12 from 200 for Version 2
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
	// fmt.Println("**************")
	// for i := 28; i <= 126; i++ {
	//  fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	// }
}

// HashBucket helps create buckets to categorize the words in Moby Dick
// this allows for quicker access by creating hashtables.

// HashBucket VERSION 3... more efficient and balanced
func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}

// HashBucket VERSION 2... dividing into 12 buckets
// func HashBucket(word string, buckets int) int {
// 	letter := int(word[0])
// 	bucket := letter % buckets
// 	return bucket
// }

// HashBucket VERSION 1...least balanced
// func HashBucket(word string) int {
// 	return int(word[0])
// }
