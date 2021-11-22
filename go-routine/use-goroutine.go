package go_routine

import (
	"fmt"
	"time"
)

func aaaa() {
	for i := 0; i <= 50; i++ {
		fmt.Printf("a")
	}
}

func bbbb() {
	for i := 0; i <= 50; i++ {
		fmt.Printf("b")
	}
}

// 이대로 실행 시 출력값이 보이지 않음.
// 왜?
//  - go 프로그램은 main 고루틴이 종료되면 다른 고루틴이 아직 실행 중이더라도 그 즉시 실행을 중단한다.
// main 함수는 a와 b 함수가 실행될 기회를 얻기도 전에 종료된다.
// 따라서 aaaa 와 bbbb 함수를 실행하는 고루틴이 완료되기 전까지 main 고루틴을 실행중인 상태로 유지해야 한다.
// - 이 작업을 정석대로 하려면... go 의 또 다른 기능인 채널이라는 것을 사용해야 한다.
// 일단 지금은 슬립을 사용하여 UseGoRoutine 함수 종료 시간을 늦춰 결과를 확인해보자
func UseGoRoutine() {
	go aaaa()
	go bbbb()

	time.Sleep(time.Second) //1초간 일시 중지
	fmt.Println("end main()")
}
