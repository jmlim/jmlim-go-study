package several_func

func InRange(min float64, max float64, numbers ...float64) []float64 {

	// 범위 내에 있는 인자 값들을 저장할 슬라이스
	var result []float64

	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}

	return result
}
