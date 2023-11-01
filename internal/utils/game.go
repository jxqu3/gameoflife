package utils

type Game struct {
	Width    int
	Height   int
	CellSize int
	Grid     [][]*Cell
}

func (g *Game) GetNeighbor(c Cell, x int, y int) *Cell {
	cX := c.Position.X
	cY := c.Position.Y

	if cX+x < 0 || cX+x >= g.Width {
		return g.Grid[0][0]
	}
	if cY+y < 0 || cY+y >= g.Height {
		return g.Grid[0][0]
	}

	return g.Grid[cX+x][cY+y]
}
