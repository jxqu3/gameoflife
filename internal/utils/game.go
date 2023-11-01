package utils

type Game struct {
	Width    int
	Height   int
	CellSize int
	Grid     [][]*Cell
}

func (g *Game) GetNeighbor(c *Cell, x int, y int) *Cell {
	cX := c.Position.X
	cY := c.Position.Y

	if cX+x < 0 || cX+x >= g.Width {
		return c
	}
	if cY+y < 0 || cY+y >= g.Height {
		return c
	}

	return g.GetCell(cX+x, cY+y)
}

func (g *Game) Exists(x, y int) bool {
	return x >= 0 && x < len(g.Grid) && y >= 0 && y < len(g.Grid[x])
}

func (g *Game) GetCell(x, y int) *Cell {
	if g.Exists(x, y) {
		return g.Grid[x][y]
	}
	return g.GetCell(0, 0)
}

func (g *Game) GetNumberAliveNeighbors(c *Cell) int {
	var neighbors int

	y := c.Position.Y
	x := c.Position.X

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if g.GetCell(j, i).Alive {
				neighbors++
			}
		}
	}
	return neighbors
}
