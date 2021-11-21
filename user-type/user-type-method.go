package user_type

import "fmt"

// 메서드 정의
// 메서드는 이름앞에 리시버 매개변수를 추가로 선언해 줘야 한다는 점이 다르다
// 정의한 메서드를 호출하기 위해서는 메서드를 호출할 값 다음에 도트 연산자를 사용하여 메서드의 이름과 괄호를 붙여주면 됨.
// 메서드를 호출하고 있는 값을 메서드 리시버 라고 한다.

type MyType string

// 고는 self 나 this 대신 리시버 매개변수 사용 . 즉 m 이 디스고 셀프임.
func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

// 함수와 아주 유사한 메서드..
/**
다은 함수와 마찬가지로 메서드에도 추가 매개변수를 정의할 수 있음.
메서드 블록에서는 리시버 매개변수와 함께 추가로 정의한 매개변수에 접근할 수 있음.
메서드를 호출할 때에는 각 매개변수로 인자를 전달해야 함.
*/
func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) WithReturn() int {
	return len(m)
}

func BasicUserTypeMethod() {
	value := MyType("a MyType value")
	value.sayHi()

	anotherValue := MyType("Another value")
	anotherValue.sayHi()

	fmt.Println("=====")

	value.MethodWithParameters(4, true)

	fmt.Println(value.WithReturn())
}

// 얀습문제

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func UserTypeExam() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(5)

	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}

// 포인터 리시버 메소드
func (n Number) Double() {
	n *= 2
}
func (n *Number) PointerReceiverDouble() {
	*n *= 2
}

func PointerReceiver() {
	number := Number(4)
	fmt.Println("Original value of number: ", number)
	number.Double()
	fmt.Println("number after calling Double: ", number)

	number.PointerReceiverDouble()
	fmt.Println("number after calling PointerReceiverDouble: ", number)
}

// ==== 예제니 이렇게 하였지만 코드의 일관성을 위해 특정 타입의 메서드를 정의할 때 가급적이면 값 리시버와
// 포인터 리시버의 혼용은 피사고 두 타입중 하나만 사용하는게 좋다.
func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

//  포인터 리시버 스트링
func PointerReceiverStringEx() {
	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()

	pointer.method()
	pointer.pointerMethod()

}
