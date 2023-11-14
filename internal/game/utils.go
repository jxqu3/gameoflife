package game

type Vec2 struct {
	X, Y int
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
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
