package game

type Vec2 struct {
	X, Y int
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}
