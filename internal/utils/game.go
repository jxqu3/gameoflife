package utils

type Game struct {
	Width        int
	Height       int
	CellSize     int
	InitCellSize int
	Grid         Grid
	NextGrid     Grid
}

type Grid struct {
	Width  int
	Height int
	Cells  [][]*Cell
}

func (g *Game) GetNeighbor(c *Cell, x int, y int) *Cell {
	cX := c.Position.X
	cY := c.Position.Y

	return g.GetCell(cX+x, cY+y)
}

func (g *Game) GetCell(x, y int) *Cell {
	y = (g.Grid.Height + y) % g.Grid.Height
	x = (g.Grid.Width + x) % g.Grid.Width
	return g.Grid.Cells[x][y]
}

func (g *Game) GetNumberAliveNeighbors(c *Cell) int {
	var neighbors int

	y := c.Position.Y
	x := c.Position.X

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && g.GetCell(x+i, y+j).Alive {
				neighbors++
			}
		}
	}
	return neighbors
}

func (g *Game) Next(x, y int) bool {
	c := g.GetCell(x, y)
	n := g.GetNumberAliveNeighbors(c)

	if n == 3 || (n == 2 && c.Alive) {
		return true
	}
	return false
}
