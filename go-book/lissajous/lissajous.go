package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black,
	color.RGBA{0x03, 0xA0, 0x62, 0xff},
	color.RGBA{0xff, 0x14, 0x93, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x63, 0x47, 0xff},
	color.RGBA{0x1e, 0x90, 0xff, 0xff}}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 8      // number of complete x oscillator revolutions
		res     = 0.0001 // angular resolution
		size    = 250    // image canvas covers [-size..+size]
		nframes = 64     // number of animation frames
		delay   = 5      // delay between frames
	)
	freq := rand.Float64() + 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		colorIndex := 1
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if colorIndex >= len(palette) {
				colorIndex = 1
			}

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))

			colorIndex++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
