package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should calculate sum", func(t *testing.T) {

		t.Run("collection of 5 numbers", func(t *testing.T) {
			numbers := []int{1, 2, 3, 4, 5}

			got := Sum(numbers)

			want := 15

			if got != want {
				t.Errorf("got %d want %d given, %v", got, want, numbers)
			}
		})

 		t.Run("collection of 7 numbers", func(t *testing.T) {
			numbers := []int{1, 2, 3, 4, 5, 6, 7}

			got := Sum(numbers)

			want := 28

			if got != want {
				t.Errorf("got %d want %d given, %v", got, want, numbers)
			}
		}) 

	})

	t.Run("should calculate sum of slices", func(t *testing.T) {
		got := SumAll([]int{2,4,5}, []int{0, 6, 2})	
		want := []int{11, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("should calculate sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{2,4,5}, []int{0, 6, 2})	
		want := []int{9, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})


	t.Run("should calculate sum of tail of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{0}, []int{6, 2})	
		want := []int{0, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
