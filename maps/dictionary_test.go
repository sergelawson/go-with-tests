package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	dictionnaire := Dictionary{"cat": "Domestic animal", "dog": "Human best friend"}

	t.Run("should find keyword definition in dictionary", func(t *testing.T) {

		got, _ := dictionnaire.Search("cat")

		want := "Domestic animal"

		assertString(t, got, want)

	})

	t.Run("should return error on  unknown keyword", func(t *testing.T) {

		got, err := dictionnaire.Search("unknown")

		want := ""

		if err == nil {
			t.Errorf("Should return error: %q", ErrorUnknownWord.Error())
		}

		assertError(t, ErrorUnknownWord, err)

		assertString(t, got, want)

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

