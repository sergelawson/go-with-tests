package generics

import "testing"


func TestAssertFuction(t *testing.T){
	t.Run("Asserting Integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 2, 1)
	})

	t.Run("asserting on strings", func(t *testing.T) {
	AssertEqual(t, "hello", "hello")
	AssertNotEqual(t, "hello", "Grace")
})
}

func AssertEqual[T comparable](t testing.TB, got, want T){
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got , want T){
	t.Helper()
	if got == want {
		t.Errorf("did't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}