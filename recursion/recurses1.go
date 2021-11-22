package recursion

import "fmt"

func RecursesCount(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		RecursesCount(start+1, end)
	}
	fmt.Printf("Returning from count(%d, %d) call\n", start, end)
}
