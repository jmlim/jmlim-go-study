package defer_panic_recover

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}

/**
defer 키워드는 "무슨 일이 있어도" 반드시 함수가 호출됨을 보장할 수 있기 때문에 보통 에러가 발생한 경우에도 실행해야 하는 코드에 사용하곤 한다.
 - 흔한 예가 바로 열린 파일을 닫는 일

이 작업은 아주 간단한데 CloseFile 함수의 호출을 (에러 처리 코드를 포함하는) OpenFile 함수 호출 바로 다음으로 옮긴 다음 defer 키워드만 붙여주면 된다.

defer를 사용하면 에러 발생 여부와 무관하게 GetFloats가 종료될 때 CloseFile이 호출됨을 보장할 수 있다
*/
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	// file.Close 대신 CloseFile 을 호출
	// defer 를 추가하면 GetFloats 가 종료될 때까지 실행되지 않는다
	defer CloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err // 이곳에서 에러가 반환되면 CloseFile은 호출되지 않는 문제가 있다.
		}

		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err() // CloseFile은 여기서 에러가 발생하는 경우에도 실행
	}

	return numbers, nil // GetFloats 가 정상적으로 완료되어도 호출됨
}
