package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/cramk/Go-Generative-Art-A/sketch"
)

var (
	sourceImgName   = "source.jpg"
	outputImgName   = "out.png"
	totalCycleCount = 5000
)

func main() {
	rand.Seed(time.Now().Unix())

	img, err := loadImage(sourceImgName)

	if err != nil {
		log.Panicln(err)
	}

	destWidth := 2000
	s := sketch.NewSketch(img, sketch.UserParams{
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
	for i := 0; i < totalCycleCount; i++ {
		s.Update()
	}

	saveOutput(s.Output(), outputImgName)

}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("source image could not be loaded: %w", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("source image format could not be decoded %w", err)
	}

	return img, nil
}

func saveOutput(img image.Image, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}
