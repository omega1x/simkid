package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

// Draw new frame for the time t
func update(t int) string {
	//Create Canvas
	dc := gg.NewContext(SCREEN_W, SCREEN_H)

	// Set a background color
	dc.SetColor(color.RGBA{0, 0, 0, 255})
	dc.Clear()

	// Draw the position for all objects
	for n := range x[t] {
		dc.SetColor(color.RGBA{255, 230 - uint8(23*n), 30 + uint8(25*n), 255})
		dc.DrawCircle(x[t][n], y[t][n], 12)
		dc.Fill()
	}

	// Create frame
	file_name := fmt.Sprintf(FRAME_FILE, t)
	dc.SavePNG(file_name)
	return file_name
}
