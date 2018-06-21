package sort

import (
	"math/rand"
	"time"
)

/*
Partition function notes:

If the item to be sorted is already sorted, then this implementation
will run with O(n^2) time complexity because each call to the partition
function will return the last index, which is len(input)-1). As a result,
quicksort is never able to "divide and conquer" because the "pivot" is always
the last item in the slice.

In order to make this implementation run faster for the case of already having a
sorted list, you can switch the last element with a random element in the list.

See func partition for more details.
- Remove that random swap and run the benchmark tests to see how much slower
  it runs.
*/

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

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
	//Random swap
	randIndex1 := rand.Intn(len(arr))
	arr[randIndex1], arr[len(arr)-1] = arr[len(arr)-1], arr[randIndex1]

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
