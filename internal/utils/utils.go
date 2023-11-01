package utils

import (
	"image/color"
)

type Vec2 struct {
	X, Y int
}

type Cell struct {
	Alive       bool
	Position    Vec2
	Color       color.RGBA
	JustChanged bool
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

func InitGrid(Width int, Height int, CellSize int) [][]*Cell {
	grid := make([][]*Cell, Width/CellSize)

	for x := range grid {
		grid[x] = make([]*Cell, Height/CellSize)
		for y := range grid[x] {
			grid[x][y] = &Cell{
				Alive:    false,
				Position: NewVec2(x, y),
			}
		}
	}

	return grid
}
