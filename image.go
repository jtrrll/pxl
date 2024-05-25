// Package pxl extends the standard library's image package to enhance 2-D image and color manipulation.
package pxl

import (
	"image"
)

// An Image is a finite rectangular grid of pixels.
// It satisfies and extends the standard library's [image.Image] interface.
type Image[T Color] interface {
	image.Image
	// Returns the color of the pixel at (x, y).
	// Get(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// Get(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right pixel.
	Get(x, y int) T
	// Sets the color of the pixel at (x, y).
	// Set(Bounds().Min.X, Bounds().Min.Y) updates the upper-left pixel of the grid.
	// Set(Bounds().Max.X-1, Bounds().Max.Y-1) updates the lower-right pixel.
	Set(x, y int, c T)
}
