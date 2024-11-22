package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	cx, cy, r float64
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(Titik{c.cx, c.cy}, p) <= c.r
}

func main() {
	var cx1, cy1, r1 float64
	fmt.Println("pusat (cx1, cy1) dan radius (r1) lingkaran 1:")
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 float64
	fmt.Println("pusat (cx2, cy2) dan radius (r2) lingkaran 2:")
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y float64
	fmt.Println("koordinat titik (x, y):")
	fmt.Scan(&x, &y)

	lingkaran1 := Lingkaran{cx1, cy1, r1}
	lingkaran2 := Lingkaran{cx2, cy2, r2}
	titik := Titik{x, y}

	isInCircle1 := didalam(lingkaran1, titik)
	isInCircle2 := didalam(lingkaran2, titik)

	if isInCircle1 && isInCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if isInCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if isInCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
