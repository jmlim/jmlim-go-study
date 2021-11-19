package maps

import "fmt"

func ZeroValues() {

	// zero value : 0
	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I'haven't been assigned"])

	// zero value : ""
	words := make(map[string]string)
	words["I've been assigned"] = "hi"
	fmt.Printf("%#v\n", words["I've been assigned"])
	fmt.Printf("%#v\n", words["I'haven't been assigned"])

	// zero value : 0
	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++

	fmt.Println(counters["a"], counters["b"], counters["c"])

	// map veriable zero value: nil
	var nilMap map[int]string
	fmt.Printf("%#v\n", nilMap)
	// nilMap[3] = "three" // 맵이 nil 이기 때문에 추가할 수 없음.

	// 키/값을 추가하려면 먼저 make 나 맵 리터럴을 사용해 맵을 생성한 다음 맵 변수에 할당해 줘야 한다.
	myMap := make(map[int]string)
	myMap[3] = "three"
	fmt.Printf("%#v\n", myMap)
}

// 할당된 값과 제로값 구분할 수 있도록 변경
func Status(name string) {

	// 제로값 구분 못함
	/*	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
		grade := grades[name]
		if grade < 60 {
			fmt.Printf("%s is failing!\n", name)
		}*/

	//제로값 구분 가능
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	// 이와 같 상황을 처리하기 위해 맵 키에 접근할 때 선택적으로 반환되는 부울값을 갖는 두 번째 반환값 사용... (ok)
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s. \n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing\n", name)
	}
}
