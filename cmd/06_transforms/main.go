package main

import (
	"bytes"
	"encoding/binary"
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

	tW := gocv.NewWindow("Translated")
	dims := img.Size()
	//tW.IMShow(*translate(&img, float32(-dims[0])/4.0, float32(dims[1]/4.0)))
	tW.IMShow(*translate(&img, float32(-dims[0])/4.0, 0))
	//tW.IMShow(*translate(&img, 0.0, 0.0))

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

func bytesFromFloat32(fs []float32) []byte {
	var b bytes.Buffer
	binary.Write(&b, binary.LittleEndian, fs)
	return b.Bytes()
}

func translate(imgPtr *gocv.Mat, x, y float32) *gocv.Mat {
	//tMat, err := gocv.NewMatFromBytes(2, 3, gocv.MatTypeCV32FC1, bytesFromFloat32([]float32{1, 0, x, 0, 1, y}))
	tMat, err := gocv.NewMatFromBytes(2, 3, gocv.MatTypeCV32F, bytesFromFloat32([]float32{1, 0, x, 0, 1, y}))
	if err != nil {
		log.Fatal(err)
	}
	//tMat := gocv.NewMatWithSizes([]int{2, 3}, gocv.MatTypeCV32F)
	//tMat.SetFloatAt(0, 0, 1)
	//tMat.SetFloatAt(0, 1, 0)
	//tMat.SetFloatAt(0, 2, x)
	//tMat.SetFloatAt(1, 0, 0)
	//tMat.SetFloatAt(1, 1, 1)
	//tMat.SetFloatAt(1, 2, y)

	transImg := gocv.NewMat()
	size := imgPtr.Size()
	if err := gocv.WarpAffine(*imgPtr, &transImg, tMat, image.Point{size[1], size[0]}); err != nil {
		log.Fatal(err)
	}
	return &transImg
}
