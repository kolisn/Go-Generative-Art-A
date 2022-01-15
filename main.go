package main

import (
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

	img, err := loadImage(sourceImgName)

	if err != nil {
		log.Panicln(err)
	}

	destWidth := 2000
	s := sketch.NewSketch(img, sketch.UserParams{
		StrokeRatio:              0.75,
		DestWidth:                destWidth,
		DestHeight:               2000,
		InitialAlpha:             0.1,
		StrokeReduction:          0.002,
		AlphaIncrease:            0.01,
		StrokeInversionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(destWidth)),
		MinEdgeCount:             3,
		MaxEdgeCount:             10,
	})

	rand.Seed(time.Now().Unix())

	for i := 0; i < totalCycleCount; i++ {
		s.Update()
	}

	saveOutput(s.Output(), outputImgName)
}

func loadImage(src string) (image.Image, error) {
	file, _ := os.Open(sourceImgName)
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}

func saveOutput(img image.Image, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}
