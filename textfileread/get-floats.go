package textfileread

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloats(filename string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	// 파일에 끝에 도달해 scanner.Scan 이 false 를 반환할 때 까지 순회
	i := 0

	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(numbers[i])
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil
}
