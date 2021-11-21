package embedding

import (
	"errors"
	"jmlim-go-study/encapsulation"
	"unicode/utf8"
)

type Event struct {
	title string
	//임베딩 (Date 의 노출된 메서드는 Event struct 에서 사용 가능)
	encapsulation.Date
}

func (e *Event) SetTitle(title string) error {
	// 제목의 길이가 30자를 넘어가면 에러 반호나
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

func (e *Event) Title() string {
	return e.title
}
