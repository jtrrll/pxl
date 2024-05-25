package pxl

import "fmt"

// TODO: Type comment
type Gray8 uint8

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c Gray8) RGBA() (r, g, b, a uint32) {
	gray := uint32(c) // 00000000 00000000 00000000 gggggggg
	gray *= 0x00000101
	r = gray
	g = gray
	b = gray
	a = 0x0000ffff
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c Gray8) Hex() string {
	return fmt.Sprintf("%02x%02x%02xff", c, c, c)
}

// TODO: Type comment
type Gray16 uint16

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c Gray16) RGBA() (r, g, b, a uint32) {
	gray := uint32(c) // 00000000 00000000 gggggggg gggggggg
	r = gray
	g = gray
	b = gray
	a = 0x0000ffff
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c Gray16) Hex() string {
	return fmt.Sprintf("%04x%04x%04xffff", c, c, c)
}

// TODO: Type comment
type Gray32 uint32

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c Gray32) RGBA() (r, g, b, a uint32) {
	gray := uint32(c >> 16)
	r = gray
	g = gray
	b = gray
	a = 0x0000ffff
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c Gray32) Hex() string {
	return fmt.Sprintf("%08x%08x%08xffffffff", c, c, c)
}

// TODO: Type comment
type Gray64 uint64

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c Gray64) RGBA() (r, g, b, a uint32) {
	gray := uint32(c >> 48)
	r = gray
	g = gray
	b = gray
	a = 0x0000ffff
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c Gray64) Hex() string {
	return fmt.Sprintf("%016x%016x%016xffffffffffffffff", c, c, c)
}
