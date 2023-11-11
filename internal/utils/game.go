package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Width          int
	Height         int
	Grid           Grid
	BrushSize      int
	Camera         rl.Camera2D
	ShowGrid       bool
	Speed_IPSecond int
	Paused         bool
}

type Grid struct {
	Width  int
	Height int
	Cells  []*Cell
}

func (g *Game) GetNeighbor(c *Cell, x int, y int) *Cell {
	cX := c.Position.X
	cY := c.Position.Y

	return g.Grid.GetCell(cX+x, cY+y)
}

func (g *Grid) GetCell(x, y int) *Cell {
	x = (g.Width + x) % g.Width
	y = (g.Height + y) % g.Height
	return g.Cells[x+y*g.Width]
}

func (g *Game) GetNumberAliveNeighbors(c Cell) int {
	var neighbors int

	y := c.Position.Y
	x := c.Position.X

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && g.Grid.GetCell(x+i, y+j).Alive {
				neighbors++
			}
		}
	}
	return neighbors
}

func (g *Game) Next(c Cell) bool {
	n := g.GetNumberAliveNeighbors(c)

	if n == 3 || (n == 2 && c.Alive) {
		return true
	}
	return false
}
