package two_sum

func twoSum(nums []int, target int) []int {

	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if i != j {
				num1 := nums[i]
				num2 := nums[j]
				if target == num1+num2 {
					return []int{i, j}
				}
			}
		}
	}

	return nums
}
