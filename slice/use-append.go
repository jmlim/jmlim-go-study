package slice

import "fmt"

func UseAppend() {
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))

	slice = append(slice, "c")
	fmt.Println(slice, len(slice))

	slice = append(slice, "d", "e")
	fmt.Println(slice, len(slice))
}

func UseAppend2() {

	s1 := []string{"s1", "s2"}
	s2 := append(s1, "s2", "s3")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")

	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)

	// s3, s4 는 동일한 내부 배열을 공유하고 있기 때문에 s4 슬라이스의 원소 값을 변경하면 s3 슬라이스에도 변경된 값이 반영.
	// 반대로  s2나 s1 은 다른 내부 배열을 사용하고 있기 때문에 아무 영향도 받지 않음.

	// 따라서 append 를 호출할 때에는 반환값을 append 에 전달한 것과 동일한 슬라이스 변수에 재할당하는게 일반적.
	// 이처럼 하나의 슬라이스만 저장하면 서로 다른 두 슬라이스가 동일한 내부 배열을 공유하는 문제를 걱정할 필요가 없음.

	ss1 := []string{"s1", "s1"}
	ss1 = append(ss1, "s2", "s2")
	ss1 = append(ss1, "s3", "s3")
	ss1 = append(ss1, "s4", "s4")
	fmt.Println(ss1)

	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%v\n", slice)
}
