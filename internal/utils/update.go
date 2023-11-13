package utils

import (
	"image/color"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Update() {
	nextGrid := InitGrid(g.Width/3, g.Height/3)
	for _, c := range g.Grid.Cells {
		nextGrid.GetCell(c.Position.X, c.Position.Y).Alive = g.Next(*c)
	}
	g.Grid.Cells = nextGrid.Cells
}

func (game *Game) HandleInput() {
	if rl.IsKeyPressed(rl.KeySpace) {
		game.Paused = !game.Paused
	}

	if rl.IsKeyPressed(rl.KeyC) {
		for i := range game.Grid.Cells {
			game.Grid.Cells[i].Alive = false
		}
	}

	if rl.IsKeyPressed(rl.KeyR) {
		for i := range game.Grid.Cells {
			game.Grid.Cells[i].Alive = rand.Intn(3) == 1
		}
	}

	if rl.IsKeyPressed(rl.KeyG) {
		game.ShowGrid = !game.ShowGrid
	}

	mW := rl.GetMouseWheelMove()

	if rl.IsMouseButtonDown(rl.MouseMiddleButton) {
		game.Camera.Target = rl.Vector2Subtract(game.Camera.Target, rl.Vector2Scale(rl.GetMouseDelta(), 1/game.Camera.Zoom))
	}

	if rl.IsKeyDown(rl.KeyLeftControl) {
		game.Camera.Zoom += mW * 0.1
		game.Camera.Zoom = ClampMin(game.Camera.Zoom, 0.2)

	} else if rl.IsKeyDown(rl.KeyLeftShift) {
		game.BrushSize += int(mW)
		game.BrushSize = ClampMin(game.BrushSize, 1)
	} else {
		if game.Speed_IPSecond+int(mW) >= 1 {
			game.Speed_IPSecond += int(mW)
		} else {
			game.Speed_IPSecond = 1
		}
	}
}

const cs = 10

func (g *Game) Draw() {
	rl.ClearBackground(color.RGBA{20, 20, 20, 255})

	m := rl.GetScreenToWorld2D(rl.GetMousePosition(), g.Camera)

	for _, c := range g.Grid.Cells {
		c.Color = color.RGBA{0, 0, 0, 255}
		if c.Alive {
			c.Color = color.RGBA{uint8(float64(g.GetNumberAliveNeighbors(*c)) / 8.0 * 255), 127, 0, 255}
		}

		x := c.Position.X
		y := c.Position.Y

		xS := x * cs
		yS := y * cs

		brushSize := g.BrushSize / 2.0 * cs

		highlighted := m.X >= float32(xS-brushSize) &&
			m.X <= float32(xS+cs+brushSize-1) &&
			m.Y >= float32(yS-brushSize) &&
			m.Y <= float32(yS+cs+brushSize-1)

		if highlighted {
			c.Color = color.RGBA{255, 255, 255, 255}
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				c.Alive = true
			} else if rl.IsMouseButtonDown(rl.MouseRightButton) {
				c.Alive = false
			}
		}

		if g.ShowGrid {
			rl.DrawRectangle(
				int32(xS),
				int32(yS),
				int32(cs-1),
				int32(cs-1),
				c.Color,
			)
			continue
		}

		// If there's no grid, only render the alive cells (much better performance)
		if c.Alive || highlighted {
			rl.DrawRectangle(
				int32(xS),
				int32(yS),
				int32(cs),
				int32(cs),
				c.Color,
			)
		}
	}
}
