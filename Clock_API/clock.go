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

// unit vector of the second hand at time 't'
func SecondHand(t time.Time) Point {
	return Point{150, 60}
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
