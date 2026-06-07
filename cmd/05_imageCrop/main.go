package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

const imgFile = "images/flower.png"

func main() {
	img := gocv.IMRead(imgFile, gocv.IMReadUnchanged)
	defer img.Close()
	dims := img.Size()
	fmt.Print("dimensions:")
	for _, d := range dims {
		fmt.Printf(" %d", d)
	}
	fmt.Println("\nNumber of channels:", img.Channels())

	limits := image.Rectangle{image.Point{150, 80}, image.Point{330, 280}}
	cropped := img.Region(limits)
	defer cropped.Close()

	w := gocv.NewWindow("Image Crop")
	defer w.Close()

	w.IMShow(cropped)
	fmt.Println("Press any key to close window.")
	w.WaitKey(0)
}
