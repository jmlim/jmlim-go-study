package defer_panic_recover

import "fmt"

/*func Socialize()  {
	defer fmt.Println("Goodbye!") //함수 호출 지연시키기
	fmt.Println("hello!")
	fmt.Println("Nice weather, eh?")
}*/

func Socialize() error {
	defer fmt.Println("Goodbye!") //함수 호출 지연시키기
	fmt.Println("hello!")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Nice weather, eh?")
	return nil
}
