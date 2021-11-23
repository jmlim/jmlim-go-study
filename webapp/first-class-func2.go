package webapp

import "fmt"

func callFunction(passedFunction func()) {
	passedFunction()
}

func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}

func callWithArguments(passedFunction func(string, bool)) {
	passedFunction("this sentence is", false)
}

func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}

func functionA() {
	fmt.Println("function called")
}

func functionB() string {
	fmt.Println("function called")
	return "Returing from function"
}

func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func FirstClassFunc() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}
