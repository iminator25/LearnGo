package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f but wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{10, 10}, 50.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
		}
	}

}

//func TestArea(t *testing.T) {
//	checkArea := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	t.Run("rectangle", func(t *testing.T) {
//		rectangle := Rectangle{12,6}
//		checkArea(t, rectangle, 72)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		checkArea(t, circle, 314.1592653589793 )
//	})
//}

//func TestArea(t *testing.T) {
//
//	t.Run("rectangle", func(t *testing.T) {
//		rectangle := Rectangle{10.0,10.0}
//		got := rectangle.Area()
//		want := 100.0
//
//		if got != want {
//			t.Errorf("got %.2f but wanted %.2f", got, want)
//		}
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		got := circle.Area()
//		want := 314.1592653589793
//
//		if got != want {
//			t.Errorf("got %g but wanted %g", got, want)
//		}
//
//	})
//
//}
