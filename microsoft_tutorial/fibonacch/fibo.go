package fibonacch

func CalcFibo(num int) []int {
	if num < 2 {
		return make([]int, 0)
	}
	nums := make([]int, num)

	nums[0], nums[1] = 1, 1

	for i := 2; i < num; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}
