package maps

import (
	"fmt"
	"sort"
)

func MapForRangeSort() {
	grades := map[string]float64{
		"Alma":  74.2,
		"Rohit": 86.5,
		"Carl":  59.7,
	}
	var names []string
	// 맵의 모든 키를 저장하는 슬라이스를 만듬.
	for key := range grades {
		names = append(names, key)
	}
	// 슬라이스를 알파벳 순으로 정렬
	sort.Strings(names)
	// 알파벳 순으로 정렬된 이름 목록을 순회
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}
