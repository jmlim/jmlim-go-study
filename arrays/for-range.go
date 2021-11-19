package arrays

import "fmt"

func ForRange() {

	notes := [7]string{
		"do",
		"re",
		"mi",
		"fa",
		"so",
		"la",
		"ti",
	}

	// 배열의 모든 원소 순휘
	for index, note := range notes {
		fmt.Println(index, note)
	}
	fmt.Println("====")

	// index 만 사용
	for index, _ := range notes {
		fmt.Println(index)
	}
	fmt.Println("====")
	// note(value) 만 사용
	for _, note := range notes {
		fmt.Println(note)
	}

}
