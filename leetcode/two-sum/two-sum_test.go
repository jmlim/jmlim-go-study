package two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	sum := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(sum)
	sum = twoSum([]int{3, 2, 4}, 6)
	fmt.Println(sum)
	sum = twoSum([]int{3, 3}, 6)
	fmt.Println(sum)
}
