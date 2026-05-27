package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

const img = "images/space_shuttle.jpg"

func main() {
	imgGray := gocv.IMRead(img, gocv.IMReadGrayScale)
	w := gocv.NewWindow("Image Read")
	w.IMShow(imgGray)
	fmt.Println("Press any key to exit.")
	w.WaitKey(0)
}
