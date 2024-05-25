package pxl

import "image/color"

// A Color is a color represented by a specific color model.
// It satisfies and extends the standard library's [image/color.Color] interface.
type Color interface {
	color.Color
	// Returns the hexadecimal code representing the RGBA color.
	Hex() string
}
