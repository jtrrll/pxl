// Package pxl extends the standard library's image package to enhance 2-D image and color manipulation.
package pxl

import (
	"image"
	"image/color"
)

// An Image is a finite rectangular grid of pixels.
// It satisfies and extends the standard library's [image.Image] interface.
type Image[T Color] interface {
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) T
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() image.Rectangle
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
}
