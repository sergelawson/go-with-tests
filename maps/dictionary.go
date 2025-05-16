package dictionary

import "errors"

type Dictionary map[string]string

var ErrorUnknownWord = errors.New("unknown keyword")

func (d Dictionary) Search(keyword string) (string, error) {
	definition, ok := d[keyword]

	if !ok {
		return "", ErrorUnknownWord
	}
	return definition, nil
}
