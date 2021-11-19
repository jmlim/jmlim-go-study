package maps

import (
	"fmt"
	"jmlim-go-study/datafile"
	"log"
)

func VotesCountMap() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
		//fmt.Println(line)
		//_, ok := counts[line]
		//fmt.Println(ok)
	}

	//fmt.Println(counts)

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

}
