package sort

//BubbleSort -- Bubble sort implementation
func BubbleSort(list []int) []int {
	var swapped bool //Used to check if the list is already sorted
	numElements := len(list)

	for numElements > 1 {
		swapped = false
		for i := 1; i < numElements; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				swapped = true
			}
		}
		if !swapped { //No swaps done means that we are done
			break
		}

		//After each round of swapping, the last element is sorted. Thus, we can ignore it
		//on the next round.
		numElements--
	}
	return list
}
