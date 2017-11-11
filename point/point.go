package point

import (
	"math/rand"

	"github.com/fogleman/gg"
)

type Point struct {
	X, Y float64
}

type Points []Point

func (point Point) Draw(graphics *gg.Context) {
	graphics.DrawPoint(point.X, point.Y, 2)
	graphics.SetRGB(0, 0, 0)
}

func (points Points) Draw(graphics *gg.Context) {
	for _, p := range points {
		p.Draw(graphics)
	}
}

func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func (points Points) Less(i, j int) bool {
	if points[i].X == points[j].X {
		return points[i].Y < points[j].Y
	}
	return points[i].X < points[j].X
}

func (points Points) Len() int {
	return len(points)
}

func CrossProduct(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

func RandomPoints(length, min, max int) Points {
	var result Points
	for i := 0; i < length; i++ {
		x := float64(rand.Intn(max-min) + min)
		y := float64(rand.Intn(max-min) + min)
		result = append(result, Point{X: x, Y: y})
	}
	return result
}
