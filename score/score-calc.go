package score

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ScoreInput() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
//	input, _ := reader.ReadString('\n')
//	fmt.Println(input)

	input, err := reader.ReadString('\n')
	log.Fatal(err)
	fmt.Println(input)
}
