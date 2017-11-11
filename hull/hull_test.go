package hull

import (
	"testing"

	"github.com/georgea93/gull/point"
	"github.com/stretchr/testify/assert"
)

var hullTests = []struct {
	in  point.Points
	out Hull
}{
	{
		nil,
		Hull{},
	},
	{
		point.Points{
			{X: 0, Y: 3},
			{X: 2, Y: 2},
		},
		Hull{},
	},
	{
		point.Points{
			{X: 0, Y: 3},
			{X: 2, Y: 2},
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 3, Y: 0},
			{X: 0, Y: 0},
			{X: 3, Y: 3},
		},
		Hull{
			point.Points{
				{X: 0, Y: 0},
				{X: 3, Y: 0},
				{X: 3, Y: 3},
				{X: 0, Y: 3},
			},
		},
	},
}

func TestFromPoints(t *testing.T) {
	for _, tt := range hullTests {
		result := FromPoints(tt.in)
		assert.Equal(t, tt.out.points, result.points)
	}
}
