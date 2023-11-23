package game

import (
	"image/color"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Grid) Update() {
	nextGrid := InitGrid(g.Width)
	for _, c := range g.Cells {
		nextGrid.GetCell(c.Position.X, c.Position.Y).Alive = g.Next(*c)
	}
	g.Cells = nextGrid.Cells
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

	if rl.IsMouseButtonDown(rl.MouseMiddleButton) {
		game.Camera.Target = rl.Vector2Subtract(game.Camera.Target, rl.Vector2Scale(rl.GetMouseDelta(), 1/game.Camera.Zoom))
	}

}

func (game *Game) HandleControls() {
	mW := rl.GetMouseWheelMove()
	if rl.IsKeyDown(rl.KeyLeftControl) {
		game.Camera.Zoom += mW * 0.1
	} else if rl.IsKeyDown(rl.KeyLeftShift) {
		game.BrushSize += int(mW)
		game.BrushSize = max(game.BrushSize, 1)
	} else if rl.IsKeyDown(rl.KeyLeftAlt) {
		if int(mW) != game.Grid.Width {
			game.Paused = true
			game.Grid.Width = max(game.Grid.Width+int(mW*10), 10)
			game.Grid.Cells = InitGrid(game.Grid.Width).Cells
			game.Grid.Update()
		}
	} else {
		if game.Speed_IPSecond+int(mW) >= 1 {
			game.Speed_IPSecond += int(mW)
		} else {
			game.Speed_IPSecond = 1
		}
	}
}

const cs = 10

func (game *Game) Draw() {
	rl.ClearBackground(color.RGBA{20, 20, 20, 255})

	m := rl.GetScreenToWorld2D(rl.GetMousePosition(), game.Camera)

	var cColor color.RGBA
	camWorldPos := rl.GetScreenToWorld2D(rl.Vector2{X: float32(game.Width) + 2*cs, Y: float32(game.Height) + 2*cs}, game.Camera)
	startWorldPos := rl.GetScreenToWorld2D(rl.Vector2{X: 250 - 2*cs*game.Camera.Zoom, Y: 0. - 2*cs*game.Camera.Zoom}, game.Camera)

	var x int
	var y int
	var xS int
	var yS int
	var visible bool
	brushSize := int(game.BrushSize / 2.0 * cs)

	for _, c := range game.Grid.Cells {
		x = c.Position.X
		y = c.Position.Y
		xS = x * cs
		yS = y * cs

		cellWorldPos := rl.Vector2{X: float32(xS), Y: float32(yS)}
		visible = cellWorldPos.X < camWorldPos.X && cellWorldPos.Y < camWorldPos.Y && cellWorldPos.X > startWorldPos.X && cellWorldPos.Y > startWorldPos.Y

		if !visible {
			continue
		}

		if c.Alive {
			cColor = color.RGBA{uint8(float64(game.Grid.GetNumberAliveNeighbors(c)) / 8.0 * 255), 127, 0, 255}
		} else {
			cColor = color.RGBA{0, 0, 0, 255}
		}

		highlighted := visible &&
			int(m.X) >= xS-brushSize+1 &&
			int(m.X) <= xS+cs+brushSize &&
			int(m.Y) >= yS-brushSize+1 &&
			int(m.Y) <= yS+cs+brushSize

		if highlighted {
			cColor = color.RGBA{255, 255, 255, 255}
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				c.Alive = true
			} else if rl.IsMouseButtonDown(rl.MouseRightButton) {
				c.Alive = false
			}
		}

		if game.ShowGrid {
			rl.DrawRectangle(
				int32(xS),
				int32(yS),
				int32(cs-1),
				int32(cs-1),
				cColor,
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
				cColor,
			)
		}
	}
}
