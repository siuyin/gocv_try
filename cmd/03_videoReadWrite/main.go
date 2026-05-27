package main

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

const vid = "images/small.mp4"

func main() {
	cap, err := gocv.OpenVideoCapture(vid)
	if err != nil {
		log.Fatal(err)
	}
	defer cap.Close()

	fps := cap.Get(gocv.VideoCaptureFPS)
	fmt.Printf("FPS of %s is %v frames per second.\n", vid, fps)

	frameCount := cap.Get(gocv.VideoCaptureFrameCount)
	fmt.Printf("%s has %g, %s encoded frames.\n", vid, frameCount, cap.CodecString())

	w := gocv.NewWindow(vid)
	defer w.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Println("Press q to quit.")
	for {
		if ok := cap.Read(&img); !ok {
			fmt.Println("error reading: ", vid)
			return
		}

		w.IMShow(img)
		if key := w.WaitKey(int(1.0 / fps * 1000)); key == 'q' {
			break
		}
	}

}
