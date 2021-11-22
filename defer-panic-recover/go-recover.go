package defer_panic_recover

import "fmt"

/**
go에는 패닉 상태에 빠진 프로그램을 복구할 수 있는 recover라는 내장 함수가 존재
*/

// 패직이 발생하지 않은 프로그램에서 recover를 호출하면 nil 을 반환

func FreakOut() {
	panic("oh no")
	recover() // 실행되지 않음. panic 함수를 호출하는 함수와 동일한 함수에서 recover를 호출하는 것은 의미가 없음.
}

func CalmDown() {
	//fmt.Println(recover()) // 패닉 값은 recover에서 반환됨.
	p := recover()
	err, ok := p.(error) // 패닉 값의 타입을 error 타입으로 단언합니다 .
	if ok {
		// error 타입의 값으로 Error 메서드를 호출 할 수 있음.
		fmt.Println(err.Error())
	}
}

func FreakOut2() {
	defer CalmDown() // 복구를 수행하는 함수의 호출을 지연
	err := fmt.Errorf("there's an error")
	panic(err) // 이 코드 다음으로 패닉이 발생하면 지연된 함수 호출이 프로그램을 복구할 것임.
}
