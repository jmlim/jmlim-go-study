package arrays

import "fmt"

func ArrayLiteral() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes[0], primes[2], primes[4])

	/**
	배열의 마지막 원소의 경우에도 다음에 줄 바꿈 문자가 오는 경우에는 쉼표를 붙여 주어야 함.
	*/
	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line",
	}

	fmt.Println(text[0], text[1], text[2])

	// 배열 전체 전달도 가능
	fmt.Println(primes)

	// go 코드에서 그대로 보이도록 처리.
	fmt.Printf("%#v", primes)
}
