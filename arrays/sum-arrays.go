package arrays

import (
	"fmt"
	"jmlim-go/textfileread"
	"log"
)

func AvgArrays() {
	numbers, err := textfileread.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	/*	numbers := [3]float64{
		71.8,
		56.2,
		89.5,
	}*/

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	// 같은 타입끼리 계산 가능.
	cnt := float64(len(numbers))
	fmt.Printf("Average %0.2f\n", sum/cnt)
}

func Between10and20() {
	numbers := [6]int{3, 16, -2, 10, 23, 12}
	for i, number := range numbers {
		if number >= 10 && number <= 20 {
			fmt.Println(i, number)
		}
	}
}
