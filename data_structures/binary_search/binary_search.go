package search

//BinarySearch -- Basic binary search implementation
func BinarySearch(list []int, value int) int {
	left := 0
	right := len(list) - 1

	for left < right {
		middle := (left + right) / 2
		if list[middle] == value {
			return middle
		}

		if list[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
