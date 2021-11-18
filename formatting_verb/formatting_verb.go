package formatting_verb

import "fmt"

func FormattingVerb() {

	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
	fmt.Printf("That will be $%f please\n", 0.23*5)

	fmt.Printf("A float : %f\n", 3.1415)
	fmt.Printf("An Integer : %d\n", 15)
	fmt.Printf("A string : %s\n", "hello")
	fmt.Printf("A boolean : %t\n", false)
	fmt.Printf("Values : %v %v %v\n", 1.2, "\t", true)
	
	// %#v 를 사용하면 일반적인 출력값에서 안보이는 값들도 확인할 수 있음
	fmt.Printf("Values : %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types : %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign : %%\n")
}
