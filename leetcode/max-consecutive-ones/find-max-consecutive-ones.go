package max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	temp, max := 0, 0
	for _, n := range nums {
		if n == 1 {
			temp++
		}
		if temp > max {
			max = temp
		}
		if n == 0 {
			temp = 0
		}
	}
	return max
}
