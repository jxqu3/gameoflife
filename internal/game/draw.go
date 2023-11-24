package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const CellSize int = 10

func (game *Game) Draw() {
	rl.ClearBackground(color.RGBA{20, 20, 20, 255})

	m := rl.GetScreenToWorld2D(rl.GetMousePosition(), game.Camera)

	var cColor color.RGBA
	fCs := float32(CellSize)
	camWorldPos := rl.GetScreenToWorld2D(rl.Vector2{X: float32(game.Width) + 2*fCs, Y: float32(game.Height) + 2*fCs}, game.Camera)
	startWorldPos := rl.GetScreenToWorld2D(rl.Vector2{X: 250 - 2*fCs*game.Camera.Zoom, Y: 0. - 2*fCs*game.Camera.Zoom}, game.Camera)

	var visible bool
	brushSize := int(game.BrushSize/2.0) * CellSize

	for _, c := range game.Grid.Cells {
		visible = float32(c.XS) < camWorldPos.X && float32(c.YS) < camWorldPos.Y && float32(c.XS) > startWorldPos.X && float32(c.YS) > startWorldPos.Y

		if !visible {
			continue
		}

		if c.Alive {
			cColor = color.RGBA{uint8(float64(game.Grid.GetNumberAliveNeighbors(*c)) / 8.0 * 255), 127, 0, 255}
		} else {
			cColor = color.RGBA{0, 0, 0, 255}
		}

		highlighted := visible &&
			int(m.X) >= c.XS-brushSize+1 &&
			int(m.X) <= c.XS+CellSize+brushSize &&
			int(m.Y) >= c.YS-brushSize+1 &&
			int(m.Y) <= c.YS+CellSize+brushSize

		if highlighted {
			cColor = color.RGBA{255, 255, 255, 255}
			switch {
			case rl.IsMouseButtonDown(rl.MouseLeftButton):
				c.Alive = true
			case rl.IsMouseButtonDown(rl.MouseRightButton):
				c.Alive = false
			}
		}

		if game.ShowGrid {
			rl.DrawRectangle(
				int32(c.XS),
				int32(c.YS),
				int32(CellSize-1),
				int32(CellSize-1),
				cColor,
			)
			continue
		}

		// If there's no grid, only render the alive cells (much better performance)
		if c.Alive || highlighted {
			rl.DrawRectangle(
				int32(c.XS),
				int32(c.YS),
				int32(CellSize),
				int32(CellSize),
				cColor,
			)
		}
	}
}
