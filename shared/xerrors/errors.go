package xerrors

import "errors"

type BaseError error

var (
	ErrNotFound BaseError = errors.New("not found")
)
