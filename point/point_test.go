package point

import "testing"
import "github.com/stretchr/testify/assert"

var crossProductTests = []struct {
	a   Point
	b   Point
	c   Point
	out float64
}{
	{
		Point{X: 10, Y: 12},
		Point{X: 40, Y: 2},
		Point{X: 34, Y: 34},
		900,
	},
	{
		Point{},
		Point{},
		Point{},
		0,
	},
}

func TestCrossProduct(t *testing.T) {
	for _, tt := range crossProductTests {
		result := CrossProduct(tt.a, tt.b, tt.c)
		assert.Equal(t, tt.out, result)
	}
}

var randomPointTests = []struct {
	length, min, max int
}{
	{
		10,
		1,
		9,
	},
	{
		0,
		1,
		9,
	},
}

func TestRandomPoints(t *testing.T) {
	for _, tt := range randomPointTests {
		result := RandomPoints(tt.length, tt.min, tt.max)
		assert.Equal(t, tt.length, len(result))
		for _, p := range result {
			inRange := func() bool {
				return p.X >= float64(tt.min) &&
					p.X <= float64(tt.max) &&
					p.Y >= float64(tt.min) &&
					p.Y <= float64(tt.max)
			}
			assert.Condition(t, inRange)
		}
	}
}
