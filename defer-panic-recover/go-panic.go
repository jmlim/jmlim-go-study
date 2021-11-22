package defer_panic_recover

import (
	"fmt"
	"math/rand"
	"time"
)

/**
직접 패닉 일으키기
*/
func GoPanic() {
	panic("oh, no, we're going down")
}

/**
패닉은 언제 사용하는가?
 - 접근 불가능한 파일, 네트워크 장애 및 잘못된 입력 등과 같은 것은 일반적으로 정상 적인 상태라고 간주하여 error값을 통해 처리하는 것이 좋음.
 - panic은 보통 사용자의 실수나 의도와는 무관한 프로그램 자체의 버그를 나타내는 "제어 불가능한" 상황에서 사용해야 한다.


doorNumber에 잘못된 값이 들어가는 경우? panic을 호출하는게 맞음.
 - 이런 일은 절대 발생해서는 안되며 만에 하나 발생한다면 동작을 잘못하기 전에 프로그램을 중단하는게 좋다.
*/

func awardPrize() {
	doorNumber := rand.Intn(3) + 1

	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number") // 의외의 숫자가 생성되면 패닉 발생 (위 조건에서는 안생기긴 함)
	}
}

func AwardPrizeMain() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}

func Snack() {
	defer fmt.Println("Closing refrigerator")
	fmt.Println("Opening refrigerator")

	panic("refrigerator is empty")
}
