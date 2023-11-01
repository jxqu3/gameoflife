package utils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Draw() {
	cs := g.CellSize
	g.Update()
	rl.ClearBackground(color.RGBA{0, 0, 0, 255})
	for i := range g.Grid {
		for j, c := range g.Grid[i] {
			cellColor := color.RGBA{0, 0, 0, 255}
			if c.Alive {
				cellColor = color.RGBA{255, 255, 255, 255}
			}
			rl.DrawRectangle(int32(i*cs), int32(j*cs), int32(cs), int32(cs), cellColor)
		}
	}
}
