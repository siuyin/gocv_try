package main

import (
	"fmt"
	"image"
	"log"

	"gocv.io/x/gocv"
)

const imgFile = "images/space_shuttle.jpg"

func main() {
	w := gocv.NewWindow("Original Image")
	defer w.Close()

	img := gocv.IMRead(imgFile, gocv.IMReadUnchanged)
	fmt.Println("Original Image", img.Cols(), img.Rows())
	if err := w.ResizeWindow(img.Cols(), img.Rows()); err != nil {
		log.Fatal(err)
	}
	if err := w.IMShow(img); err != nil {
		log.Fatal(err)
	}

	// Enlargement --------------------------------------------------------------------------------
	enlargeImg := gocv.NewMat()
	if err := gocv.Resize(img, &enlargeImg, image.Point{}, 2, 2, gocv.InterpolationCubic); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enlarged Image", enlargeImg.Cols(), enlargeImg.Rows())
	wEnlarge := gocv.NewWindow(fmt.Sprintf("%vx%v", enlargeImg.Cols(), enlargeImg.Rows()))
	if err := wEnlarge.ResizeWindow(enlargeImg.Cols(), enlargeImg.Rows()); err != nil {
		log.Fatal(err)
	}
	defer wEnlarge.Close()
	if err := wEnlarge.IMShow(enlargeImg); err != nil {
		log.Fatal(err)
	}

	// Reduction --------------------------------------------------------------------------------
	reduceImg := gocv.NewMat()
	if err := gocv.Resize(img, &reduceImg, image.Point{300, 300}, 0, 0, gocv.InterpolationArea); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reduced Image", reduceImg.Cols(), reduceImg.Rows())
	wReduced := gocv.NewWindow(fmt.Sprintf("%vx%v", reduceImg.Cols(), reduceImg.Rows()))
	if err := wReduced.ResizeWindow(reduceImg.Cols(), reduceImg.Rows()); err != nil {
		log.Fatal(err)
	}
	defer wReduced.Close()
	if err := wReduced.IMShow(reduceImg); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Press any key to close windows and exit.")
	w.WaitKey(0)

}
