package structPacakage

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("should calculate perimeter of a rectangle", func(t *testing.T) {

		rectangle := Rectangle{10.0, 10.0}

		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{{name: "Rectangle", shape: Rectangle{5.0, 3.0}, hasArea: 15.0}, {name: "Circle", shape: Circle{5.0}, hasArea: 78.53981633974483}}

	t.Run("should calculate the area of each  shape", func(t *testing.T) {

		for _, tt := range areaTests {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("shape %#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		}

	})
}
