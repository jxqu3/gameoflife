package main

import (

	// Lib to replace err != nil

	"github.com/checkm4ted/va/internal/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Width = 800
const Height = 800
const CellSize = 10

// Iterations Per Second
const Speed_IPSecond = 10

func main() {

	grid := utils.InitGrid(Width, Height, CellSize)

	game := utils.Game{
		Width:    Width / CellSize,
		Height:   Height / CellSize,
		CellSize: CellSize,
		Grid:     grid,
	}
	rl.InitWindow(Width, Height, "CheckM4te Automata")
	defer rl.CloseWindow()
	rl.SetTargetFPS(Speed_IPSecond)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		game.Draw()

		rl.EndDrawing()
	}
}
