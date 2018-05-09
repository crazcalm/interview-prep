package question1

import (
	"bytes"
)

/*
Is Unique: Implementing an algorithm to determine if a string has all
unique characters.

Part2: What if you cannot use an additional data structure?
*/

/*
TODO: Create a solution that sorts the string in place (Quick sort) and then
	  Checks the neighboring values
*/

/*
TODO: Learn what bit vectors are and create a solution using them.
*/

/*
TODO: Create a solution using the standard library just to compare.
*/

//BruteForce -- brute force solution
func BruteForce(s string) bool {
	for index, currentChar := range s {
		for i := index + 1; i < len(s); i++ {
			if bytes.ContainsRune([]byte{s[i]}, currentChar) {
				return false
			}
		}
	}
	return true
}

//Hash -- Solution that uses an hash
func Hash(s string) bool {
	hash := make(map[rune]int)

	for _, char := range s {
		_, ok := hash[char]
		if !ok {
			hash[char] = 1
		} else {
			return false
		}
	}
	return true
}
