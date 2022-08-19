package helpers

import "errors"

var (
	ErrorEmptyValue    = errors.New("empty value")
	ErrorNotValidUrl   = errors.New("not a valid url")
	ErrorNotTwitterUrl = errors.New("not a twitter url")
	ErrorPath          = errors.New("not a valid video url")
)
