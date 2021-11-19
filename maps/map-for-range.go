package maps

import "fmt"

func MapForRange() {

	grades := map[string]float64{
		"Alma":  74.2,
		"Rohit": 86.5,
		"Carl":  59.7,
	}

	//  for ..range 루프는 맵을 무작위 순서로 처리., 순서가 중요하지 않은 경우에는 괜찮지만 순서가 필요한 경우에는 추가 작업이 필요
	for name, grade := range grades {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grade)
	}

	fmt.Println("======")

	// key 만 순회하기
	for name := range grades {
		fmt.Println(name)
	}

	fmt.Println("======")
	//  값만 순회하기
	for _, value := range grades {
		fmt.Println(value)
	}
}
