package utils

import "image/color"

type Vec2 struct {
	X, Y int
}

type Cell struct {
	Alive    bool
	Position Vec2
	Color    color.RGBA
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}
