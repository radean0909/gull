package hull

import (
	"sort"

	"github.com/fogleman/gg"
	"github.com/georgea93/gull/point"
)

type Hull struct {
	points point.Points
}

func (hull Hull) Draw(graphics *gg.Context) {
	size := len(hull.points)
	for i, p := range hull.points {
		p.Draw(graphics)
		graphics.SetLineWidth(2)
		if i+1 < size {
			graphics.DrawLine(p.X, p.Y, hull.points[i+1].X, hull.points[i+1].Y)
		} else {
			// Make sure we close the loop!
			graphics.DrawLine(p.X, p.Y, hull.points[0].X, hull.points[0].Y)
		}
		graphics.Stroke()
	}
}

func lowerHull(points point.Points, size int) chan point.Points {
	result := make(chan point.Points)
	go func() {
		var lower point.Points
		count := 0
		for i := 0; i < size; i++ {
			for count >= 2 && point.CrossProduct(lower[count-2], lower[count-1], points[i]) <= 0 {
				count--
				lower = lower[:count]
			}
			count++
			lower = append(lower, points[i])
		}
		result <- lower[:len(lower)-1]
	}()
	return result
}

func upperHull(points point.Points, size int) chan point.Points {
	result := make(chan point.Points)
	go func() {
		var upper point.Points
		count := 0
		for i := size - 1; i >= 0; i-- {
			for count >= 2 && point.CrossProduct(upper[count-2], upper[count-1], points[i]) <= 0 {
				count--
				upper = upper[:count]
			}
			count++
			upper = append(upper, points[i])
		}
		result <- upper[:len(upper)-1]
	}()
	return result
}

func convexHull(points point.Points) *Hull {
	var result point.Points
	size := len(points)
	if size < 3 {
		return &Hull{result}
	}

	sort.Sort(points)

	lower := <-lowerHull(points, size)
	upper := <-upperHull(points, size)

	result = append(result, lower...)
	result = append(result, upper...)
	return &Hull{result}
}

func FromPoints(points point.Points) *Hull {
	return convexHull(points)
}
