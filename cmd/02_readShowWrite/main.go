package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

const (
	img          = "images/space_shuttle.jpg"
	outputFolder = "scratch"
	outputImg    = outputFolder + "/grayscaleImage.jpg"
)

func main() {

	w := gocv.NewWindow("Image Read")
	defer w.Close()

	imgGray := gocv.IMRead(img, gocv.IMReadGrayScale)
	w.IMShow(imgGray)
	fmt.Println("Press any key to close window and write image.")
	fmt.Printf("keycode: %d detected.\n", w.WaitKey(0))

	os.MkdirAll(outputFolder, 0755)
	if ok := gocv.IMWriteWithParams(outputImg, imgGray, []int{gocv.IMWriteJpegQuality, 70}); ok {
		fmt.Println("Wrote ", outputImg)
		return
	}
	fmt.Println("Image write failed")
}
