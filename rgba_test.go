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
	t.Run("it implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA8(0x00)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("it implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA8(0x00)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("it does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 1, int(unsafe.Sizeof(pxl.RGBA8(0x00))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		testCases := []struct {
			c pxl.RGBA8
			r uint32
			g uint32
			b uint32
			a uint32
		}{{c: pxl.RGBA8(0x00), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
			{c: pxl.RGBA8(0xff), r: 0xffffffff, g: 0xffffffff, b: 0xffffffff, a: 0xffffffff},
			{c: pxl.RGBA8(0x9b), r: 0xaaaaaaaa, g: 0x55555555, b: 0xaaaaaaaa, a: 0xffffffff}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				r, g, b, a := testCase.c.RGBA()
				assert.Equal(t, testCase.r, r)
				assert.Equal(t, testCase.g, g)
				assert.Equal(t, testCase.b, b)
				assert.Equal(t, testCase.a, a)
			})
		}
	})
	t.Run("Hex()", func(t *testing.T) {
		testCases := []struct {
			c   pxl.RGBA8
			hex string
		}{{c: pxl.RGBA8(0x00), hex: "00"},
			{c: pxl.RGBA8(0xff), hex: "ff"},
			{c: pxl.RGBA8(0x9b), hex: "9b"}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				assert.Equal(t, testCase.hex, testCase.c.Hex())
			})
		}
	})
}

func BenchmarkRGBA8(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA8(i).RGBA()
		}
	})
}

func TestRGBA16(t *testing.T) {
	t.Run("it implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA16(0x0000)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("it implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA16(0x0000)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("it does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 2, int(unsafe.Sizeof(pxl.RGBA16(0x0000))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		testCases := []struct {
			c pxl.RGBA16
			r uint32
			g uint32
			b uint32
			a uint32
		}{{c: pxl.RGBA16(0x0000), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
			{c: pxl.RGBA16(0xffff), r: 0xffffffff, g: 0xffffffff, b: 0xffffffff, a: 0xffffffff},
			{c: pxl.RGBA16(0xa5af), r: 0xaaaaaaaa, g: 0x55555555, b: 0xaaaaaaaa, a: 0xffffffff}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				r, g, b, a := testCase.c.RGBA()
				assert.Equal(t, testCase.r, r)
				assert.Equal(t, testCase.g, g)
				assert.Equal(t, testCase.b, b)
				assert.Equal(t, testCase.a, a)
			})
		}
	})
	t.Run("Hex()", func(t *testing.T) {
		testCases := []struct {
			c   pxl.RGBA16
			hex string
		}{{c: pxl.RGBA16(0x0000), hex: "0000"},
			{c: pxl.RGBA16(0xffff), hex: "ffff"},
			{c: pxl.RGBA16(0xa5af), hex: "a5af"}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				assert.Equal(t, testCase.hex, testCase.c.Hex())
			})
		}
	})
}

func BenchmarkRGBA16(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.RGBA16(i).RGBA()
		}
	})
}

func TestRGBA32(t *testing.T) {
	t.Run("it implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA32{}
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("it implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA32{}
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("it does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 4, int(unsafe.Sizeof(pxl.RGBA32{})))
	})
	t.Run("RGBA()", func(t *testing.T) {
		testCases := []struct {
			c pxl.RGBA32
			r uint32
			g uint32
			b uint32
			a uint32
		}{{c: pxl.RGBA32{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
			{c: pxl.RGBA32{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, r: 0xffffffff, g: 0xffffffff, b: 0xffffffff, a: 0xffffffff},
			{c: pxl.RGBA32{R: 0xaa, G: 0x55, B: 0xaa, A: 0xff}, r: 0xaaaaaaaa, g: 0x55555555, b: 0xaaaaaaaa, a: 0xffffffff}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				r, g, b, a := testCase.c.RGBA()
				assert.Equal(t, testCase.r, r)
				assert.Equal(t, testCase.g, g)
				assert.Equal(t, testCase.b, b)
				assert.Equal(t, testCase.a, a)
			})
		}
	})
	t.Run("Hex()", func(t *testing.T) {
		testCases := []struct {
			c   pxl.RGBA32
			hex string
		}{{c: pxl.RGBA32{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, hex: "00000000"},
			{c: pxl.RGBA32{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, hex: "ffffffff"},
			{c: pxl.RGBA32{R: 0xaa, G: 0x55, B: 0xaa, A: 0xff}, hex: "aa55aaff"}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				assert.Equal(t, testCase.hex, testCase.c.Hex())
			})
		}
	})
}

func TestRGBA64(t *testing.T) {
	t.Run("it implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA64{}
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("it implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA64{}
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("it does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 8, int(unsafe.Sizeof(pxl.RGBA64{})))
	})
	t.Run("RGBA()", func(t *testing.T) {
		testCases := []struct {
			c pxl.RGBA64
			r uint32
			g uint32
			b uint32
			a uint32
		}{{c: pxl.RGBA64{R: 0x0000, G: 0x0000, B: 0x0000, A: 0x0000}, r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
			{c: pxl.RGBA64{R: 0xffff, G: 0xffff, B: 0xffff, A: 0xffff}, r: 0xffffffff, g: 0xffffffff, b: 0xffffffff, a: 0xffffffff},
			{c: pxl.RGBA64{R: 0xaaaa, G: 0x5555, B: 0xaaaa, A: 0xffff}, r: 0xaaaaaaaa, g: 0x55555555, b: 0xaaaaaaaa, a: 0xffffffff}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				r, g, b, a := testCase.c.RGBA()
				assert.Equal(t, testCase.r, r)
				assert.Equal(t, testCase.g, g)
				assert.Equal(t, testCase.b, b)
				assert.Equal(t, testCase.a, a)
			})
		}
	})
	t.Run("Hex()", func(t *testing.T) {
		testCases := []struct {
			c   pxl.RGBA64
			hex string
		}{{c: pxl.RGBA64{R: 0x0000, G: 0x0000, B: 0x0000, A: 0x0000}, hex: "0000000000000000"},
			{c: pxl.RGBA64{R: 0xffff, G: 0xffff, B: 0xffff, A: 0xffff}, hex: "ffffffffffffffff"},
			{c: pxl.RGBA64{R: 0xaaaa, G: 0x5555, B: 0xaaaa, A: 0xffff}, hex: "aaaa5555aaaaffff"}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				assert.Equal(t, testCase.hex, testCase.c.Hex())
			})
		}
	})
}

func TestRGBA128(t *testing.T) {
	t.Run("it implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.RGBA128{}
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("it implements the std color interface", func(t *testing.T) {
		var v any = pxl.RGBA128{}
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("it does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 16, int(unsafe.Sizeof(pxl.RGBA128{})))
	})
	t.Run("RGBA()", func(t *testing.T) {
		testCases := []struct {
			c pxl.RGBA128
			r uint32
			g uint32
			b uint32
			a uint32
		}{{c: pxl.RGBA128{R: 0x00000000, G: 0x00000000, B: 0x00000000, A: 0x00000000}, r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x00000000},
			{c: pxl.RGBA128{R: 0xffffffff, G: 0xffffffff, B: 0xffffffff, A: 0xffffffff}, r: 0xffffffff, g: 0xffffffff, b: 0xffffffff, a: 0xffffffff},
			{c: pxl.RGBA128{R: 0xaaaaaaaa, G: 0x55555555, B: 0xaaaaaaaa, A: 0xffffffff}, r: 0xaaaaaaaa, g: 0x55555555, b: 0xaaaaaaaa, a: 0xffffffff}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				r, g, b, a := testCase.c.RGBA()
				assert.Equal(t, testCase.r, r)
				assert.Equal(t, testCase.g, g)
				assert.Equal(t, testCase.b, b)
				assert.Equal(t, testCase.a, a)
			})
		}
	})
	t.Run("Hex()", func(t *testing.T) {
		testCases := []struct {
			c   pxl.RGBA128
			hex string
		}{{c: pxl.RGBA128{R: 0x00000000, G: 0x00000000, B: 0x00000000, A: 0x00000000}, hex: "00000000000000000000000000000000"},
			{c: pxl.RGBA128{R: 0xffffffff, G: 0xffffffff, B: 0xffffffff, A: 0xffffffff}, hex: "ffffffffffffffffffffffffffffffff"},
			{c: pxl.RGBA128{R: 0xaaaaaaaa, G: 0x55555555, B: 0xaaaaaaaa, A: 0xffffffff}, hex: "aaaaaaaa55555555aaaaaaaaffffffff"}}
		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("%#v", testCase.c), func(t *testing.T) {
				assert.Equal(t, testCase.hex, testCase.c.Hex())
			})
		}
	})
}
