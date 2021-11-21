package encapsulation

import (
	"errors"
)

// Date 구조체 필드에 잘못된 값을 넣지 않도록 처리
type Date struct {
	year  int
	month int
	day   int
}

// 설정자 메서드
// 포인터 리시버를 사용해야 함.
// 리시버 매개변수도 일반 매개변수와 마찬가지로 원래 값의 복사본을 사용하기 때문
func (d *Date) SetYear(year int) error {

	// 유효성 체크
	if year < 1 {
		return errors.New("invalid year")
	}

	d.year = year // 복사본이 아닌 원래값을 변경
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

// 접근자 메서드 (getter)
/**
GetName, GetCity 등과 같은 형태는 ?
go 에서는 권장하지 않음. 고 커뮤니티는 접근자 메서드명에 Get 접두사를 붙이는 컨벤션을 버리기로 결정.
Set 은 붙이는데 설정자 메서드와 접근자 메서드를 구분하기 위함.
*/
func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}
