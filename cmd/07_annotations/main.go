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

	ln := line(img, image.Point{133, 144}, image.Point{259, 133}, color.RGBA{0, 32, 200, 64}, 32)
	defer ln.Close()
	wLn := display("Line", ln)
	defer wLn.Close()

	c := circ(img, image.Point{196, 174}, 24, color.RGBA{255, 0, 0, 127}, -1) // use thickness of -1 for filled circle
	defer c.Close()
	wC := display("Circle", c)
	defer wC.Close()

	e := ellipse(img, image.Point{196, 174}, image.Point{45, 27}, 0, 90, 270, color.RGBA{255, 0, 0, 127}, -1)
	defer e.Close()
	wE := display("Ellipse", e)
	defer wE.Close()

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
	ovr := src.Clone()
	if err := gocv.RectangleWithParams(&ovr, image.Rectangle{start, end}, c, thickness, gocv.LineAA, 0); err != nil {
		log.Fatal(err)
	}

	alpha := float64(c.A) / 255.0
	img := alphaBlend(src, ovr, alpha)
	return img
}

func circ(src gocv.Mat, ctr image.Point, radius int, c color.RGBA, thickness int) gocv.Mat {
	ovr := src.Clone()
	defer ovr.Close()

	if err := gocv.CircleWithParams(&ovr, ctr, radius, c, thickness, gocv.LineAA, 0); err != nil {
		log.Fatal(err)
	}
	alpha := float64(c.A) / 255.0
	img := alphaBlend(src, ovr, alpha)

	return img
}

func alphaBlend(src, ovr gocv.Mat, alpha float64) gocv.Mat {
	img := gocv.NewMat()
	if err := gocv.AddWeighted(src, 1-alpha, ovr, alpha, 0, &img); err != nil {
		log.Fatal(err)
	}
	return img
}

func ellipse(src gocv.Mat, ctr, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int) gocv.Mat {
	ovr := src.Clone()
	defer ovr.Close()

	if err := gocv.EllipseWithParams(&ovr, ctr, axes, angle, startAngle, endAngle, c, thickness, gocv.LineAA, 0); err != nil {
		log.Fatal(err)
	}

	alpha := float64(c.A) / 255.0
	img := alphaBlend(src, ovr, alpha)
	return img
}
