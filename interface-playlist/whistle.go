package interface_playlist

import "fmt"

type Whistle string
type Horn string
type Robot string

type NoiseMaker interface {
	MakeSound()
}

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

func (r Robot) MakeSound() {
	fmt.Println("Beep Beep!")
}

func (r Robot) Walk() { //인터페이스와 관련없는 추가된 로봇타입의 메소드
	fmt.Println("Powering legs")
}

//인터페이스 활용
func play(n NoiseMaker) {
	n.MakeSound()
	// n.Walk() NoiseMaker 에 정의된 메서드가 아니므로 실행불가
}

func NoiseMakerMain() {
	/*var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()*/
	play(Whistle("Toyco Canary")) // 인터페이스를 만족하는 type
	play(Horn("Toyco Canary"))    // 인터페이스를 만족하는 type
	play(Robot("Botco Ambler"))   // 인터페이스를 만족하는 type
}

/**
타입 단언 ( type assertion) 을 사용하면 구체(constructure) 타입의 값을 가져올 수 있음 (다운 캐스팅 같은것)
var noiseMaker NoiseMaker = Robot("Botco Ambler")
var robot Robot = noiseMaker.(Robot)
                  /             \
              인터페이스 값      단언할 값
*/

func TypeAssertExample() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}
