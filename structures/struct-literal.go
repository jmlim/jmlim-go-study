package structures

import (
	"fmt"
	structures_geo "jmlim-go-study/structures-geo"
)

func StructLiteral() {
	subscriber := Subscriber{
		Name:   "Aman Singh",
		Rate:   4.99,
		Active: true,
	}

	fmt.Println("Name: ", subscriber.Name)
	fmt.Println("Rate: ", subscriber.Rate)
	fmt.Println("Active: ", subscriber.Active)

	fmt.Println("====")
	subscriber2 := Subscriber{Rate: 4.99}
	fmt.Println("Name: ", subscriber2.Name)
	fmt.Println("Rate: ", subscriber2.Rate)
	fmt.Println("Active: ", subscriber2.Active)
}

func StructLiteralExam() {
	location := structures_geo.Coordinates{Latitude: 37.42, Longitude: -122.08}
	fmt.Println("Latitude: ", location.Latitude)
	fmt.Println("Longitude: ", location.Longitude)

}
