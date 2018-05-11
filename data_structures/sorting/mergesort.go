package sort

/*
Source: https://en.wikipedia.org/wiki/Merge_sort
- Top-down implementation using lists
*/

//MergeSort -- Merge sort
func MergeSort(list []int) []int {
	//Base Case: A list of zero or one elements is sorted
	if len(list) <= 1 {
		return list
	}

	//Recusive Case:
	//First, divide the list into equal-sized sublists consisting
	//of the first half and second half of the list
	middlePoint := len(list) / 2
	left := list[:middlePoint]
	right := list[middlePoint:]

	//The recursive Call
	left = MergeSort(left)
	right = MergeSort(right)

	//Then merge the now sorted lists
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	for {
		//Base Case to break loop
		if len(left) == 0 || len(right) == 0 {
			break
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]

		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	//Either left or right may have elements left. Consume them.
	//(Only one of the below cases will be true)
	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
