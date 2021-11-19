package structures

import "fmt"

// 사용자 정의 타입 구조체

/**
변수와 마찬가지로 타입 또한 함수 내에서 정의 가능. 하지만 이 경우에는 타입의 스코프가 함수 블록으로 한정되기 때문에 함수 외부에선 사용불가.
이러한 이유로 타입은 보통 함수 외부의 패키지 수준에서 선언하는게 일반적임.

백문이 불여일견이라고 곧바로 타입 정의 예제 코드를 보겠다.

변수 선언 시 긴 구조체 정의를 일일이 작성할 필요 없이 사용자 정의 타입의 이름만 사용하면 됨.
*/

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func DefinedTypeStruct() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name : ", porsche.name)
	fmt.Println("Name : ", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description: ", bolts.description)
	fmt.Println("Count: ", bolts.count)
}

func Exam1() {
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func Exam2() {
	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
}

func doublePack(p *part) {
	p.count *= 2
}
