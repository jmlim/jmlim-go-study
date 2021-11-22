package go_routine

import (
	"fmt"
	"time"
)

/***
func greeting() string {
	return "hi"
}

func main() {
	fmt.Println(go greeting()) // 컴파일 에러 발생.
}
*/

/**
매개변수로 채널 사용
*/
func greeting(myChannel chan string) {
	myChannel <- "hi"
}
func BasicChannel() {
	// 새로운 채널 생성
	myChannel := make(chan string)
	// 채널을 새로운 고루틴에서 실행되는 함수로 전달
	go greeting(myChannel)
	// 채널에서 값을 가져옴
	fmt.Println(<-myChannel)
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func BasicChannel2() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)
	// 채널에서 가져온 값을 순서대로 출력
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}

// 고루틴 동기화 관찰
// 다음은 블로킹을 확인할 수 있도록 의도적으로 속도를 늦추고 있는 또 다른 프로그램
// ---- 현재 고루틴을 일정 시간만큼 일시 중지하는 reportNap 함수로 시작. 고루틴은 매초 중지될 때마다 중지되었다는 상태 메세지를 출력한다.

/**
main 고루틴에서 채널을 하나 생성하여 send 함수로 전달.
그리고 reportNap을 호출하여 5초간 일시 중지한 다음 마지막으로 채널에서 두번의 수신 연산을 수행

이 프로그램을 실행하면 첫 2초 동안은 두 고루틴 모두 중지. 2초 후에는 send 고루틴이 깨어나 채널에 값을 전달.
하지만 send 고루틴의 송신 연산은 main 고루틴이 채널에서 값을 받아가기 전 까지 send 고루틴을 블로킹하기 때문에 더이상 아무것도 실행되지 않음.

main 고루틴은 3초 (5-2=3) 간 더 중지되기 때문에 당장은 아무일도 일어나지 않는다. 그리고 3초 후 main 고루틴이 깨어나면 채널에서 값을 받아온다.
그제서야 send 고루틴의 블로킹이 해제되고 두 번째 값을 전달할 수 있게 된다.
*/

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending go routine", 2)
	fmt.Println("*** sending value ***")
	myChannel <- "a" // BasicChannel3 중지된 상태에서 이 전달은 블로킹 됨
	fmt.Println(" *** sending value ***")
	myChannel <- "b"
}

func BasicChannel3() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
