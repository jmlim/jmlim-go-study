package interface_playlist

import (
	"fmt"
	"github.com/headfirstgo/gadget"
)

// Player 인터페이스를 활용해 playList 함수 수정 후 공통 코드로 사용
type Player interface {
	Play(string)
	Stop()
}

func playList(player Player, song []string) {
	for _, song := range song {
		player.Play(song)
	}
	player.Stop()
}

func PlayListMain() {
	var player Player
	player = gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}

	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}

// 인터페이스.
// go 에서 인터페이스는 특정 값이 가지고 있기를 기대하는 메서드의 집합으로 정의.
// 인터페이스는 interface 키워드를 사용하여 정의할 수 있음.
// interface 키워드 다음으로는 메서드가 가지고 있기를 기대하는 매개변수 또는 반환 값과 함께 메서드 이름의 목록이 중괄호 안에 감싸여 따라옴.
/*
type myInterface interface {
	methodWithoutParameters()
	methodWithParameter(f float64)
	methodWithReturnValue() string
}*/

// 타입 단언 사용하기.
func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	//	recorder := player.(gadget.TapeRecorder) // 이대로 코드를 짜면 타입 단언이 되지 않는 타입에 대해선 패닉상태로 넘어간다.

	// 아래와 같이 구체 타입의 값을 저장할 변수와 두 번째 ok 변수에 할당하고 있음.
	// ok (bool)  값을 통해 다른 타입을 가졌을 경우에 대한 대비책을 마련해 런타임 패닉을 방지하도록 해야 한다. (
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func MainTryOut() {
	TryOut(gadget.TapeRecorder{})
	fmt.Println("====")
	TryOut(gadget.TapePlayer{}) // 런타임 패닉 발생 (타입 단언 실패, TapePlayer 에는 Record 가 없으므로.. )
}
