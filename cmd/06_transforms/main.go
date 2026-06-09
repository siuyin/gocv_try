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
	r, c := center(&img)
	fmt.Println(r, c)
	rW := gocv.NewWindow("Rotated")
	rW.IMShow(*rotate(&img, image.Point{r, c}, 45.0, 1.0))
	fmt.Println("Press any key to close window.")
	w.WaitKey(0)

}

func center(imgPtr *gocv.Mat) (row, col int) {
	dims := imgPtr.Size()
	return dims[1] / 2, dims[0] / 2
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
