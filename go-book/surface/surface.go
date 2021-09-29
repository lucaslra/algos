package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 300            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		drawShape(w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func drawShape(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, _, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, _, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, z, ok := corner(i+1, j+1)
			if !ok {
				continue
			}

			selectedColor := getColor(z)
			fmt.Fprintf(w, "<polygon style='fill:%s;stroke:purple;stroke-width:1;%g' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", selectedColor, z, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func getColor(z float64) string {
	if z > 0.9 {
		return "#FF0000"
	} else if z <= 0.9 && z > 0.6 {
		return "#990033"
	} else if z <= 0.6 && z > 0.3 {
		return "#330099"
	} else if z <= 0.3 && z > 0.1 {
		return "#003399"
	} else if z <= 0.1 && z > -0.1 {
		return "#009933"
	} else {
		return "#FFFFFF"
	}
}

func corner(i, j int) (float64, float64, float64, bool) {
	ok := true
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := zcomp(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(sx, 0) || math.IsInf(sy, 0) {
		ok = false
	}

	return sx, sy, z, ok
}

func zcomp(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
