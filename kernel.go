package main

import (
	"unsafe"

	. "github.com/nithsua/goos/bootboi"
)

var (
	bootboi BOOTBOI
	fb      uint32
)

type psf2 struct {
	magic        uint32
	version      uint32
	headerSize   uint32
	flags        uint32
	numGlyph     uint32
	bytePerGlyph uint32
	height       uint32
	width        uint32
	glyphs       uint8
}

var binaryFontPSF []byte

/* offset = stride*y + x*pixelStride
 * the formula represents offset calculation for a 2D array of pixels
 * where each pixel stride is 4 bytes (XRGB (32-bit color(24 bit color + 8 bit padding)))
 * s represents the current row being processed
 * w represents the width of the frame buffer
 */

func _start() {
	var x, y int
	w := int(bootboi.FbWidth)
	h := int(bootboi.FbHeight)
	s := int(bootboi.FbScanline)

	if s > 0 {
		// Draw a Crosshair to validate the calibration of the framebuffer.
		for y = 0; y < h; y++ {
			// Jumps to the middle of the row and plots a pixel
			// s*y+((w*4)/2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(&fb), uintptr((s*y)+(w*2)))) = 0x00FFFFFF
		}

		for x = 0; x < w; x++ {
			// Jumps to the middle of the column and plots a pixel
			// s*(y/2)+(x*4). y would be equal to h due to the loop above.
			*(*uint32)(unsafe.Add(unsafe.Pointer(&fb), uintptr((s*(y/2))+(x*4)))) = 0x00FFFFFF
		}
	}
}

func main() {}
