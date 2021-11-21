package interface_playlist

import "fmt"

/**
타입이 포인터 리시버 메서드를 선언하고 있는 경우 인터페이스 변수에는 해당 타입의 포인터 값만 할당할 수 있음.
 - 아래 코드에서 Switch 타입의 toggle 메서드는 리시버 값을 변경해야 하기 때문에 포인터 리시버를 사용해야 한다.
*/
type Switch string

type Toggleable interface {
	toggle()
}

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

func PointerReceiverMain() {
	s := Switch("off")
	// 	var t Toggleable = s // Toggable 인터페이스 타입의 변수에 Switch 변수를 할당하려고 하면 에러가 발생 함.
	var t Toggleable = &s
	t.toggle()
	t.toggle()
	t.toggle()
	t.toggle()
}
