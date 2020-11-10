package main

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return .5 * t.Height * t.Base
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2*rectangle.Width + 2*rectangle.Height
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
