package game

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
