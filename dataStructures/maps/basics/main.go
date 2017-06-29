// _Maps_ are Go's built-in [associative data type]
// (sometimes called _hashes_ or _dicts_ in other languages).

package main

import "fmt"

func main() {

	// variable m = making a map with key type of string and value type of int
	m := make(map[string]int)

	// Set key/value pairs using typical 'name[key] = val' syntax.
	m["key1"] = 7
	m["key2"] = 14

	// Printing a map with e.g. "Println" will show all of
	// its key/value pairs.
	fmt.Println("map: ", m)

	// Get a value for a key with 'name[key]'
	value1 := m["key1"]
	fmt.Println("value1: ", value1)

	// The built-in 'len' returns the number of key/value pairs when called on a map.
	fmt.Println("len: ", len(m))

	// The built-in 'delete' removes key/value pairs from a map.
	delete(m, "key2")
	fmt.Println("map: ", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like '0' or ' "" '. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_ ' _
	// aka "Comma ok idiom" _, ok := m["key"]
	// _, present := m["key2"]
	// or you can get the value
	v, present := m["key2"]
	fmt.Println("present? ", present, v)

	val, pres := m["key1"]
	fmt.Println("present? ", pres, val)

	// You can also declare and initialize a new map in
	// the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
