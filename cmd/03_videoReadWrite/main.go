package main

import (
	"fmt"
	"log"

	"gocv.io/x/gocv"
)

// const src = "images/small.mp4"
const src = 0

func main() {
	imgPtr := captureLastVideoFrame()
	defer imgPtr.Close()

	if ok := gocv.IMWriteWithParams("scratch/capturedImage.jpg", *imgPtr, []int{gocv.IMWriteJpegQuality, 70}); ok {
		return
	}
	fmt.Println("Image write failed")

}

func captureLastVideoFrame() *gocv.Mat {
	//cap, err := gocv.OpenVideoCapture(src)
	cap, err := gocv.OpenVideoCapture(src) // 0: /dev/video0
	if err != nil {
		log.Fatal(err)
	}
	defer cap.Close()

	printStructuralMetadata(cap)

	w := gocv.NewWindow("Video Capture")
	defer w.Close()

	img := gocv.NewMat()

	fmt.Println("Press c to capture image and close window.")
	w.ResizeWindow(640, 480)
	for {
		if ok := cap.Read(&img); !ok {
			fmt.Println("error reading: ", src)
			return &img
		}

		//if key := w.WaitKey(int(1.0 / fps * 1000)); key == 'q' {
		if key := w.WaitKey(1); key == 'c' {
			fmt.Println("'c' pressed. Writing scratch/capturedImage.jpg")
			break
		}
		// w.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
		w.IMShow(img)
	}

	return &img
}

func printStructuralMetadata(cap *gocv.VideoCapture) {
	fmt.Println("Video Metadata")

	fps := cap.Get(gocv.VideoCaptureFPS)
	fmt.Printf("FPS of %v is %v frames per second.\n", src, fps)

	frameCount := cap.Get(gocv.VideoCaptureFrameCount)
	fmt.Printf("%v has %g, %s encoded frames.\n", src, frameCount, cap.CodecString())
}
