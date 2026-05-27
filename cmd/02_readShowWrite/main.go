package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

const (
	img       = "images/space_shuttle.jpg"
	outputImg = "scratch/grayscaleImage.jpg"
)

func main() {
	imgGray := gocv.IMRead(img, gocv.IMReadGrayScale)

	w := gocv.NewWindow("Image Read")
	defer w.Close()

	w.IMShow(imgGray)
	fmt.Println("Press any key to close window and write image.")
	fmt.Printf("keycode: %d detected.\n", w.WaitKey(0))

	if ok := gocv.IMWrite(outputImg, imgGray); ok {
		fmt.Println("Wrote ", outputImg)
		return
	}
	fmt.Println("Image write failed")
}
