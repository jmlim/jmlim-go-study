package interface_playlist

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

// 인터페이스
type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}

func CarTruckExample() {
	var vehicle Vehicle = Car("Toyoda Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()

	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("test cargo")
	}
}

func TryTruck() {
	TryVehicle(Truck("Fnord F180"))
}
