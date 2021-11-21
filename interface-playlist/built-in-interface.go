package interface_playlist

import "fmt"

// error 인터페이스 사용
/**
다음과 같이 정의되어 있음. 나만의 에러 타입을 정의할 수 있으며 그렇게 정의한 타입은 error 타입의 값이 필요한 모든 곳에서 사용가능.

type error interface {
	Error() string
}
*/

// 어떤 장비의 과열 여부를 모니터링 하는 프로그램을 작성한다고 가정.
// OverheatError 타입을 만들어서 사용해보자.

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func CheckTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

// Stringer 인터페이스 사용
/**
type Stringer interface {
	String() string
}
*/
type Gallons float64
type Liters float64
type Milliliters float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func CheckStringer() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Milliliters(12.09248342))
}
