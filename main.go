package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/georgea93/gull/hull"
	"github.com/georgea93/gull/point"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	points := point.RandomPoints(400, 100, 1900)
	convexHull := hull.FromPoints(points)

	graphics := gg.NewContext(2000, 2000)
	graphics.DrawRectangle(0, 0, 2000, 2000)
	graphics.SetRGB(1, 1, 1)
	graphics.Fill()
	points.Draw(graphics)
	convexHull.Draw(graphics)
	graphics.SavePNG("output.png")
}
