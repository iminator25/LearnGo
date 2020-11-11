package Clock_API

import (
	"math"
	"time"
)

// A Point represents a two D Coordinate
type Point struct {
	X float64
	Y float64
}

const SecondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// unit vector of the second hand at time 't'
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * SecondHandLength, p.Y * SecondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}

}
