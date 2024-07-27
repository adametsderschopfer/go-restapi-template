package repository

import "errors"

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExists   = errors.New("URL exists")
)
