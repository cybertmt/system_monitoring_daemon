package storage

import "errors"

var (
	ErrObjectAlreadyExists = errors.New("stats object already exists")
	ErrObjectDoesNotExist  = errors.New("stats object does not exist")
)
