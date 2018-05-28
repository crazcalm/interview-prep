package question1

/*
Triple Step: A child is running up a staircase with n steps and can
hop either 1 step, 2 steps, or 3 steps at a time. Implement a method
to count how many possible ways the child can run up the steps
*/

//StepsRecursion -- Recursive solution
func StepsRecursion(steps int) int {
	if steps == 0 {
		return 1
	}
	if steps < 0 {
		return 0
	}

	return StepsRecursion(steps-1) + StepsRecursion(steps-2) + StepsRecursion(steps-3)
}

//StepsMemoization -- Recursive solution with memoization
func StepsMemoization(steps int) int {
	var cache = make(map[int]int)
	return stepsMemoization(steps, cache)

}

func stepsMemoization(steps int, cache map[int]int) int {
	if steps == 0 {
		return 1
	}
	if steps < 0 {
		return 0
	}

	answer, ok := cache[steps]
	if !ok {
		cache[steps] = stepsMemoization(steps-1, cache) + stepsMemoization(steps-2, cache) + stepsMemoization(steps-3, cache)
		answer = cache[steps]
	}
	return answer
}
