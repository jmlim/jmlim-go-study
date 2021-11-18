package pointer

import (
	"fmt"
	"reflect"
)

/**
포인터 타입 출력
*/
func PointerType() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
}

func PointerValueChange() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println("포인터 값 자체를 출력", myIntPointer)         // 포인터 값 자체를 출력
	fmt.Println("포인터 주소에 있는 값을 출력합니다", *myIntPointer) // 포인터 주소에 있는 값을 출력합니다

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)  // 포인터 값 자체를 출력
	fmt.Println(*myFloatPointer) // 포인터 주소에 있는 값을 출력

	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)  //  포인터 값 자체를 출력
	fmt.Println(*myBoolPointer) // 포인터 주소에 있는 값 출력

	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
