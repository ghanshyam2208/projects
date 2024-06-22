package structinterface

import "testing"

func TestShapes(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		actual := shape.Area()
		if expected != actual {
			t.Errorf("actual %.2f expected %.2f", actual, expected)
		}
	}

	t.Run("Test perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{height: 10.0, width: 10.0}
		actual := Perimeter(rectangle)
		expected := 40.0

		if expected != actual {
			t.Errorf("actual %.2f expected %.2f", actual, expected)
		}
	})

	t.Run("Test area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{height: 10.0, width: 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Test area of circle", func(t *testing.T) {
		circle := Circle{Radius: 10.0}

		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("Table driven tests", func(t *testing.T) {
		areaTest := []struct {
			shape  Shape
			actual float64
		}{
			{Rectangle{height: 12.0, width: 6.0}, 72.0},
			{Circle{Radius: 10}, 314.1592653589793},
			{Triangle{Base: 12, Height: 6}, 36.0},
		}

		for _, tt := range areaTest {
			actual := tt.shape.Area()
			if actual != tt.actual {
				t.Errorf("%#v actual %g want %g", tt.shape, actual, tt.actual)
			}
		}
	})

}
