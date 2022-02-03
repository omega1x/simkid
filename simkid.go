package main

import (
	"fmt"
)

const (
	FRAMES_N = 720
	OBJECT_N = 10

	SCREEN_W = 1920
	SCREEN_H = 1080

	FRAME_FILE = "./frameset/frame-%04d.png"
)

// Space initializers
var x, y, vx, vy [FRAMES_N][OBJECT_N]float64

/* var m [OBJECT_N]float32 */

func main() {

	// Trajectory simulator
	for t := range x {
		sim(t)
	}

	// Draw frames
	for t := range x {
		fmt.Printf("\r[simkid]: write frame to <%s>", update(t))
	}
	fmt.Println("\n[simkid]: all done, thanks!  :)")
}
