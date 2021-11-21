package user_type

import "fmt"

type Liters float64
type MilliliLiters float64
type Gallons float64

// Liters 의 메서드
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

// 밀리리터의 메서드
func (m MilliliLiters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() MilliliLiters {
	return MilliliLiters(g * 3785.41)
}

func (l Liters) ToMilliliters() MilliliLiters {
	return MilliliLiters(l * 1000)
}

func (m MilliliLiters) ToLiters() Liters {
	return Liters(m / 1000)
}

func Soda() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())

	water := MilliliLiters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
}

func Milk() {
	milk := Gallons(2)
	fmt.Printf("%0.3f  gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f  gallons equals %0.3f millililiters\n", milk, milk.ToMilliliters())
}

func ChangeLitersOrMilliliters() {
	l := Liters(3)
	fmt.Printf("%0.1f  liters is %0.1f milliliters\n", l, l.ToMilliliters())

	ml := MilliliLiters(3)
	fmt.Printf("%0.1f  millililiters is %0.1f liters\n", ml, ml.ToLiters())
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

// 함수 오버로딩 미 지원
/*func ToGallons(m Milliliters) Gallons  {
	return Gallons(m * 0.000264)
}*/
func UserTypeBasic() {
	var carFuel Gallons
	var busFuel Liters

	carFuel = Gallons(10.0)
	busFuel = Liters(10.0)

	// 기본 타입이 같더라도 할당 불가
	//	carFuel = Liters(10.0) X
	fmt.Println(carFuel, busFuel)

	// 하지만.. 기본타입이 같은 경우에는 타입간 변환 가능.
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
}

type Title string

func UserTypeOperator() {
	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Zodiac"))
	fmt.Println(Title("Alien") > Title("Zodiac"))
	fmt.Println(Title("Alien") + "s")
	// fmt.Println(Title("Jaws 2") - " 2")

	// 기본 타입의 리터럴 값과의 연산도 가능.
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	fmt.Println(Gallons(1.2) == 1.2)
	fmt.Println(Liters(1.2) < 3.4)

	// 아래 코드는 타입이 달라 연산 불가.
	// 이러한 특징은 서로 다른 타입을 혼용할 여지를 미연에 방지하기 위함.
	/*	fmt.Println(Liters(1.2) + Gallons(3.4))
		fmt.Println(Gallons(1.2) == Liters(3.4))*/
}

type Population int

func UserTypePopulation() {
	var population Population
	population = Population(572)
	fmt.Println("Sleepy Creek Country population : ", population)
	fmt.Println("Congratulations, Kevin and Anna ! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek County population: ", population)
}
