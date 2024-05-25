package pxl_test

import (
	"fmt"
	"image/color"
	"pxl"
	"unsafe"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRGBA8(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA8(0x00)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA8(0x00)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 1, int(unsafe.Sizeof(pxl.RGBA8(0x00))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.RGBA8
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.RGBA8(0x00), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
				{c: pxl.RGBA8(0xff), r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.RGBA8(0x9b), r: 0x0000aaaa, g: 0x00005555, b: 0x0000aaaa, a: 0x0000ffff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					r, g, b, a := testCase.c.RGBA()
					assert.Equal(t, testCase.r, r)
					assert.Equal(t, testCase.g, g)
					assert.Equal(t, testCase.b, b)
					assert.Equal(t, testCase.a, a)
				})
			}
		})
	})
	t.Run("Hex()", func(t *testing.T) {
		t.Run("returns the correct value", func(t *testing.T) {
			testCases := []struct {
				c   pxl.RGBA8
				hex string
			}{{c: pxl.RGBA8(0x00), hex: "00"},
				{c: pxl.RGBA8(0xff), hex: "ff"},
				{c: pxl.RGBA8(0x9b), hex: "9b"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkRGBA8(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA8(i).RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA8(i).Hex()
		}
	})
}

func TestRGBA16(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA16(0x0000)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA16(0x0000)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 2, int(unsafe.Sizeof(pxl.RGBA16(0x0000))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.RGBA16
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.RGBA16(0x0000), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
				{c: pxl.RGBA16(0xffff), r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.RGBA16(0xa5af), r: 0x0000aaaa, g: 0x00005555, b: 0x0000aaaa, a: 0x0000ffff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					r, g, b, a := testCase.c.RGBA()
					assert.Equal(t, testCase.r, r)
					assert.Equal(t, testCase.g, g)
					assert.Equal(t, testCase.b, b)
					assert.Equal(t, testCase.a, a)
				})
			}
		})
	})
	t.Run("Hex()", func(t *testing.T) {
		t.Run("returns the correct value", func(t *testing.T) {
			testCases := []struct {
				c   pxl.RGBA16
				hex string
			}{{c: pxl.RGBA16(0x0000), hex: "0000"},
				{c: pxl.RGBA16(0xffff), hex: "ffff"},
				{c: pxl.RGBA16(0xa5af), hex: "a5af"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkRGBA16(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA16(i).RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA16(i).Hex()
		}
	})
}

func TestRGBA32(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA32{}
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA32{}
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 4, int(unsafe.Sizeof(pxl.RGBA32{})))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.RGBA32
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.RGBA32{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
				{c: pxl.RGBA32{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.RGBA32{R: 0xaa, G: 0x55, B: 0xaa, A: 0xff}, r: 0x0000aaaa, g: 0x00005555, b: 0x0000aaaa, a: 0x0000ffff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					r, g, b, a := testCase.c.RGBA()
					assert.Equal(t, testCase.r, r)
					assert.Equal(t, testCase.g, g)
					assert.Equal(t, testCase.b, b)
					assert.Equal(t, testCase.a, a)
				})
			}
		})
		t.Run("is equivalent to std lib's NRGBA", func(t *testing.T) {
			testCases := []struct {
				r uint8
				g uint8
				b uint8
				a uint8
			}{{r: 0x00, g: 0x00, b: 0x00, a: 0x00},
				{r: 0xff, g: 0xff, b: 0xff, a: 0xff},
				{r: 0xaa, g: 0x55, b: 0xaa, a: 0xff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					er, eg, eb, ea := color.NRGBA{R: testCase.r, G: testCase.g, B: testCase.b, A: testCase.a}.RGBA()
					ar, ag, ab, aa := pxl.RGBA32{R: testCase.r, G: testCase.g, B: testCase.b, A: testCase.a}.RGBA()
					assert.Equal(t, er, ar)
					assert.Equal(t, eg, ag)
					assert.Equal(t, eb, ab)
					assert.Equal(t, ea, aa)
				})
			}
		})
	})
	t.Run("Hex()", func(t *testing.T) {
		t.Run("returns the correct value", func(t *testing.T) {
			testCases := []struct {
				c   pxl.RGBA32
				hex string
			}{{c: pxl.RGBA32{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, hex: "00000000"},
				{c: pxl.RGBA32{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, hex: "ffffffff"},
				{c: pxl.RGBA32{R: 0xaa, G: 0x55, B: 0xaa, A: 0xff}, hex: "aa55aaff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkRGBA32(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA32{R: uint8(i), G: uint8(i), B: uint8(i), A: uint8(i)}.RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA32{R: uint8(i), G: uint8(i), B: uint8(i), A: uint8(i)}.Hex()
		}
	})
}

func TestRGBA64(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA64{}
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA64{}
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 8, int(unsafe.Sizeof(pxl.RGBA64{})))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.RGBA64
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.RGBA64{R: 0x0000, G: 0x0000, B: 0x0000, A: 0x0000}, r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
				{c: pxl.RGBA64{R: 0xffff, G: 0xffff, B: 0xffff, A: 0xffff}, r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.RGBA64{R: 0xaaaa, G: 0x5555, B: 0xaaaa, A: 0xffff}, r: 0x0000aaaa, g: 0x00005555, b: 0x0000aaaa, a: 0x0000ffff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					r, g, b, a := testCase.c.RGBA()
					assert.Equal(t, testCase.r, r)
					assert.Equal(t, testCase.g, g)
					assert.Equal(t, testCase.b, b)
					assert.Equal(t, testCase.a, a)
				})
			}
		})
		t.Run("is equivalent to std lib's NRGBA64", func(t *testing.T) {
			testCases := []struct {
				r uint16
				g uint16
				b uint16
				a uint16
			}{{r: 0x0000, g: 0x0000, b: 0x0000, a: 0x0000},
				{r: 0xffff, g: 0xffff, b: 0xffff, a: 0xffff},
				{r: 0xaaaa, g: 0x5555, b: 0xaaaa, a: 0xffff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					er, eg, eb, ea := color.NRGBA64{R: testCase.r, G: testCase.g, B: testCase.b, A: testCase.a}.RGBA()
					ar, ag, ab, aa := pxl.RGBA64{R: testCase.r, G: testCase.g, B: testCase.b, A: testCase.a}.RGBA()
					assert.Equal(t, er, ar)
					assert.Equal(t, eg, ag)
					assert.Equal(t, eb, ab)
					assert.Equal(t, ea, aa)
				})
			}
		})
	})
	t.Run("Hex()", func(t *testing.T) {
		t.Run("returns the correct value", func(t *testing.T) {
			testCases := []struct {
				c   pxl.RGBA64
				hex string
			}{{c: pxl.RGBA64{R: 0x0000, G: 0x0000, B: 0x0000, A: 0x0000}, hex: "0000000000000000"},
				{c: pxl.RGBA64{R: 0xffff, G: 0xffff, B: 0xffff, A: 0xffff}, hex: "ffffffffffffffff"},
				{c: pxl.RGBA64{R: 0xaaaa, G: 0x5555, B: 0xaaaa, A: 0xffff}, hex: "aaaa5555aaaaffff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkRGBA64(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA64{R: uint16(i), G: uint16(i), B: uint16(i), A: uint16(i)}.RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA64{R: uint16(i), G: uint16(i), B: uint16(i), A: uint16(i)}.Hex()
		}
	})
}

func TestRGBA128(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA128{}
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA128{}
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 16, int(unsafe.Sizeof(pxl.RGBA128{})))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.RGBA128
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.RGBA128{R: 0x00000000, G: 0x00000000, B: 0x00000000, A: 0x00000000}, r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
				{c: pxl.RGBA128{R: 0xffffffff, G: 0xffffffff, B: 0xffffffff, A: 0xffffffff}, r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.RGBA128{R: 0xaaaaaaaa, G: 0x55555555, B: 0xaaaaaaaa, A: 0xffffffff}, r: 0x0000aaaa, g: 0x00005555, b: 0x0000aaaa, a: 0x0000ffff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					r, g, b, a := testCase.c.RGBA()
					assert.Equal(t, testCase.r, r)
					assert.Equal(t, testCase.g, g)
					assert.Equal(t, testCase.b, b)
					assert.Equal(t, testCase.a, a)
				})
			}
		})
	})
	t.Run("Hex()", func(t *testing.T) {
		t.Run("returns the correct value", func(t *testing.T) {
			testCases := []struct {
				c   pxl.RGBA128
				hex string
			}{{c: pxl.RGBA128{R: 0x00000000, G: 0x00000000, B: 0x00000000, A: 0x00000000}, hex: "00000000000000000000000000000000"},
				{c: pxl.RGBA128{R: 0xffffffff, G: 0xffffffff, B: 0xffffffff, A: 0xffffffff}, hex: "ffffffffffffffffffffffffffffffff"},
				{c: pxl.RGBA128{R: 0xaaaaaaaa, G: 0x55555555, B: 0xaaaaaaaa, A: 0xffffffff}, hex: "aaaaaaaa55555555aaaaaaaaffffffff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkRGBA128(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA128{R: uint32(i), G: uint32(i), B: uint32(i), A: uint32(i)}.RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA128{R: uint32(i), G: uint32(i), B: uint32(i), A: uint32(i)}.Hex()
		}
	})
}

func TestRGBA256(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA256{}
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA256{}
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 32, int(unsafe.Sizeof(pxl.RGBA256{})))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.RGBA256
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.RGBA256{R: 0x0000000000000000, G: 0x0000000000000000, B: 0x0000000000000000, A: 0x0000000000000000}, r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
				{c: pxl.RGBA256{R: 0xffffffffffffffff, G: 0xffffffffffffffff, B: 0xffffffffffffffff, A: 0xffffffffffffffff}, r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.RGBA256{R: 0xaaaaaaaa00ff00ff, G: 0x5555555500ff00ff, B: 0xaaaaaaaa00ff00ff, A: 0xffffffff00ff00ff}, r: 0x0000aaaa, g: 0x00005555, b: 0x0000aaaa, a: 0x0000ffff}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					r, g, b, a := testCase.c.RGBA()
					assert.Equal(t, testCase.r, r)
					assert.Equal(t, testCase.g, g)
					assert.Equal(t, testCase.b, b)
					assert.Equal(t, testCase.a, a)
				})
			}
		})
	})
	t.Run("Hex()", func(t *testing.T) {
		t.Run("returns the correct value", func(t *testing.T) {
			testCases := []struct {
				c   pxl.RGBA256
				hex string
			}{{c: pxl.RGBA256{R: 0x0000000000000000, G: 0x0000000000000000, B: 0x0000000000000000, A: 0x0000000000000000}, hex: "0000000000000000000000000000000000000000000000000000000000000000"},
				{c: pxl.RGBA256{R: 0xffffffffffffffff, G: 0xffffffffffffffff, B: 0xffffffffffffffff, A: 0xffffffffffffffff}, hex: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"},
				{c: pxl.RGBA256{R: 0xaaaaaaaa00ff00ff, G: 0x5555555500ff00ff, B: 0xaaaaaaaa00ff00ff, A: 0xffffffff00ff00ff}, hex: "aaaaaaaa00ff00ff5555555500ff00ffaaaaaaaa00ff00ffffffffff00ff00ff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkRGBA256(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA256{R: uint64(i), G: uint64(i), B: uint64(i), A: uint64(i)}.RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA256{R: uint64(i), G: uint64(i), B: uint64(i), A: uint64(i)}.Hex()
		}
	})
}
