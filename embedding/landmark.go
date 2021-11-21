package embedding

import (
	"errors"
	"jmlim-go-study/encapsulation"
)

type Landmark struct {
	name string
	encapsulation.Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}
