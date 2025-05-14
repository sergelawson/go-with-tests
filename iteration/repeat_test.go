package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat a character", func(t *testing.T) {

		repeated := Repeat("a", 1)

		expected := "a"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat 5 characters", func(t *testing.T) {

		repeated := Repeat("y", 5)

		expected := "yyyyy"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat(){
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}
