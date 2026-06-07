package main

import (
	"fmt"

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
}
