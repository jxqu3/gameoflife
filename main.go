package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/checkm4ted/gameoflife/internal/utils"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Width = 1200
const Height = 800

func main() {
	game := utils.Game{
		Width:     Width,
		Height:    Height,
		Grid:      utils.InitGrid(80, 80),
		BrushSize: 1,
		Camera: rl.Camera2D{
			Offset:   rl.Vector2{X: 300, Y: 0},
			Target:   rl.Vector2{X: 0, Y: 0},
			Rotation: 0,
			Zoom:     0.4,
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
			game.Grid.Width = int(gW)
			game.Grid = utils.InitGrid(int(gW), int(gW))
		}

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
