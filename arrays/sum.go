package arrays

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numsToSum ...[]int) []int {

	sums := make([]int, len(numsToSum))

	// Could use here builtin append function
	// but it is more efficient because uses less space.
	for i, nums := range numsToSum {
		sums[i] = Sum(nums)
	}

	return sums
}
