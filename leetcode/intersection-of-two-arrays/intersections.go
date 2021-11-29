package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	num1Map := make(map[int]bool)
	num2Map := make(map[int]bool)
	var rtn []int

	for _, val := range nums1 {
		num1Map[val] = true
	}

	for _, val := range nums2 {
		_, ok1 := num1Map[val]
		if ok1 {
			_, ok2 := num2Map[val]
			if !ok2 {
				rtn = append(rtn, val)
				num2Map[val] = true
			}
		}
	}

	return rtn
}
