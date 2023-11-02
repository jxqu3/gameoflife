package main

import (

	// Lib to replace err != nil

	"fmt"
	"image/color"
	"time"

	"github.com/checkm4ted/no/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Width = 800
const Height = 800
const CellSize = 5

// Iterations Per Second
var Speed_IPSecond = 5

var Paused = true

func main() {

	grid := utils.InitGrid(Width, Height, CellSize)

	game := utils.Game{
		Width:        Width,
		Height:       Height,
		CellSize:     CellSize,
		InitCellSize: CellSize,
		Grid: utils.Grid{
			Width:  Width / CellSize,
			Height: Height / CellSize,
			Cells:  grid,
		},
		NextGrid: utils.Grid{
			Width:  Width / CellSize,
			Height: Height / CellSize,
			Cells:  grid,
		},
	}
	rl.InitWindow(Width, Height, "CheckM4te Game Of Life")
	defer rl.CloseWindow()
	rl.SetTargetFPS(144)

	go func() {
		for {
			if !Paused {
				game.Update()
				time.Sleep(time.Duration(1000/Speed_IPSecond) * time.Millisecond)
			}
		}
	}()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			Paused = !Paused
		}
		rl.BeginDrawing()

		game.Draw()

		rl.DrawText(fmt.Sprint("Iterations/Sec: ", Speed_IPSecond), 10, 10, 20, color.RGBA{255, 255, 255, 255})

		if Paused {
			rl.DrawText("PAUSED (Space)", 600, 10, 20, color.RGBA{255, 255, 255, 255})
		}

		mW := rl.GetMouseWheelMove()

		if rl.IsKeyDown(rl.KeyLeftControl) {
			game.CellSize += int(mW)
		} else {
			if Speed_IPSecond+int(mW) >= 1 {
				Speed_IPSecond += int(mW)
			} else {
				Speed_IPSecond = 1
			}
		}

		rl.EndDrawing()
	}

}
