package sort

//SelectionSort -- Implements Selection Sort
func SelectionSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		min := list[i]
		for j := i; j < len(list); j++ {
			if min > list[j] {
				min, list[j] = list[j], min
			}
		}
		list[i] = min
	}
	return list
}
