package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/bake/dither"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s in.png out.png", os.Args[0])
		os.Exit(1)
	}
	r, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	img = dither.FloydSteinberg(img)
	w, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(w, img); err != nil {
		log.Fatal(err)
	}
}
