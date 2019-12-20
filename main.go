package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/radean0909/gull/point"

	"github.com/fogleman/gg"
	"github.com/radean0909/gull/hull"
)

var width = 2000
var height = 2000
var centreX = float64(width / 2)
var centreY = float64(height / 2)
var maxPoints = 10000
var minPoints = 10
var maxRadius = 900
var minRadius = 10
var numNoisePoints = 5
var minNoisePoint = 10
var maxNoisePoint = width - 10

func generate(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	// generate a circle of points
	numPoints := rand.Intn(maxPoints-minPoints) + minPoints
	radius := rand.Intn(maxRadius-minRadius) + minRadius
	points := point.RandomPointsCircle(numPoints, float64(radius), centreX, centreY)
	// add a little noise
	points = append(points, point.RandomPoints(numNoisePoints, minNoisePoint, maxNoisePoint)...)
	convexHull := hull.FromPoints(points)

	graphics := gg.NewContext(width, height)
	graphics.DrawRectangle(0, 0, float64(width), float64(height))
	graphics.SetRGB(1, 1, 1)
	graphics.Fill()
	points.Draw(graphics)
	convexHull.Draw(graphics)
	path := fmt.Sprintf("results/output_%d.png", i)
	graphics.SavePNG(path)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go generate(&wg, i)
	}

	wg.Wait()
}
