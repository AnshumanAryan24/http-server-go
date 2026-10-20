package constants

import (
	"errors"
)

type ListenMode interface {
	Desc() string
}

type listenMode struct {
	code int
	desc string
}

const (
	ACCEPT_ALL int = iota
)

func NewListenMode(code int) (ListenMode, error) {
	if code == ACCEPT_ALL {
		return &listenMode{code, "ACCEPT ALL"}, nil
	}
	return nil, errors.New("Invalid code for ListenMode")
}

func (l *listenMode) Desc() string {
	return l.desc
}
