package pxl_test

import (
	"fmt"
	"image/color"
	"pxl"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestGray8(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.Gray8(0x00)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.Gray8(0x00)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 1, int(unsafe.Sizeof(pxl.Gray8(0x00))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.Gray8
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.Gray8(0x00), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x0000ffff},
				{c: pxl.Gray8(0xff), r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.Gray8(0x9b), r: 0x00009b9b, g: 0x00009b9b, b: 0x00009b9b, a: 0x0000ffff}}
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
				c   pxl.Gray8
				hex string
			}{{c: pxl.Gray8(0x00), hex: "000000ff"},
				{c: pxl.Gray8(0xff), hex: "ffffffff"},
				{c: pxl.Gray8(0x9b), hex: "9b9b9bff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkGray8(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray8(i).RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray8(i).Hex()
		}
	})
}

func TestGray16(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.Gray16(0x0000)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.Gray16(0x0000)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 2, int(unsafe.Sizeof(pxl.Gray16(0x0000))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.Gray16
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.Gray16(0x0000), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x0000ffff},
				{c: pxl.Gray16(0xffff), r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.Gray16(0xa5af), r: 0x0000a5af, g: 0x0000a5af, b: 0x0000a5af, a: 0x0000ffff}}
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
				c   pxl.Gray16
				hex string
			}{{c: pxl.Gray16(0x0000), hex: "000000000000ffff"},
				{c: pxl.Gray16(0xffff), hex: "ffffffffffffffff"},
				{c: pxl.Gray16(0xa5af), hex: "a5afa5afa5afffff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkGray16(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray16(i).RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray16(i).Hex()
		}
	})
}

func TestGray32(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.Gray32(0x00000000)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.Gray32(0x00000000)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 4, int(unsafe.Sizeof(pxl.Gray32(0x00000000))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.Gray32
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.Gray32(0x00000000), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x0000ffff},
				{c: pxl.Gray32(0xffffffff), r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.Gray32(0xaa55aaff), r: 0x0000aa55, g: 0x0000aa55, b: 0x0000aa55, a: 0x0000ffff}}
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
				c   pxl.Gray32
				hex string
			}{{c: pxl.Gray32(0x00000000), hex: "000000000000000000000000ffffffff"},
				{c: pxl.Gray32(0xffffffff), hex: "ffffffffffffffffffffffffffffffff"},
				{c: pxl.Gray32(0xaa55aaff), hex: "aa55aaffaa55aaffaa55aaffffffffff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkGray32(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray32(i).RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray32(i).Hex()
		}
	})
}

func TestGray64(t *testing.T) {
	t.Parallel()
	t.Run("implements the pxl color interface", func(t *testing.T) {
		var v any = pxl.Gray64(0x0000000000000000)
		_, ok := v.(pxl.Color)
		assert.True(t, ok)
	})
	t.Run("implements the std color interface", func(t *testing.T) {
		var v any = pxl.Gray64(0x0000000000000000)
		_, ok := v.(color.Color)
		assert.True(t, ok)
	})
	t.Run("does not exceed the expected number of bytes", func(t *testing.T) {
		assert.Equal(t, 8, int(unsafe.Sizeof(pxl.Gray64(0x0000000000000000))))
	})
	t.Run("RGBA()", func(t *testing.T) {
		t.Run("returns the correct values", func(t *testing.T) {
			testCases := []struct {
				c pxl.Gray64
				r uint32
				g uint32
				b uint32
				a uint32
			}{{c: pxl.Gray64(0x0000000000000000), r: 0x00000000, g: 0x00000000, b: 0x00000000, a: 0x0000ffff},
				{c: pxl.Gray64(0xffffffffffffffff), r: 0x0000ffff, g: 0x0000ffff, b: 0x0000ffff, a: 0x0000ffff},
				{c: pxl.Gray64(0xaaaa5555aaaaffff), r: 0x0000aaaa, g: 0x0000aaaa, b: 0x0000aaaa, a: 0x0000ffff}}
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
				c   pxl.Gray64
				hex string
			}{{c: pxl.Gray64(0x0000000000000000), hex: "000000000000000000000000000000000000000000000000ffffffffffffffff"},
				{c: pxl.Gray64(0xffffffffffffffff), hex: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"},
				{c: pxl.Gray64(0xaaaa5555aaaaffff), hex: "aaaa5555aaaaffffaaaa5555aaaaffffaaaa5555aaaaffffffffffffffffffff"}}
			for _, testCase := range testCases {
				t.Run(fmt.Sprintf("%+v", testCase), func(t *testing.T) {
					assert.Equal(t, testCase.hex, testCase.c.Hex())
				})
			}
		})
	})
}

func BenchmarkGray64(b *testing.B) {
	b.Run("RGBA()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray64(i).RGBA()
		}
	})
	b.Run("Hex()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pxl.Gray64(i).Hex()
		}
	})
}
