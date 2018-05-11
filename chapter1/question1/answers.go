package question1

import (
	"bytes"
	"strings"

	"github.com/crazcalm/interview-prep/data_structures/sorting"
)

/*
Is Unique: Implementing an algorithm to determine if a string has all
unique characters.

Part2: What if you cannot use an additional data structure?
*/

/*
QuickSort

Note: Strings are immutable in Golang, so there is not way to do an in place sort via Quick Sort
or any other sorting algorithm.
*/

//QuickSort -- Quick Sort Solution
func QuickSort(s string) bool {
	stringSlice := strings.Split(s, "")
	sort.QuickSortString(stringSlice)
	for i := 0; i < len(stringSlice)-1; i++ {
		if strings.EqualFold(stringSlice[i], stringSlice[i+1]) {
			return false
		}
	}
	return true
}

/*
TODO: Learn what bit vectors are and create a solution using them.
*/

/*
StandardLibrary

Note: This is a O(n^2) because the strings.Count does 1 loop over the entire array

The bellow was taken from the Golang source code: strings.go

// countGeneric implements Count.
func countGeneric(s, substr string) int {
        // special case
        if len(substr) == 0 {
                return utf8.RuneCountInString(s) + 1
        }
        n := 0
        for {
                i := Index(s, substr)
                if i == -1 {
                        return n
                }
                n++
                s = s[i+len(substr):]
        }
}

*/

//StandardLibrary -- standard library solution
func StandardLibrary(s string) bool {
	for _, char := range s {
		if strings.Count(s, string(char)) != 1 {
			return false
		}
	}
	return true
}

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
