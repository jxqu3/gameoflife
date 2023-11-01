package main

import (

	// Lib to replace err != nil

	"math"

	ut "github.com/checkm4ted/va/internal/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const Width = 800
const Height = 800
const CellSize = 10

// Iterations Per Second
const Speed_IPSecond = 10

func main() {
	grid := make([][]*ut.Cell, Width/CellSize)
	for x := 0; x < Width/CellSize; x++ {
		grid[x] = make([]*ut.Cell, Height/CellSize)
		for y := 0; y < Width/CellSize; y++ {
			grid[x][y] = &ut.Cell{
				Alive:    math.Mod(float64(x), 2) == 0,
				Position: ut.NewVec2(x, y),
			}
		}
	}

	game := ut.Game{
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
