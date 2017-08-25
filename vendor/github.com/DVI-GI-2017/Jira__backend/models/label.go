package models

import "errors"

var (
	ErrEmptyLabel = errors.New("label can not be empty!")
)

type Label string

type LabelsList []Label

func (l Label) Validate() error {
	if len(l) == 0 {
		return ErrEmptyLabel
	}
	return nil
}
