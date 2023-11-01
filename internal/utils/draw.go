package utils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Draw() {
	cs := g.CellSize
	g.Update()
	rl.ClearBackground(color.RGBA{0, 0, 0, 255})
	for x := range g.Grid {
		for y := range g.Grid[x] {
			cellColor := color.RGBA{0, 0, 0, 255}
			if g.GetCell(x, y).Alive {
				cellColor = color.RGBA{255, 255, 255, 255}
			}
			rl.DrawRectangle(int32(x*cs), int32(y*cs), int32(cs), int32(cs), cellColor)
		}
	}
}
