package intersection_of_two_arrays

import (
	"sort"
	"testing"
)

/**
Input
[4,7,9,7,6,7]
[5,0,0,6,1,6,2,2,4]
Output
[5,0,6,1]
Expected
[6,4]
*/

func TestIntersection(t *testing.T) {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	outputArr := intersection(num1, num2)

	sort.Ints(outputArr)
	if !IntArrayEquals(outputArr, []int{2}) {
		t.Error("test failed... answer i solved: ", outputArr, ", ndesired answer : ", []int{2})
	}

	num1 = []int{4, 9, 5}
	num2 = []int{9, 4, 9, 8, 4}
	outputArr = intersection(num1, num2)

	sort.Ints(outputArr)
	if !IntArrayEquals(outputArr, []int{4, 9}) {
		t.Error("test failed... answer i solved: ", outputArr, ", ndesired answer : ", []int{4, 9})
	}

	num1 = []int{4, 7, 9, 7, 6, 7}
	num2 = []int{5, 0, 0, 6, 1, 6, 2, 2, 4}
	outputArr = intersection(num1, num2)

	sort.Ints(outputArr)
	if !IntArrayEquals(outputArr, []int{4, 6}) {
		t.Error("test failed... answer i solved: ", outputArr, ", ndesired answer : ", []int{4, 6})
	}

}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
