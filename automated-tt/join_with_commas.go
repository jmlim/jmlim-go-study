package automated_tt

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ")

	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		result += " and "
	} else {
		result += ", and "
	}
	result += phrases[len(phrases)-1]
	return result
}

func JoinWithCommasMain() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
}
