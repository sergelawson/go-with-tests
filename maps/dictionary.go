package dictionary

type Dictionary map[string]string

type DictionaryError string

func (error DictionaryError) Error() string {
	return string(error)
}

const (
	ErrorWordNotFound = DictionaryError("word not found")
	ErrorWordExists   = DictionaryError("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, exists := d[word]

	if !exists {
		return "", ErrorWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, exists := d[word]
	if exists {
		return ErrorWordExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, exists := d[word]

	if !exists {
		return ErrorWordNotFound

	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, exists := d[word]

	if !exists {
		return ErrorWordNotFound
	}

	delete(d, word)

	return nil
}
