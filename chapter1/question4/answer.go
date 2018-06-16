package question4

/*
Given a string, write a function to check if it is permutation of a palindrome.
A palindrome is a word of phrase that is the same forward and backwards. A
permutation is a rearragnment of letters. The palindrome does not need to be
limited to just dictionary words.

Example:

input - "tact coa"
output - true (permutations: "taco cat", "atco cta", etc)
*/

/*
Assumptions, spaces do not matter. Thus, my input will not have spaces.
*/
func answer(input string) bool {
	//fill table
	table := make(map[string]int)
	for _, char := range input {
		_, ok := table[string(char)]
		if ok {
			table[string(char)]++
		} else {
			table[string(char)] = 1
		}
	}

	oddCount := 0
	for _, value := range table {
		if value%2 == 1 {
			oddCount++
		}

		if oddCount > 1 {
			return false
		}
	}

	return true
}
