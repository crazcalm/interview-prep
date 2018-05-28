package question5

/*
Recursive Multiply: Write a recursive function to multiply two positive
integers without using the * operator. You can use addition, subtraction,
and bit shifting, but you should minimize the number of those operations
*/

//SimpleRecursion -- simplest answer possible
//O(b)
func SimpleRecursion(a, b uint32) uint32 {
	if b == 1 {
		return a
	}
	return SimpleRecursion(a, b-1) + a
}

func findMin(a, b uint32) (uint32, uint32) {
	if a < b {
		return a, b
	}
	return b, a
}

//SimpleIteration -- simplest iterative answer
func SimpleIteration(a, b uint32) uint32 {
	var result uint32
	a, b = findMin(a, b)

	for i := uint32(0); i < a; i++ {
		result += b
	}
	return result
}

//SlightlyBetter -- O(c) given if a < b, then c = a else c = b
func SlightlyBetter(a, b uint32) uint32 {
	if a < b {
		return SimpleRecursion(b, a)
	}
	return SimpleRecursion(a, b)
}

//Better -- O(b/2)
func Better(a, b uint32) uint32 {
	if b%2 == 0 {
		half := SimpleIteration(a, b/2)
		return half + half
	}
	half := SimpleIteration(a, (b-1)/2)
	return half + half + a
}
