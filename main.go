package main

import (
	"fmt"
	"image/color"
	"time"

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
			Offset:   rl.Vector2{X: 350, Y: 0},
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

	rl.SetTargetFPS(300)

	go func() {
		for {
			if !game.Paused {
				game.Update()
				time.Sleep(time.Duration(1_000_000_000/game.Speed_IPSecond) * time.Nanosecond)
			}
		}
	}()

	for !rl.WindowShouldClose() {
		go game.HandleInput()

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
