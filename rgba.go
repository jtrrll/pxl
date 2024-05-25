package pxl

import (
	"fmt"
)

// An RGBA8 is an 8-bit color represented by the additive RGBA color model.
// Each channel is represented by 2 bits, in the order `rrggbbaa`.
// RGBA8 is not alpha-premultiplied.
type RGBA8 uint8

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA8) RGBA() (r, g, b, a uint32) {
	a = uint32(c & 0x03)    // 00000000 00000000 00000000 000000aa
	r = uint32(c&0xC0) >> 6 // 00000000 00000000 00000000 000000rr
	g = uint32(c&0x30) >> 4 // 00000000 00000000 00000000 000000gg
	b = uint32(c&0x0C) >> 2 // 00000000 00000000 00000000 000000bb
	a *= 0x00005555
	r *= a / 0x00000003
	g *= a / 0x00000003
	b *= a / 0x00000003
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

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA16) RGBA() (r, g, b, a uint32) {
	a = uint32(c & 0x0f)       // 00000000 00000000 00000000 0000aaaa
	r = uint32(c&0xf000) >> 12 // 00000000 00000000 00000000 0000rrrr
	g = uint32(c&0x0f00) >> 8  // 00000000 00000000 00000000 0000gggg
	b = uint32(c&0x00f0) >> 4  // 00000000 00000000 00000000 0000bbbb
	a *= 0x00001111
	r *= a / 0x0000000f
	g *= a / 0x0000000f
	b *= a / 0x0000000f
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

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA32) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A) * 0x00000101
	r = uint32(c.R) * a / 0x000000ff
	g = uint32(c.G) * a / 0x000000ff
	b = uint32(c.B) * a / 0x000000ff
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

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA64) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A)
	r = uint32(c.R) * a / 0x0000ffff
	g = uint32(c.G) * a / 0x0000ffff
	b = uint32(c.B) * a / 0x0000ffff
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

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA128) RGBA() (r, g, b, a uint32) {
	a = c.A >> 16
	r = c.R >> 16 * a / 0x0000ffff
	g = c.G >> 16 * a / 0x0000ffff
	b = c.B >> 16 * a / 0x0000ffff
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c RGBA128) Hex() string {
	return fmt.Sprintf("%08x%08x%08x%08x", c.R, c.G, c.B, c.A)
}

// An RGBA256 is a 256-bit color represented by the additive RGBA color model.
// Each channel is represented by 64 bits.
// RGBA256 is not alpha-premultiplied.
type RGBA256 struct {
	R, G, B, A uint64
}

// Returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c RGBA256) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A >> 48)
	r = uint32(c.R>>48) * a / 0x0000ffff
	g = uint32(c.G>>48) * a / 0x0000ffff
	b = uint32(c.B>>48) * a / 0x0000ffff
	return
}

// Returns the hexadecimal code representing the RGBA color.
func (c RGBA256) Hex() string {
	return fmt.Sprintf("%016x%016x%016x%016x", c.R, c.G, c.B, c.A)
}
