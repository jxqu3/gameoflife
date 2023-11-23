package main

import (
	"fmt"
	"image/color"

	g "github.com/checkm4ted/gameoflife/internal/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Width = 1200
const Height = 800

func main() {
	game := g.Game{
		Width:     Width,
		Height:    Height,
		Grid:      g.InitGrid(100),
		BrushSize: 1,
		Camera: rl.Camera2D{
			Offset:   rl.Vector2{X: 250, Y: 0},
			Target:   rl.Vector2{X: 0, Y: 0},
			Rotation: 0,
			Zoom:     1,
		},
		ShowGrid:       true,
		Paused:         true,
		Speed_IPSecond: 1,
	}

	rl.InitWindow(Width, Height, "CheckM4te Game Of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(360)

	t := 0.0
	go func() {
		for {
			if !game.Paused && t >= 1./float64(game.Speed_IPSecond) {
				game.Grid.Update()
				t = 0
			}
		}
	}()

	for !rl.WindowShouldClose() {
		t += float64(rl.GetFrameTime())
		go game.HandleInput()
		game.HandleControls()

		rl.BeginDrawing()
		rl.BeginMode2D(game.Camera)

		game.Draw()

		rl.EndMode2D()

		rl.DrawRectangle(0, 0, 250, Height, color.RGBA{220, 220, 220, 255})

		game.DrawGUI()

		rl.DrawText(
			fmt.Sprint(
				"FPS: ", rl.GetFPS(),
			),
			10, 10, 20, color.RGBA{0, 0, 0, 255},
		)

		if game.Paused {
			rl.DrawText("PAUSED (Space)", 1020, 10, 20, color.RGBA{255, 255, 255, 255})
		}

		rl.EndDrawing()
	}

}
