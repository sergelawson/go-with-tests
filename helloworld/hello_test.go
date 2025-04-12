package helloworld

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello to person", func(t *testing.T) {
		name := "Serge"
		got := Hello(name)
		want := "Hello, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello when input is empty string", func(t *testing.T) {	
			got := Hello("")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
		
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}