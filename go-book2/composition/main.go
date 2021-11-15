package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	return hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type Cache struct {
	sync.Mutex
	m map[string]string
}

func (c *Cache) String() string {
	x := fmt.Sprintf("%v", c.m)
	return x
}

func main() {
	cp := ColoredPoint{Point{1, 2}, color.RGBA{1, 1, 1, 1}}

	fmt.Println(cp.X)
	fmt.Println(cp.Y)

	fmt.Println(cp.Point.X)
	fmt.Println(cp.Point.Y)

	fmt.Println(cp.X == cp.Point.X)
	fmt.Println(cp.Distance(Point{4, 6}))

	cache := Cache{m: make(map[string]string)}

	cache.Lock()
	cache.m["foo"] = "bar"
	cache.m["fizz"] = "buzz"
	cache.m["win"] = "11"
	cache.Unlock()

	fmt.Println(cache.String())
}
