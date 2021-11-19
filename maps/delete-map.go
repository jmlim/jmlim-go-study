package maps

import "fmt"

func DeleteMap() {

	var ok bool
	ranks := make(map[string]int)

	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze") // bronze 키와 해당 키의 값을 삭제
	rank, ok = ranks["bronze"]
	fmt.Printf("rank :%d, ok: %v\n", rank, ok)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5) // 5키와 해당 키의 값을 삭제
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)

}
