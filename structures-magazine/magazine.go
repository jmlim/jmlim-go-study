package structures_magazine

/**
외부 구조체의 익명 필드로 선언된 내부 구조체를 외부 구조체 안에 임베딩(embedded) 되었다고 합니다.
임베딩된 구조체의 필드는 외부 구조체로 승격(promoted) 되는데 승격되었다는 말은 내부 구조체의 필드를 마치 외부 구조체에 속해 있는 것처럼 접근할 수 있음을 의미합니다
*/

/**
Address 구조체 타입은 Subscriber와 Employee 구조체 타입에 임베딩 되었기 때문에
subscriber.Address.City 대신 subscriber.City 만으로 City 필드의 값을 가지고 올 수 있음.
*/

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address // 익명 필드
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
