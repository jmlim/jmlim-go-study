package maps

import "fmt"

func MapLiteral() {

	//	myMap := map[string]float64{"a": 1.2, "b": 5.6}
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["bronze"])

	elements := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
	}

	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
}
