* go는 자동 테스트를 작성할 수 있는 testing 패키지와 테스트를 실행할 수 있는 go test 명령어를 제공

* 같은 패키지 디렉터리에 테스트할파일_test.go 생성.
    * go test 명령어가 이 접미사를 가진 이름의 파일을 검색.

~~~go
package automated_test // 동일한 패키지에 테스트 코드 작성 필요

import "testing" //test 표준 라이브러리 패키지

/**
함수 이름은 Test 로 시작해야 함. Test 뒤로는 아무 이름이나 사용 가능.
 testing.T 값의 포인터를 argument 로 전달
*/
func TestTwoElements(t *testing.T) {
	// 테스트 실패를 위해 testing.T의 메서드 호출
	t.Error("no test written yet")
}

func TestThreeElements(t *testing.T) {
	t.Error("no test here either")
}

~~~
* 컨벤션을 따라야 go test 도구 사용 가능
  * 테스트 파일이 테스트 하려는 코드와 반드시 동일한 패키지에 속할 필요는 없으나 패키지의 노출되지 않은 타입이나 함수에 접근하려면 동일한 패키지에 속해야 함.
  * testing 패키지의 타입을 사용하려면 테스트가 필요하기 때문에 각 테스트 파일의 맨 위에서 해당 패키지를 가져와야 함.
  * 테스트 함수의 이름은 Test로 시작해야 한다.
    * 이외 나머지 부분에는 원하는 이름을 사용할 수 있지만 대문자로 시작하는 것이 좋음.
  * 테스트 함수는 단일 매개변수로 testing.T 값의 포인터를 받음.
  * testing.T값의 (Error와 같은) 메서드를 호출하여 실패한 테스트를 보고할 수 있다.
    * 대부분 메서드는 테스트가 실패한 원인을 설명하는 문자열 메시지를 매개변수로 받는다.


* 테스트 실패 시
~~~shell
adminui-iMac-3:jmlim-go-study jmlim$ go test jmlim-go-study/automated-tt
--- FAIL: TestTwoElements (0.00s)
    join_with_commans_test.go:8: didn't match expected value
FAIL
FAIL    jmlim-go-study/automated-tt     0.006s
FAIL
~~~

* 성공 시
~~~shell
adminui-iMac-3:jmlim-go-study jmlim$ go test jmlim-go-study/automated-tt
ok      jmlim-go-study/automated-tt     0.005s
adminui-iMac-3:jmlim-go-study jmlim$ 
~~~

* go build, go install 과 같은 go 도구는 _test.go 로 끝나는 파일들은 무시함.

* 테스트 부분 집합 실행
~~~shell
adminui-iMac-3:jmlim-go-study jmlim$ go test jmlim-go-study/automated-tt -v
=== RUN   TestOneElements
--- PASS: TestOneElements (0.00s)
=== RUN   TestTwoElements
--- PASS: TestTwoElements (0.00s)
=== RUN   TestThreeElements
--- PASS: TestThreeElements (0.00s)
=== RUN   TestFirstLarger
--- PASS: TestFirstLarger (0.00s)
=== RUN   TestSecondLarger
--- PASS: TestSecondLarger (0.00s)
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestSubtract
--- PASS: TestSubtract (0.00s)
PASS
ok      jmlim-go-study/automated-tt     0.006s
adminui-iMac-3:jmlim-go-study jmlim$ go test jmlim-go-study/automated-tt -v -run Two
=== RUN   TestTwoElements
--- PASS: TestTwoElements (0.00s)
PASS
ok      jmlim-go-study/automated-tt     0.005s
adminui-iMac-3:jmlim-go-study jmlim$ 
adminui-iMac-3:jmlim-go-study jmlim$ go test jmlim-go-study/automated-tt -v -run Elements
=== RUN   TestOneElements
--- PASS: TestOneElements (0.00s)
=== RUN   TestTwoElements
--- PASS: TestTwoElements (0.00s)
=== RUN   TestThreeElements
--- PASS: TestThreeElements (0.00s)
PASS
ok      jmlim-go-study/automated-tt     0.006s
adminui-iMac-3:jmlim-go-study jmlim$ 
~~~