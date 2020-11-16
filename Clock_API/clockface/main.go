package main

import (
	"LearningGo/Clock_API"
	"os"
	"time"
)

func main() {
	t := time.Now()
	Clock_API.SVGWriter(os.Stdout, t)
}
