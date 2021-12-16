package business

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrEmptyForm = errors.New("fill all required form")

	ErrUser = errors.New("username or password wrong")
)
