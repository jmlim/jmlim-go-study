package slice

import "fmt"

/**
내부 배열을 변경하면 슬라이스도 변경됨.
이러한 잠재적인 문제 때문에 일반적으로 배열을 먼저 만들고 슬라이스 연산자를 사용하는 방법보단
make 나 슬라이스 리터럴을 사용해 슬라이스를 만드는 방법이 더 낫다.
make나 슬라이스 리터럴을 사용하면 내부 배열을 건드릴 일이 전혀 없기 때문...
*/
func SliceChange() {

	array1 := [5]string{"a", "b", "c", "d", "e"}

	slice1 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice1)

}
