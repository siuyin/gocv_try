package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"gocv.io/x/gocv"
)

var colors = []color.RGBA{
	{R: 255, G: 255, B: 255, A: 0}, // White
	{R: 255, G: 255, B: 0, A: 0},   // Yellow
	{R: 0, G: 255, B: 255, A: 0},   // Cyan
	{R: 0, G: 255, B: 0, A: 0},     // Green
	{R: 255, G: 0, B: 255, A: 0},   // Magenta
	{R: 255, G: 0, B: 0, A: 0},     // Red
	{R: 0, G: 0, B: 255, A: 0},     // Blue
	{R: 0, G: 0, B: 0, A: 0},       // Black
}

func main() {
	for i := 0; i < 300000; i++ {
		displayFor(fmt.Sprintf("%d", i+1), 1000*time.Millisecond)
	}
}

func displayFor(title string, t time.Duration) {
	img := perturbColorBars()
	defer img.Close()
	w := display(title, img)
	defer w.Close()

	w.WaitKey(int(t / 1000000))
}

func display(title string, img gocv.Mat) *gocv.Window {
	w := gocv.NewWindow(title)
	if err := w.ResizeWindow(img.Cols(), img.Rows()); err != nil {
		log.Fatal(err)
	}
	w.IMShow(img)
	return w
}

func perturbColorBars() gocv.Mat {
	const (
		width  = 1280
		height = 720
	)

	img := gocv.NewMatWithSize(height, width, gocv.MatTypeCV8UC3)
	numBars := len(colors)
	barWidth := width / numBars
	for i := 0; i < numBars; i++ {
		startX := i * barWidth
		endX := (i + 1) * barWidth
		if i == numBars-1 {
			endX = width
		}
		rect := image.Rect(startX, 0, endX, height)
		gocv.Rectangle(&img, rect, colors[i], -1)
	}

	perturbColors(127)

	return img
}

func perturbColors(maxPerturb int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range colors {
		rOffset := r.Intn(maxPerturb*2+1) - maxPerturb
		gOffset := r.Intn(maxPerturb*2+1) - maxPerturb
		bOffset := r.Intn(maxPerturb*2+1) - maxPerturb

		colors[i].R = clamp(int(colors[i].R) + rOffset)
		colors[i].G = clamp(int(colors[i].G) + gOffset)
		colors[i].B = clamp(int(colors[i].B) + bOffset)
	}
}

func clamp(val int) uint8 {
	if val < 0 {
		return 0
	}
	if val > 255 {
		return 255
	}
	return uint8(val)
}
