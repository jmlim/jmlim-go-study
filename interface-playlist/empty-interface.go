package interface_playlist

import "fmt"

// 메서드가 전혀 필요 없는 인터페이스 타입을 선언한다면?
// 모든 타입이 이 인터페이스를 만족할 수 있게 됨.

/**
interface{}  타입을 빈 인터페이스 (empty interface) 라고 하며, 이 타입은 모든 타입의 값을 저장할 수 있음.
빈 인터페이스에는 인터페이스를 만족하기 위해 필요한 메서드가 없기 때문에 모든 값이 만족할 수 있다.
*/

// 빈 인터페이스 타입의 매개변수를 받음
func AcceptAnything(thing interface{}) {
	fmt.Println(thing)

	// 빈 인터페이스 타입의 값에서 메서드를 호출하려면?
	// 먼저 타입 단언으로 구체 타입의 값을 가져와야 함.
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}
}

func Main() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("Toyco Canary"))
}

// 빈 인터페이스 타입의 값으로는 할 수 있는게 그리 많지 않으므로.. 무턱대고 사용하지 않을 것.
