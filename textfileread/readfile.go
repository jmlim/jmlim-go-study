package textfileread

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// 파일에 끝에 도달해 scanner.Scan 이 false 를 반환할 때 까지 순회
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
