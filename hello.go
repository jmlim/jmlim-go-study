package main

import (
	"fmt"
	"jmlim-go-study/structures"
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

	/*truth := true
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
	fmt.Printf("%0.2f degrees Celsius\n", celsius)*/

	//	slice.MakeSlice()
	//	slice.SliceLiteral()
	// slice.SliceExam()

	// slice.SliceChange()
	// slice.UseAppend()
	//	slice.UseAppend2()

	/*numbers, err := textfileread.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)*/

	// 	slice.SliceExam2()

	//	fmt.Println(os.Args[1:])
	/*
		fmt.Println(several_func.Maximum(71.8, 56.2, 89.5))
		fmt.Println(several_func.Maximum(98.7, 89.7, 98.5, 92.3))

		fmt.Println(several_func.InRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
		fmt.Println(several_func.InRange(-10, 10, 4.1, 12, -12, -5.2))
	*/

	/*intSlice := []int{1, 2, 3}
	several_func.SeveralInts(intSlice...)

	stringSlice := []string{"a", "b", "c", "d"}
	several_func.Mix(1, true, stringSlice...)*/

	/*lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)*/

	//	slice.VotesCount()

	/*	maps.SampleMap()
		fmt.Println("=======")
		maps.MapLiteral()*/

	//	maps.ExampleMap()

	// 	maps.ZeroValues()

	//maps.Status("Alma")
	//maps.Status("Jmlim") // jmlim 점수는 기록되지 않음...

	// maps.DeleteMap()

	// maps.VotesCountMap()

	//	maps.MapForRange()

	// maps.MapForRangeSort()
	// maps.VotesCountMap()

	// structures.BasicStruct()
	// 	structures.DefinedTypeStruct()

	/*var bolts structures.Part
	bolts.Description = "Hex bolts"
	bolts.Count = 24
	structures.ShowInfo(bolts)

	p := structures.MinimumOrder("Hex Bolts")
	fmt.Println(p.Description, p.Count)*/

	/*	subscriber1 := structures.DefaultSubscriber("Aman Singh")
		subscriber1.Rate = 4.99
		structures.PrintInfo(subscriber1)
		fmt.Println("====")
		subscriber2 := structures.DefaultSubscriber("Beth Ryan")
		structures.PrintInfo(subscriber2)*/

	/*var s structures.Subscriber
	structures.ApplyDiscount(&s)
	fmt.Println(s)

	// 포인터를 통한 구조체 필드 접근
	var value structures.Subscriber
	var pointer *structures.Subscriber = &value
	pointer.Name = "jmlim"
	fmt.Println(pointer.Name)
	(*pointer).Active = true
	fmt.Println(pointer)*/

	/*subscriber1 := structures_pointer.DefaultSubscriber("Aman Singh")
	structures_pointer.ApplyDiscount(subscriber1)
	structures_pointer.PrintInfo(subscriber1)

	subscriber2 := structures_pointer.DefaultSubscriber("Beth Ryan")
	structures_pointer.PrintInfo(subscriber2)*/

	//	structures.Exam1()
	structures.Exam2()
}
