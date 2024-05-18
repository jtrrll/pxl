package pxl

import (
	"fmt"
)

// An RGBA8 is an 8-bit color represented by the additive RGBA color model.
// Each channel is represented by 2 bits, in the order `rrggbbaa`.
// RGBA8 is not alpha-premultiplied.
type RGBA8 uint8

// Returns the alpha-premultiplied red, green, blue and alpha channels
// as 32 bit values.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA8) RGBA() (r, g, b, a uint32) {
	a = uint32(c & 0x03)    // 00000000 00000000 00000000 000000aa
	r = uint32(c&0xC0) >> 6 // 00000000 00000000 00000000 000000rr
	r *= a / 0x03
	g = uint32(c&0x30) >> 4 // 00000000 00000000 00000000 000000gg
	g *= a / 0x03
	b = uint32(c&0x0C) >> 2 // 00000000 00000000 00000000 000000bb
	b *= a / 0x03
	for i := 2; i <= 16; i *= 2 {
		r |= r << i
		g |= g << i
		b |= b << i
		a |= a << i
	}
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c RGBA8) Hex() string {
	return fmt.Sprintf("%02x", c)
}

// An RGBA16 is a 16-bit color represented by the additive RGBA color model.
// Each channel is represented by 4 bits, in the order `rrrrgggg bbbbaaaa`.
// RGBA16 is not alpha-premultiplied.
type RGBA16 uint16

// Returns the alpha-premultiplied red, green, blue and alpha channels
// as 32 bit values.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA16) RGBA() (r, g, b, a uint32) {
	a = uint32(c & 0x0f)       // 00000000 00000000 00000000 0000aaaa
	r = uint32(c&0xf000) >> 12 // 00000000 00000000 00000000 0000rrrr
	r *= a / 0x0f
	g = uint32(c&0x0f00) >> 8 // 00000000 00000000 00000000 0000gggg
	g *= a / 0x0f
	b = uint32(c&0x00f0) >> 4 // 00000000 00000000 00000000 0000bbbb
	b *= a / 0x0f
	for i := 4; i <= 16; i *= 2 {
		r |= r << i
		g |= g << i
		b |= b << i
		a |= a << i
	}
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c RGBA16) Hex() string {
	return fmt.Sprintf("%04x", c)
}

// An RGBA32 is a 32-bit color represented by the additive RGBA color model.
// Each channel is represented by 8 bits.
// RGBA32 is not alpha-premultiplied, and is equivalent to the standard library's [image.NRGBA].
type RGBA32 struct {
	R, G, B, A uint8
}

// Returns the alpha-premultiplied red, green, blue and alpha channels
// as 32 bit values.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA32) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A)
	r = uint32(c.R) * a / 0xff
	g = uint32(c.G) * a / 0xff
	b = uint32(c.B) * a / 0xff
	for i := 8; i <= 16; i *= 2 {
		r |= r << i
		g |= g << i
		b |= b << i
		a |= a << i
	}
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c RGBA32) Hex() string {
	return fmt.Sprintf("%02x%02x%02x%02x", c.R, c.G, c.B, c.A)
}

// An RGBA64 is a 64-bit color represented by the additive RGBA color model.
// Each channel is represented by 16 bits.
// RGBA64 is not alpha-premultiplied, and is equivalent to the standard library's [image.NRGBA64].
type RGBA64 struct {
	R, G, B, A uint16
}

// Returns the alpha-premultiplied red, green, blue and alpha channels
// as 32 bit values.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA64) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A)
	r = uint32(c.R) * a / 0xffff
	g = uint32(c.G) * a / 0xffff
	b = uint32(c.B) * a / 0xffff
	r |= r << 16
	g |= g << 16
	b |= b << 16
	a |= a << 16
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c RGBA64) Hex() string {
	return fmt.Sprintf("%04x%04x%04x%04x", c.R, c.G, c.B, c.A)
}

// An RGBA128 is a 128-bit color represented by the additive RGBA color model.
// Each channel is represented by 32 bits.
// RGBA128 is not alpha-premultiplied.
type RGBA128 struct {
	R, G, B, A uint32
}

// Returns the alpha-premultiplied red, green, blue and alpha channels
// as 32 bit values.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA128) RGBA() (r, g, b, a uint32) {
	a = c.A
	r = uint32(uint64(c.R) * uint64(c.A) / 0xffffffff)
	g = uint32(uint64(c.G) * uint64(c.A) / 0xffffffff)
	b = uint32(uint64(c.B) * uint64(c.A) / 0xffffffff)
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c RGBA128) Hex() string {
	return fmt.Sprintf("%08x%08x%08x%08x", c.R, c.G, c.B, c.A)
}
