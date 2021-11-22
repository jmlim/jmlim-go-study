package go_routine

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ExampleCom() {

	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close() // ExampleCom 함수 종료 시 네트워크 연결 해제
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 응답 데이터를 문자열로 변환한 뒤 출력
	fmt.Println(string(body))

}

/*
func ResponseSize(url string) {
	fmt.Println("Getting", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body)) //byte 슬라이스의 크기는 페이지의 크기와 같습니다
}*/

/**
go 에서는 동시에 실행되는 작업을 고루틴이라고 부른다.
이는 다른 프로그래밍 언어에 있는 쓰레드와 유사한 개념이지만 고루틴은 쓰레드보다 좀 더 적은 메모리를 사용하여 좀 더 빠른 시작 및 종료시간을 가졌기 때문에 한 번에 더 많은 고루틴을 실행할 수 있다.

고루틴은 사용하기도 숩다. 또 다른 고루틴을 시작하기 위해서는 함수 메서드 호출 앞에 go 키워드만 붙여 주면 된다.

모든 go 프로그램의 main 함수는 고루틴을 사용하여 실행되기 때문에 모든 go 프로그램은 최소 하나의 고루틴 위에서 실행됨


 go 문은 반환값과 함께 사용할 수 없음.. (참고)
*/

/*func ResponseSize(url string, channel chan int) {
	fmt.Println("Getting", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 채널에 길이값 반환
	channel <- len(body) //byte 슬라이스의 크기는 페이지의 크기와 같습니다
}*/

type Page struct {
	URL  string
	Size int
}

func ResponseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 채널에 길이값 반환
	channel <- Page{URL: url, Size: len(body)}
}
