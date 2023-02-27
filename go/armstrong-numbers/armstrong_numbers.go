package armstrong

func IsNumber(n int) bool {
	var nums []int
	cur := n
	for cur > 0 {
		nums = append(nums, cur%10)
		cur /= 10
	}

	var sum int
	for _, num := range nums {
		sum += powInt(num, len(nums))
	}
	return sum == n
}

func powInt(num int, pow int) int {
	if pow < 0 {
		return -powInt(num, -pow)
	}
	result := 1
	for i := 0; i < pow; i++ {
		result *= num
	}
	return result
}
