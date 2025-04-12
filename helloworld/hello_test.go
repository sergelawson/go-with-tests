package helloworld

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello to person", func(t *testing.T) {
		name := "Serge"
		got := Hello(name)
		want := "Hello, " + name
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello to empty string", func(t *testing.T) {	
			got := Hello("")
			want := "Hello, World"
		
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
	})

}