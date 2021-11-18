package main

import (
	"fmt"
	"jmlim-go-study/keyboard"
	"jmlim-go-study/pointer"
	"log"
)

func main() {
	fmt.Println("Hello World!")
	//fmt.Println(type_change.TypeChange(3.75))

	//type_change.TypeChangeExam()

	//score.ScoreInput()
	//	type_change.StringToInt()

	// guess.GuessNumbers()

	// formatting_verb.FormattingVerb()

	/*myInt, myBool, myString := many_returns.ManyReturns()
	fmt.Println(myInt, myBool, myString)

	cans, remainder := many_returns.FloatParts(1.26)
	fmt.Println(cans, remainder)*/

	// paint needed
	/*amount, err := many_returns.PaintNeeded(4.2, -3.0)
	if err != nil {
		// 이 경우에는 에러를 출력하고 프로그램을 종료함
		log.Fatal(err)
	}

	fmt.Printf("%0.2f liters needed\n", amount)*/
	/*
		root, err := many_returns.SquareRoot(-9.3)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%0.3f", root)
		}*/

	/*quotient, err := many_returns.Divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}*/

	/*	pointer.PointerType()
		pointer.PointerValueChange()

		var myFloatPointer *float64 = pointer.CreatePointer()
		fmt.Println(myFloatPointer)
		fmt.Println(*myFloatPointer)

		var myBool bool = true
		pointer.PrintPointer(&myBool)

		amount := 6
		pointer.Double(&amount)
		fmt.Println(amount)*/

	truth := true
	pointer.Negate(&truth)
	fmt.Println(truth)

	lies := false
	pointer.Negate(&lies)
	fmt.Println(lies)

	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
