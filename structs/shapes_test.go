package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// Menggunakan table driven test
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 13}, hasArea: 156},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 7}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
	/*
		testArea := func(t *testing.T, s Shape, want float64) {
			t.Helper()
			got := s.Area()
			if got != want {
				t.Errorf("got %g want %g", got, want)
			}
		}
		t.Run("Rectangle", func(t *testing.T) {
			r := Rectangle{12, 13}
			testArea(t, r, 156.0)
		})

		t.Run("Circle", func(t *testing.T) {
			c := Circle{10}
			testArea(t, c, 314.1592653589793)
		})
	*/
}
