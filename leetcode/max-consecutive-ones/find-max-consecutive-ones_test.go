package max_consecutive_ones

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	inputs := []int{1, 1, 0, 1, 1, 1}
	output := findMaxConsecutiveOnes(inputs)
	if output != 3 {
		t.Error("test failed... answer i solved: ", output, ", ndesired answer : ", 3)
	}

	inputs2 := []int{1, 0, 1, 1, 0, 1}
	output2 := findMaxConsecutiveOnes(inputs2)
	if output2 != 2 {
		t.Error("test failed... answer i solved: ", output2, ", ndesired answer : ", 2)
	}

	inputs3 := []int{0}
	output3 := findMaxConsecutiveOnes(inputs3)
	if output3 != 0 {
		t.Error("test failed... answer i solved: ", inputs3, ", ndesired answer : ", 0)
	}

	inputs4 := []int{1}
	output4 := findMaxConsecutiveOnes(inputs4)
	if output4 != 1 {
		t.Error("test failed... answer i solved: ", inputs4, ", ndesired answer : ", 1)
	}
}
