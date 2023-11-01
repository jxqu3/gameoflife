package utils

import (
	"image/color"
	"math"
)

func (g *Game) Update() {
	for x := range g.Grid {
		for y := range g.Grid[x] {
			c := g.GetCell(x, y)
			c.Color = color.RGBA{0, 0, 0, 255}
			if c.Alive {
				if math.Mod(float64(x), 2) == 0 {
					c.Color = color.RGBA{0, 255, 0, 255}
				} else {
					c.Color = color.RGBA{255, 255, 255, 255}
				}
			}
			if g.GetNumberAliveNeighbors(c) <= 3 {
				c.Alive = !c.Alive
			}
		}
	}
}
