package textfileread

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	// 파일에 끝에 도달해 scanner.Scan 이 false 를 반환할 때 까지 순회
	//i := 0

	for scanner.Scan() {
		/*numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(numbers[i])
		i++*/
		floatnumber, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, floatnumber)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	// 슬라이스 대신 nil 반환..
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
