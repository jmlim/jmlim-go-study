package webapp

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func MainSayHi() {
	var myFunction func()
	myFunction = sayHi

	myFunction()
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func MainTwice() {
	twice(sayHi)
	twice(sayBye)
}

/*
func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}*/

func DivideMain() {
	var greeterFunction func()
	var mathFunction func(int, int) float64
	greeterFunction = sayHi
	mathFunction = divide
	greeterFunction()
	fmt.Println(mathFunction(5, 2))
}

/**
타입으로서의 함수.
함수를 매개변수로 받는 함수는 전달받을 함수의 매개변수 및 반환 값의 타입도 지정해줘야 한다.
*/

func doMath(passedFunction func(int, int) float64) {
	result := passedFunction(10, 2)
	fmt.Println(result)
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func multiply(a int, b int) float64 {
	return float64(a * b)
}

func MainDoMath() {
	doMath(divide)
	doMath(multiply)
}
