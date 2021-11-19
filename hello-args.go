package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// 실행파일.exe 71.8 56.2 89.5
func main() {
	//명령줄 인자 예제
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("%0.2f", sum/sampleCount)
}
