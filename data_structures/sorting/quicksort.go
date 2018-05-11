package sort

func partitionString(s []string) (primeIdx int) {
	primeIdx = 0
	for i := 0; i < len(s); i++ {
		if s[i] < s[len(s)-1] {
			s[i], s[primeIdx] = s[primeIdx], s[i]
			primeIdx++
		}
	}
	s[primeIdx], s[len(s)-1] = s[len(s)-1], s[primeIdx]
	return primeIdx
}

//Source: https://github.com/shady831213/algorithms/blob/master/sort/quickSort.go
func partition(arr []int) (primeIdx int) {
	primeIdx = 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[len(arr)-1] {
			arr[i], arr[primeIdx] = arr[primeIdx], arr[i]
			primeIdx++
		}
	}
	arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
	return
}

//QuickSortString -- quicksort for strings
func QuickSortString(s []string) {
	if len(s) > 1 {
		primeIdx := partitionString(s)
		QuickSortString(s[:primeIdx])
		QuickSortString(s[primeIdx+1:])
	}
}

//QuickSort -- quicksort
func QuickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := partition(arr)
		QuickSort(arr[:primeIdx])
		QuickSort(arr[primeIdx+1:])
	}
}
