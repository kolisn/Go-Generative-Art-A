package sketch

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type UserParams struct {
	StrokeRatio              float64
	DestWidth                int
	DestHeight               int
	InitialAlpha             float64
	StrokeReduction          float64
	AlphaIncrease            float64
	StrokeInversionThreshold float64
	StrokeJitter             int
	MinEdgeCount             int
	MaxEdgeCount             int
}

type Sketch struct {
	UserParams        // embed for easier access
	source            image.Image
	dc                *gg.Context
	sourceWidth       int
	sourceHeight      int
	strokeSize        float64
	initialStrokeSize float64
}

func NewSketch(source image.Image, userParams UserParams) *Sketch {
	s := &Sketch{UserParams: userParams}
	bounds := source.Bounds()
	s.sourceWidth, s.sourceHeight = bounds.Max.X, bounds.Max.Y
	s.initialStrokeSize = s.StrokeRatio * float64(s.DestWidth)
	s.strokeSize = s.initialStrokeSize

	canvas := gg.NewContext(s.DestWidth, s.DestHeight)
	canvas.SetColor(color.Black)
	canvas.DrawRectangle(0, 0, float64(s.DestWidth), float64(s.DestHeight))
	canvas.FillPreserve()

	s.source = source
	s.dc = canvas
	return s
}

func (s *Sketch) Update() {
	rndX := rand.Float64() * float64(s.sourceWidth)
	rndY := rand.Float64() * float64(s.sourceHeight)
	r, g, b := rgb255(s.source.At(int(rndX), int(rndY)))

	destX := rndX * float64(s.DestWidth) / float64(s.sourceWidth)
	destX += float64(randRange(s.StrokeJitter))
	destY := rndY * float64(s.DestHeight) / float64(s.sourceHeight)
	destY += float64(randRange(s.StrokeJitter))
	edges := s.MinEdgeCount + rand.Intn(s.MaxEdgeCount-s.MinEdgeCount+1)

	random := rand.Intn(20-1) + 1
	if random == 1 {

		const S = 1024
		for i := 0; i < rand.Intn(360); i += rand.Intn(20) {
			s.dc.Push()

			rndX1 := rand.Float64() * float64(s.sourceWidth)
			rndY1 := rand.Float64() * float64(s.sourceHeight)
			r1, g1, b1 := rgb255(s.source.At(int(rndX1), int(rndY1)))
			gradSize := rand.Float64()
			s.dc.SetRGBA255(r1, g1, b1, int(s.InitialAlpha))
			s.dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
			edges := s.MinEdgeCount + rand.Intn(s.MaxEdgeCount-s.MinEdgeCount+1)
			s.dc.DrawRegularPolygon(edges, destX, destY, s.strokeSize, gradSize)
			s.dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
			s.dc.Fill()
			s.dc.SetStrokeStyle(nil)
			s.dc.Pop()
		}

	} else {
		s.dc.SetRGBA255(r, g, b, int(s.InitialAlpha))
		s.dc.DrawRegularPolygon(edges, destX, destY, s.strokeSize, rand.Float64())
	}

	s.dc.Fill()

	if s.strokeSize <= s.StrokeInversionThreshold*s.initialStrokeSize {
		if (r+g+b)/3 < 128 {
			s.dc.SetRGBA255(255, 255, 255, int(s.InitialAlpha*2))
		} else {
			s.dc.SetRGBA255(0, 0, 0, int(s.InitialAlpha*2))
		}
	}

	s.strokeSize -= s.StrokeReduction * s.strokeSize
	s.InitialAlpha += s.AlphaIncrease
}

func (s *Sketch) Output() image.Image {
	return s.dc.Image()
}

func rgb255(c color.Color) (r, g, b int) {
	r0, g0, b0, _ := c.RGBA()
	return int(r0 / 257), int(g0 / 257), int(b0 / 257)
}

func randRange(max int) int {
	return -max + rand.Intn(2*max)
}
