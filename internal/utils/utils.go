package utils

import (
	"image/color"
)

type Vec2 struct {
	X, Y int
}

type Cell struct {
	Alive       bool
	Position    Vec2
	Color       color.RGBA
	JustChanged bool
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

func InitGrid(Width int, Height int) Grid {
	cells := make([]*Cell, Width*Height)
	for i := range cells {
		x := (i % Width)
		y := (i / Height)
		cells[i] = &Cell{
			Alive:    false,
			Position: NewVec2(x, y),
		}
	}

	return Grid{
		Width:  Width,
		Height: Height,
		Cells:  cells,
	}
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func ClampMin[T Number](x, min T) T {
	if x < min {
		return min
	}
	return x
}
