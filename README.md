# dither

[![GoDoc](https://godoc.org/github.com/bakerolls/dither?status.svg)](https://godoc.org/github.com/bakerolls/dither)
[![Go Report Card](https://goreportcard.com/badge/github.com/bakerolls/dither)](https://goreportcard.com/report/github.com/bakerolls/dither)

Package dither implements [ordered](https://en.wikipedia.org/wiki/Ordered_dithering) and [Floyd–Steinberg dithering](https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering).

```bash
$ go get github.com/bakerolls/dither
```

```go
func main() {
	r, err := os.Open("dag.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	img = dither.FloydSteinberg(img)
}
```

|                 Original                  |                     Ordered                      |                         Floyd-Steinberg                          |
| :---------------------------------------: | :----------------------------------------------: | :--------------------------------------------------------------: |
| ![Original](/cmd/dither/michelangelo.png) | ![Ordered](/cmd/dither/michelangelo-ordered.png) | ![Floyd-Steinberg](/cmd/dither/michelangelo-floyd-steinberg.png) |
