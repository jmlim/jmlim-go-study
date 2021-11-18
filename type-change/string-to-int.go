package type_change

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func StringToInt() {
	var status string
	fmt.Printf("Enter a grade:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	// 문자열을 float64 값으로 변환
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	// input 의 문자열이 아닌 grade 과 비교
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
