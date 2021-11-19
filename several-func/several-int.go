package several_func

import "fmt"

func SeveralInts(numbers ...int) {
	fmt.Println(numbers)
}

func Mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}
