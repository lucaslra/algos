package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black,
	color.RGBA{0x03, 0xA0, 0x62, 0xff},
	color.RGBA{0xff, 0x14, 0x93, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x63, 0x47, 0xff},
	color.RGBA{0x1e, 0x90, 0xff, 0xff}}

var (
	cycles  = 8
	res     = 0.0001
	size    = 250
	nframes = 64
	delay   = 5
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		var err error
		for k, v := range r.Form {
			switch k {
			case "cycles":
				cycles, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Printf("Failed to convert %q to int. %q value is %v", k, k, v)
				}
			case "size":
				size, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Printf("Failed to convert %q to int. %q value is %v", k, k, v)
				}
			case "nframes":
				nframes, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Printf("Failed to convert %q to int. %q value is %v", k, k, v)
				}
			case "delay":
				delay, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Printf("Failed to convert %q to int. %q value is %v", k, k, v)
				}
			}
		}

		lissajous(w, cycles, size, nframes, delay, res)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles, size, nframes, delay int, res float64) {
	freq := rand.Float64() + 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		colorIndex := 1
		for t := 0.0; t < (float64(cycles) * 2 * math.Pi); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if colorIndex >= len(palette) {
				colorIndex = 1
			}

			xPos := size + int(x*float64(size)+0.5)
			yPos := size + int(y*float64(size)+0.5)
			img.SetColorIndex(xPos, yPos, uint8(colorIndex))

			colorIndex++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
