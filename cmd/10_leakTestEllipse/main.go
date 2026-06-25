package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Drawing multiple ellipses on to an canvas.")
	img := gocv.NewMatWithSize(1024, 1024, gocv.MatTypeCV8UC3)
	defer img.Close()

	w := gocv.NewWindow("Ellipses")
	defer w.Close()

	if err := w.ResizeWindow(img.Cols(), img.Rows()); err != nil {
		log.Fatal(err)
	}

	angle := 0.0
	for i := 0; i < 200; i++ {
		ellipse(&img, angle, randomColor())
		if err := w.IMShow(img); err != nil {
			log.Fatal(err)
		}

		w.WaitKey(1)
		angle += 7
	}

	fmt.Println("Press any key to exit.")
	w.WaitKey(0)
}

func ellipse(imgPtr *gocv.Mat, angle float64, c color.RGBA) {
	el := struct {
		maj       int
		mnr       int
		ctr       image.Point
		thickness int
		lineType  gocv.LineType
	}{
		500,
		64,
		image.Point{1024 / 2, 1024 / 2},
		2,
		gocv.LineAA,
	}

	if err := gocv.EllipseWithParams(imgPtr, el.ctr, image.Point{el.maj, el.mnr},
		angle, 0, 360, c, el.thickness, el.lineType, 0); err != nil {
		log.Fatal(err)
	}

}

func randomColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 0}
}
