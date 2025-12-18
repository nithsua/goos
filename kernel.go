package main

import . "github.com/nithsua/goos/bootboi"

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
