package guess

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GuessNumbers() {

	// 관련해서 아래 설명
	rand.Seed(time.Now().UnixNano())

	const MAX_COUNT int = 10
	// 1부터 100 까지
	targetNum := rand.Intn(100) + 1
	for i := 0; i < MAX_COUNT; i++ {
		fmt.Print("Please enter a number : ")

		// 숫자 입력 받기
		inputInt := inputNum()

		if inputInt < targetNum {
			fmt.Println("Oops, Your guess was LOW.", (MAX_COUNT-1)-i)
		} else if inputInt > targetNum {
			fmt.Println("Oops, Your guess was HIGH.", (MAX_COUNT-1)-i)
		} else {
			fmt.Println("Good job! You guessed it!")
			break
		}
		// def
		//fmt.Println(targetNum, inputInt)
	}

	/*	for i := 0; i <= 100; i++ {
		randomInt := rand.Intn(100) + 1
		fmt.Println(randomInt)
	}*/
}

func inputNum() int {
	reader := bufio.NewReader(os.Stdin)
	//사용자가 엔터 키를 누를때까지 내용을 읽음
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	inputInt, err := strconv.Atoi(input)
	//	inputInt, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	return inputInt
}



/***
Go에서 범위에 걸쳐 난수 생성하기
Golang은 표준 라이브러리에서 난수 생성을 기본적으로 지원합니다. 특히 의사 난수 생성기를 구현 하는 math/rand 패키지가 있습니다.
그러나 정의된 두 값 사이에 임의의 정수를 생성하는 기능은 없습니다. 그러나 0과 최대 값 사이에 주어진 임의의 정수를 생성하는 것이 있습니다.
func Intn(n int) int
func Intn(n int) int
Intn은 기본 소스에서 [0,n)의 음이 아닌 의사 난수를 int로 반환합니다. n <= 0이면 패닉이 발생합니다.
이 함수를 효과적으로 사용하여 두 값 사이에 난수를 생성할 수 있습니다.


package main

import (
    "fmt"
    "math/rand"
)

func main() {
    min := 10
    max := 30
    fmt.Println(rand.Intn(max - min) + min)
}

이것은 꽤 잘 작동하지만, 여러 번 실행하려고 하면 출력에 항상 같은 숫자가 표시됩니다.

실제로 공식 문서에서 읽을 수 있습니다.

Float64 및 Int와 같은 최상위 함수는 프로그램이 실행될 때마다 값의 결정적 시퀀스를 생성하는 기본 공유 소스를 사용합니다.
실행마다 다른 동작이 필요한 경우 Seed 기능을 사용하여 기본 소스를 초기화합니다.

적절한 예는 다음이 되어야 합니다.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) <----
    min := 10
    max := 30
    fmt.Println(rand.Intn(max - min + 1) + min)
}

보안에 민감한 작업에 적합한 난수의 경우 crypto/rand 패키지를 대신 사용하십시오.
*/