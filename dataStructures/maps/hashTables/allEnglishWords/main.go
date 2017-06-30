package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// found this link on a github repo that converted an infochimps .xls file to a simple
// .txt file. The link to the repo: https://github.com/dwyl/english-words
func main() {
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)
}
