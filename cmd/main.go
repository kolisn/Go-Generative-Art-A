package main

import (
	"math/rand"
	"time"
	"log"
	"github.com/cramk/Go-Generative-Art-A/sketch"
)

var (
	sourceImgName = "source.jpg"
	outputImgName = "out.png"
	totalCycleCount = 5000
)

func main() {
	rand.Seed(time.Now().Unix())

	img, err := loadImage(sourceImgName)

	destWidth := 2000
	sketch := sketch.NewSketch{img, sketch.UserParams{		
		DestWidth:                destWidth,
		DestHeight:               2000,		
		StrokeRatio:              0.75,
		StrokeReduction:          0.002,
		AlphaIncrease:            0.06,
		StrokeInversionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(destWidth)),
		InitialAlpha:             0.1,
		MinEdgeCount:             3,
		MaxEdgeCount:             4,
	})

	// main loop
	for i := 0; i< totalCycleCount; i++ {
		sketch.Update()
	}

	saveOutput(s.Output(), outputImgName)

}
