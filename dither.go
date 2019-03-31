// Package dither allows dithering an image.
package dither

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Ordered dithers an image with a 2x2 matrix.
func Ordered(img image.Image) *image.Gray {
	b := img.Bounds()
	dst := image.NewGray(b)
	mask := [][]uint32{
		{math.MaxUint16 / 4, 2 * math.MaxUint16 / 4},
		{3 * math.MaxUint16 / 4, 0},
	}
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			dst.Pix[index(x, y, dst)] = math.MaxUint8
			dst.Set(x, y, color.White)
			if colorToGray(img.At(x, y)) <= mask[y%2][x%2] {
				dst.Pix[index(x, y, dst)] = 0
			}
		}
	}
	return dst
}

// FloydSteinberg dithering.
func FloydSteinberg(img image.Image) *image.Gray {
	b := img.Bounds()
	dst := image.NewGray(b)
	draw.Draw(dst, b, img, image.ZP, draw.Src)

	neighbours := []struct {
		dx, dy int
		pen    float64
	}{{1, 0, 7}, {-1, 1, 3}, {0, 1, 5}, {1, 1, 1}}
	pix := make([]int, len(dst.Pix))
	for i, p := range dst.Pix {
		pix[i] = int(p)
	}
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			prev := float64(pix[index(x, y, dst)])
			curr := math.Round(prev/math.MaxUint8) * math.MaxUint8
			dst.Pix[index(x, y, dst)] = uint8(curr)
			quantErr := prev - curr
			for _, n := range neighbours {
				i := index(x+n.dx, y+n.dy, dst)
				if 0 >= i || i >= len(pix) {
					continue
				}
				pix[i] += int(quantErr * n.pen / 16)
			}
		}
	}

	return dst
}

func index(x, y int, img *image.Gray) int {
	b := img.Bounds()
	return (y-b.Min.Y)*img.Stride + (x-b.Min.X)*1
}

func colorToGray(c color.Color) uint32 {
	r, g, b, _ := c.RGBA()
	return uint32(0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b))
}
