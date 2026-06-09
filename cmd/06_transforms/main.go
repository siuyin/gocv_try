package main

import (
	"fmt"
	"image"
	"log"

	"gocv.io/x/gocv"
)

const imgFile = "images/flower.png"

func main() {
	img := gocv.IMRead(imgFile, gocv.IMReadUnchanged)
	defer img.Close()

	w := gocv.NewWindow("Source")
	defer w.Close()

	w.IMShow(img)
	c := center(&img)
	rW := gocv.NewWindow("Rotated")
	rW.IMShow(*rotate(&img, c, 45.0, 2))
	fmt.Println("Press any key to close window.")
	w.WaitKey(0)

}

func center(imgPtr *gocv.Mat) image.Point {
	dims := imgPtr.Size()
	return image.Point{dims[1] / 2, dims[0] / 2}
}

func rotate(imgPtr *gocv.Mat, center image.Point, angle float64, scale float64) *gocv.Mat {
	rotMat := gocv.GetRotationMatrix2D(center, angle, scale)
	rotImg := gocv.NewMat()
	size := imgPtr.Size()
	if err := gocv.WarpAffine(*imgPtr, &rotImg, rotMat, image.Point{size[1], size[0]}); err != nil {
		log.Fatal(err)
	}
	return &rotImg
}
