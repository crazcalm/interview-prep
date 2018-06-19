package question2

import (
	"strings"
)

/*
Word Frequencies: Design a method to find the frequency of occurrences of
any given word in a book. What if we were running this algorithm multiple times?
*/

//TrimWord -- splits up strings into words
func TrimWord(source string) (words []string) {
	noSpaces := strings.TrimSpace(strings.ToLower(source))
	tempt := strings.Split(noSpaces, " ")

	for _, item := range tempt {
		item = strings.TrimSpace(item)
		if len(item) > 0 {
			if strings.EqualFold(string(item[len(item)-1]), ".") || strings.EqualFold(string(item[len(item)-1]), "?") || strings.EqualFold(string(item[len(item)-1]), "!") {
				words = append(words, item[:len(item)-1])
			} else {
				words = append(words, item)
			}
		}
	}
	return words
}

//Frequency -- Returns a hash that maps words to counts
func Frequency(words []string) map[string]int {
	frequency := make(map[string]int)
	for _, word := range words {
		_, ok := frequency[word]
		if ok {
			frequency[word]++
		} else {
			frequency[word] = 1
		}
	}
	return frequency
}

//WordFrequency -- returns the word frequency for the passed in word
func WordFrequency(freq map[string]int, word string) (numerator, denominator int, result float32) {
	var okay bool
	numerator, okay = freq[word]
	if !okay {
		numerator = 0
	}

	for _, value := range freq {
		denominator += value
	}

	result = float32(numerator) / float32(denominator)
	return
}

//FreqMerge -- combines the two frequency maps into one
func FreqMerge(freq1, freq2 map[string]int) map[string]int {
	for key, value := range freq2 {
		_, okay := freq1[key]
		if okay {
			freq1[key] += value
		} else {
			freq1[key] = value
		}
	}
	return freq1
}
