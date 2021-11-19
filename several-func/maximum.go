package several_func

import "math"

func Maximum(numbers ...float64) float64 {

	// 아주 작은 값으로 시작
	max := math.Inf(-1)

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
