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
}

func (g *Game) GetNeighbor(c *Cell, x int, y int) *Cell {
	cX := c.Position.X
	cY := c.Position.Y

	return g.Grid.GetCell(cX+x, cY+y)
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

func (game *Game) DrawGUI() {
	bS := rg.Slider(
		rl.NewRectangle(100, 100, 100, 20),
		"Brush Size",
		fmt.Sprint(game.BrushSize),
		float32(game.BrushSize),
		1, 20,
	)
	game.BrushSize = int(bS)
	ips := rg.Slider(
		rl.NewRectangle(100, 150, 100, 20),
		"Iterations/Sec",
		fmt.Sprint(game.Speed_IPSecond),
		float32(game.Speed_IPSecond),
		1, 1000,
	)
	game.Speed_IPSecond = int(ips)
	zm := rg.Slider(
		rl.NewRectangle(100, 200, 100, 20),
		"Zoom",
		fmt.Sprintf("%.1f", game.Camera.Zoom),
		float32(game.Camera.Zoom),
		0.1, 20,
	)
	game.Camera.Zoom = zm
	gW := rg.Slider(
		rl.NewRectangle(100, 250, 100, 20),
		"Grid Size",
		fmt.Sprint(game.Grid.Width),
		float32(game.Grid.Width),
		1, 1000,
	)

	if int(gW) != game.Grid.Width {
		game.Paused = true
		game.Grid.Width = int(gW)
		game.Grid.Cells = InitGrid(game.Grid.Width).Cells
		game.Update()
	}
}
