package contains

import (
	"errors"
	"gomockexample/value"
)

type ValueWrapper struct {
	id    int
	value value.Value
}
type Contains struct {
	values []ValueWrapper
}

var (
	ErrValueNotFound = errors.New("Value Not Found")
)

func (c *Contains) searchId(id int) (*ValueWrapper, error) {
	for _, value := range c.values {
		if value.id == id {
			value.value.Print(true)
			return &value, nil
		}
	}
	return nil, ErrValueNotFound
}
