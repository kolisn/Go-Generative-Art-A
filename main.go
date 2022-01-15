package main

import (
	"github.com/cramk/Go-Generative-Art-A/sketch"
)

func main() {
	destWidth := 2000
	params := sketch.UserParams{
		StrokeRatio:              0.75,
		DestWidth:                destWidth,
		DestHeight:               2000,
		InitialAlpha:             0.1,
		StrokeReduction:          0.002,
		AlphaIncrease:            0.06,
		StrokeInversionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(destWidth)),
		MinEdgeCount:             3,
		MaxEdgeCount:             4,
	}
}
