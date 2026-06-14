package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"log"

	"gocv.io/x/gocv"
)

const imgFile = "images/wood.png"

func main() {
	img := gocv.IMRead(imgFile, gocv.IMReadUnchanged)
	defer img.Close()

	fmt.Println("size:", img.Size())
	w := display("Source", img)
	defer w.Close()

	iden := identity(img)
	defer iden.Close()
	wI := display("Identity", iden)
	defer wI.Close()

	b := blur(img)
	defer b.Close()
	wB := display("Blurred", b)
	defer wB.Close()

	g := gaussBlur(img)
	defer g.Close()
	wG := display("Gaussian Blurred", g)
	defer wG.Close()

	fmt.Println("Press any key to close window.")
	w.WaitKey(0)
}

func display(title string, img gocv.Mat) *gocv.Window {
	w := gocv.NewWindow(title)
	if err := w.ResizeWindow(img.Cols(), img.Rows()); err != nil {
		log.Fatal(err)
	}
	w.IMShow(img)
	return w
}

func bytesFromFloat32(fs []float32) []byte {
	var b bytes.Buffer
	binary.Write(&b, binary.LittleEndian, fs)
	return b.Bytes()
}

func identity(src gocv.Mat) gocv.Mat {
	knl, err := gocv.NewMatFromBytes(3, 3, gocv.MatTypeCV32F, bytesFromFloat32([]float32{0, 0, 0, 0, 1, 0, 0, 0, 0}))
	if err != nil {
		log.Fatal(err)
	}

	dst := gocv.NewMat()
	if err := gocv.Filter2D(src, &dst, -1, knl, image.Point{-1, -1}, 0, gocv.BorderConstant); err != nil {
		log.Fatal(err)
	}

	return dst
}

func blur(src gocv.Mat) gocv.Mat {
	knl := gocv.Ones(5, 5, gocv.MatTypeCV32F)
	(&knl).DivideFloat(25)

	dst := gocv.NewMat()
	if err := gocv.Filter2D(src, &dst, -1, knl, image.Point{-1, -1}, 0, gocv.BorderConstant); err != nil {
		log.Fatal(err)
	}

	return dst
}

func gaussBlur(src gocv.Mat) gocv.Mat {
	dst := gocv.NewMat()
	if err := gocv.GaussianBlur(src, &dst, image.Point{5, 5}, 0, 0, gocv.BorderConstant); err != nil {
		log.Fatal(err)
	}
	return dst
}
