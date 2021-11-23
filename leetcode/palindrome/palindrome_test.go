package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	palindrome := isPalindrome(121)
	if palindrome != true {
		t.Error("test failed... answer i solved: ", palindrome, ", ndesired answer : ", true)
	}

	palindrome = isPalindrome(-121)
	if palindrome != false {
		t.Error("test failed... answer i solved: ", palindrome, ", ndesired answer : ", false)
	}

	palindrome = isPalindrome(10)
	if palindrome != false {
		t.Error("test failed... answer i solved: ", palindrome, ", ndesired answer : ", false)
	}

	palindrome = isPalindrome(-101)
	if palindrome != false {
		t.Error("test failed... answer i solved: ", palindrome, ", ndesired answer : ", false)
	}

}
