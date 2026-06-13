package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

const imgFile = "images/dog.png"

func main() {
	img := gocv.IMRead(imgFile, gocv.IMReadUnchanged)
	defer img.Close()

	fmt.Println("size:", img.Size())
	w := display("Source", img)
	defer w.Close()

	ln := line(img, image.Point{133, 144}, image.Point{259, 133}, color.RGBA{0, 32, 200, 255}, 32)
	defer ln.Close()
	wLn := display("Line", ln)
	defer wLn.Close()

	c := circ(img, image.Point{196, 174}, 24, color.RGBA{255, 0, 0, 127}, 4)
	defer c.Close()
	wC := display("Circle", c)
	defer wC.Close()

	fmt.Println("Press any key to close window.")
	w.WaitKey(0)
}

func display(title string, img gocv.Mat) *gocv.Window {
	w := gocv.NewWindow(title)
	if err := w.ResizeWindow(img.Cols(), img.Rows()); err != nil {
		log.Fatal(err)
	}
	w.IMShow(img)
	return w
}

func line(src gocv.Mat, start, end image.Point, c color.RGBA, thickness int) gocv.Mat {
	img := src.Clone()
	if err := gocv.Line(&img, start, end, c, thickness); err != nil {
		log.Fatal(err)
	}
	return img
}

func circ(src gocv.Mat, ctr image.Point, radius int, c color.RGBA, thickness int) gocv.Mat {
	ovr := src.Clone()
	defer ovr.Close()

	if err := gocv.Circle(&ovr, ctr, radius, c, thickness); err != nil {
		log.Fatal(err)
	}
	img := gocv.NewMat()

	alpha := float64(c.A) / 255.0
	if err := gocv.AddWeighted(src, 1-alpha, ovr, alpha, 0, &img); err != nil {
		log.Fatal(err)
	}

	return img
}
