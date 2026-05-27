package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

const img = "images/space_shuttle.jpg"

func main() {
	imgGray := gocv.IMRead(img, gocv.IMReadColor)
	w := gocv.NewWindow("Image Read")
	w.IMShow(imgGray)
	fmt.Println("Press any key to exit.")
	fmt.Printf("keycode: %d detected. Exiting.\n", w.WaitKey(0))
}
