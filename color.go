package pxl

// A Color is a color represented by a specific color model.
// It satisfies and extends the standard library's [color.Color] interface.
type Color interface {
	// Returns the hexadecimal code representing the RGBA color.
	Hex() string
	// Returns the alpha-premultiplied red, green, blue and alpha channels
	// as 32 bit values.
	//
	// An alpha-premultiplied color component c has been scaled by alpha (a),
	// so has valid values 0 <= c <= a.
	RGBA() (r, g, b, a uint32)
}
