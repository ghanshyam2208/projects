package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound          = errors.New("could not fund any value for provided key")
	ErrWordAlreadyExists = errors.New("word already exists")
)

func Search(dict Dictionary, searchStr string) (string, error) {
	stringValue, isFound := dict[searchStr]
	if !isFound {
		return "", ErrNotFound
	}
	return stringValue, nil
}

func (dict Dictionary) Add(key string, value string) error {
	_, isFound := dict[key]
	if !isFound {
		dict[key] = value
		return nil
	}
	return ErrWordAlreadyExists
}
