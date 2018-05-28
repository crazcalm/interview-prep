package practice

/*
Practice Recursion, Memoization and Goroutines:

Fib: Fib(n) = Fib(n - 1) + Fib(n - 2) where Fib(0) = 0 and Fib(1) = 1

Fib(0) = 0
Fib(1) = 1
Fib(2) = 1
Fib(3) = 2
Fib(4) = 3
Fib(5) = 5
Fib(6) = 8
Fib(7) = 13
Fib(8) = 21
*/

//RecursiveFib -- Typical recursive solution
func RecursiveFib(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return RecursiveFib(num-1) + RecursiveFib(num-2)
}

//IterativeFib -- Iterative solution for fib
func IterativeFib(num int) int {
	if num == 0 {
		return 0
	}
	a := 0
	b := 1
	for i := 2; i < num; i++ {
		c := a + b
		a, b = b, c
	}
	return a + b
}

/*TopDownMemoizationFib -- Fib using top down memoization

Study the recursive tree. Where do you see identical nodes [tree of fib(5)]?

There are lots of identical nodes. For example, fib(3) appears twice and fib(2)
appears three times. Why should we recompute these from scratch each time?

In fact, when we call fib(n), we shouldn't have to do much more than O(n) calls, since
there's only O(n) possible values we can throw at fib. Each time we computer fib, we
should just cache this result and use it later.
*/
func TopDownMemoizationFib(num int) int {
	var cache = make(map[int]int)
	return fibonacciTDM(num, cache)
}

func fibonacciTDM(num int, cache map[int]int) int {
	if num == 0 || num == 1 {
		return num
	}

	answer, ok := cache[num]
	if !ok {
		cache[num] = fibonacciTDM(num-1, cache) + fibonacciTDM(num-2, cache)
		answer = cache[num]
	}
	return answer
}

/*
BottomUpMemoizationFib -- Fib using bottom up memoziotion

We can also take this approach and implement it with bottom-up dynamic programming.
Think about doing the same recursive memoized approach, but in reverse.

First, we compute fib(1) and fib(0), which are already known from the base cases.
Then we use those to compute fib(2). Then we use the prior answers to compute fib(3)
and fib(4)
*/
func BottomUpMemoizationFib(num int) int {
	if num == 0 {
		return num
	}

	if num == 1 {
		return 1
	}

	var cache = make(map[int]int)
	cache[0] = 0
	cache[1] = 1

	for i := 2; i < num; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[num-1] + cache[num-2]
}
