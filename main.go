package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/checkm4ted/no/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Width = 800
const Height = 800

func main() {
	game := utils.Game{
		Width:     Width,
		Height:    Height,
		Grid:      utils.InitGrid(Width/3, Height/3),
		BrushSize: 1,
		Camera: rl.Camera2D{
			Offset:   rl.Vector2{X: 0, Y: 0},
			Target:   rl.Vector2{X: 0, Y: 0},
			Rotation: 0,
			Zoom:     0.3,
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
				time.Sleep(time.Duration(1000/game.Speed_IPSecond) * time.Millisecond)
			}
		}
	}()

	for !rl.WindowShouldClose() {
		go game.HandleInput()

		rl.BeginDrawing()
		rl.BeginMode2D(game.Camera)

		game.Draw()

		rl.EndMode2D()

		rl.DrawText(
			fmt.Sprint(
				"Iterations/Sec: ", game.Speed_IPSecond, "\nFPS: ",
				rl.GetFPS(), "\nZoom: ",
				fmt.Sprintf("%.2f", game.Camera.Zoom),
			),
			10, 10, 20, color.RGBA{255, 255, 255, 255},
		)

		if game.Paused {
			rl.DrawText("PAUSED (Space)", 600, 10, 20, color.RGBA{255, 255, 255, 255})
		}

		rl.EndDrawing()
	}

}
