package game

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"

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
	EnormousGrid   bool
}

func (g *Game) GetNeighbor(c *Cell, x int, y int) *Cell {
	cX := c.Position.X
	cY := c.Position.Y

	return g.Grid.GetCell(cX+x, cY+y)
}

func (g *Game) GetNumberAliveNeighbors(c *Cell) int {
	var neighbors int

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && g.Grid.GetCell(c.Position.X+i, c.Position.Y+j).Alive {
				neighbors++
			}
		}
	}
	return neighbors
}

func (g *Game) Next(c Cell) bool {
	n := g.GetNumberAliveNeighbors(&c)
	if n == 3 || (n == 2 && c.Alive) {
		return true
	}
	return false
}

func (g *Game) DrawGUI() {
	bS := rg.Slider(
		rl.NewRectangle(100, 100, 100, 20),
		"Brush Size",
		fmt.Sprint(g.BrushSize),
		float32(g.BrushSize),
		1, 20,
	)
	g.BrushSize = int(bS)
	ips := rg.Slider(
		rl.NewRectangle(100, 150, 100, 20),
		"Iterations/Sec",
		fmt.Sprint(g.Speed_IPSecond),
		float32(g.Speed_IPSecond),
		1, 1000,
	)
	g.Speed_IPSecond = int(ips)
	zm := rg.Slider(
		rl.NewRectangle(100, 200, 100, 20),
		"Zoom",
		fmt.Sprintf("%.1f", g.Camera.Zoom),
		float32(g.Camera.Zoom),
		0.1, 20,
	)
	g.Camera.Zoom = zm

	g.EnormousGrid = rg.CheckBox(rl.NewRectangle(100, 300, 20, 20), "Enormous Grid (*10)", g.EnormousGrid)
	maxGrid := 1200
	if g.EnormousGrid {
		maxGrid *= 10
	}

	gW := rg.Slider(
		rl.NewRectangle(100, 250, 100, 20),
		"Grid Size",
		fmt.Sprint(g.Grid.Width),
		float32(g.Grid.Width),
		1, float32(maxGrid),
	)

	if int(gW) != g.Grid.Width {
		g.Paused = true
		g.Grid.Width = int(gW)
		g.Grid.Cells = InitGrid(g.Grid.Width).Cells
		g.Update()
	}
}
