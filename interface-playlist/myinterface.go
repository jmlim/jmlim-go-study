package interface_playlist

import "fmt"

// 인터페이스
type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (m MyType) MethodNotInInterface() { // 인터페이스와 무관한 메서드를 가져도 MyType 에 대한 인터페이스 만족
	fmt.Println("MethodNotInInterface called")
}
