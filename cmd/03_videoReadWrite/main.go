package main

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

// const src = "images/small.mp4"
const src = 0

func main() {
	//cap, err := gocv.OpenVideoCapture(src)
	cap, err := gocv.OpenVideoCapture(src) // 0: /dev/video0
	if err != nil {
		log.Fatal(err)
	}
	defer cap.Close()

	fps := cap.Get(gocv.VideoCaptureFPS)
	fmt.Printf("FPS of %v is %v frames per second.\n", src, fps)

	frameCount := cap.Get(gocv.VideoCaptureFrameCount)
	fmt.Printf("%v has %g, %s encoded frames.\n", src, frameCount, cap.CodecString())

	w := gocv.NewWindow("Video Capture")
	defer w.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Println("Press q to quit.")
	for {
		if ok := cap.Read(&img); !ok {
			fmt.Println("error reading: ", src)
			return
		}

		//if key := w.WaitKey(int(1.0 / fps * 1000)); key == 'q' {
		if key := w.WaitKey(1); key == 'q' {
			fmt.Println("'q' pressed. Exiting now.")
			break
		}
		w.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
		w.IMShow(img)
	}

}
