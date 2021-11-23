package palindrome

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	strnum := strconv.Itoa(x)
	fmt.Println(strnum, reverseString(strnum))
	return strnum == reverseString(strnum)
}

func reverseString(s string) string {
	chars := []rune(s)
	length := len(s)

	for i := 0; i < length/2; i++ {
		ch := chars[i]
		chars[i] = chars[length-i-1]
		chars[length-i-1] = ch
	}

	return string(chars)
}
