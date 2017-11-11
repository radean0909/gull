package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/fogleman/gg"
	"github.com/georgea93/gull/hull"
	"github.com/georgea93/gull/point"
)

func generate(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	points := point.RandomPoints(400, 100, 1900)
	convexHull := hull.FromPoints(points)

	graphics := gg.NewContext(2000, 2000)
	graphics.DrawRectangle(0, 0, 2000, 2000)
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
