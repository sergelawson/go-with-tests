package dictionary

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"cat": "Domestic animal", "dog": "Human best friend"}

	t.Run("should find keyword definition in dictionary", func(t *testing.T) {

		word := "cat"
		want := "Domestic animal"

		assertDefinition(t, dictionary, word ,want)

	})

	t.Run("should return error on  unknown keyword", func(t *testing.T) {

		got, err := dictionary.Search("unknown")

		want := ""

		if err == nil {
			t.Errorf("Should return error: %q", ErrorWordNotFound.Error())
		}

		assertError(t, ErrorWordNotFound, err)

		assertString(t, got, want)

	})

	t.Run("should add word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "bird"
		definition := "Animal with wings"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition )


	})

	t.Run("should not add word that already exits", func(t *testing.T) {
		word := "bird"
		definition := "Animal with wings"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)

		assertError(t, err, ErrorWordExists)

	})

	t.Run("should update word from dictionary", func(t *testing.T) {

		word := "bird"
		definition := "Animal with wings"

		dictionary := Dictionary{word: definition}


		update := "Animal with wings that can fly"

		dictionary.Update(word, update)

		assertDefinition(t, dictionary, word, update)
	})

	t.Run("update should fail if word does not exists", func(t *testing.T) {

		word := "bird"
		definition := "Animal with wings"

		dictionary := Dictionary{word: definition}


		update := "Animal with wings that can fly"

		err := dictionary.Update("unknown", update)


		assertDefinition(t, dictionary, word, definition)

		assertError(t, err, ErrorWordNotFound)


	})

	t.Run("should delete word from dictionary", func(t *testing.T) {

		word := "bird"
		definition := "Animal with wings"

		dictionary := Dictionary{word: definition}



		err := dictionary.Delete(word)

		if err != nil {

			t.Errorf("Error deleting: %q", err)
		}

		_, err = dictionary.Search(word)

		assertError(t, err, ErrorWordNotFound)
	})

	t.Run("delete should return error if word  does not exists", func(t *testing.T) {

		dictionary := Dictionary{}

		err := dictionary.Delete("unkown")


		assertError(t, err, ErrorWordNotFound)
	})
}

func assertString(t testing.TB, got string, want string) {
	t.Helper()

	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {

		t.Errorf("Unexpected error message want: %q, got: %q", want, got)
	}
}

func assertDefinition(t testing.TB, d Dictionary, word string, definition string){

	got, err := d.Search(word)

	if err != nil {
		t.Errorf("Error : %q", err.Error())
	}

	assertString(t, got, definition)
}
