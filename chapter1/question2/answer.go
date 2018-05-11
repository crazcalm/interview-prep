package question2

import (
	"strings"
)

/*
Check Permutation: Given two strings, write a method to decide if one
                   is a permutation of the other.
*/

/*
TODO: Even though it should be slower than the Hash solution, I should still write a solution that

- Sorts both strings
- Checks that every char for index i matches
- O(n log(n))
*/

//Counter -- A helper function to create the hash char counters
func Counter(s []string) map[string]int {
	counter := make(map[string]int)
	for _, char := range s {
		_, okay := counter[char]
		if !okay {
			counter[char] = 0
		} else {
			counter[char]++
		}
	}
	return counter
}

//Hash -- hash solution
func Hash(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	string1 := strings.Split(s1, "")
	string2 := strings.Split(s2, "")

	hash1 := Counter(string1)
	hash2 := Counter(string2)

	for key, value := range hash1 {
		hash2Value, okay := hash2[key]
		if !okay {
			return false
		}
		if value != hash2Value {
			return false
		}
	}

	return true
}
