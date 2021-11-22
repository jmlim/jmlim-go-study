package defer_panic_recover

import "fmt"

/**
스택 트레이스
-  호출된 함수는 자기 자신을 호출한 함수로 되돌아 가야 한다.
go는 다른 프로그래밍 언어와 마찬가지로 특정 시점에서 활성화된 함수 호출의 목록을 저장하는 호출 스택(call stack)을 유지한다.

프로그램에서 패닉이 발생하면 스택 트레이스 또는 호출 스택 목록이 패닉 에러 메세지에 포함된다. 스택 트레이스는 크래시의 원인을 파악할 때 유용하게 사용할 수 있다.
*/

func Start() {
	one()
}

func one() {
	two()
}

func two() {
	three()
}

func three() {
	panic("This call stack's too deep for me!")
}

/**
지연된 호출은 크래시가 발생하기 전에 실행된다.
 - 프로그램에서 패닉이 발생해도 모든 지연된(defer) 함수 호출은 계속해서 실행된다.
 - 만약 지연된 호출이 두개 이상이라면 지연된 순서의 역순으로 실행됨
*/
func StartDefer() {
	oneDefer()
}

func oneDefer() {
	defer fmt.Println("Deferred in one()")
	twoDefer()
}

func twoDefer() {
	defer fmt.Println("deferred in two()")
	panic("Let's see what's been deferred!")
}
