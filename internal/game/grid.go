package game

type Grid struct {
	Width int
	Cells []*Cell
}

type Vec2 struct {
	X int
	Y int
}

func (g *Grid) GetCell(x, y int) *Cell {
	cx := (g.Width + x) % g.Width
	cy := (g.Width + y) % g.Width
	return g.Cells[cx+cy*g.Width]
}

func InitGrid(Width int) Grid {
	cells := make([]*Cell, Width*Width)
	for i := range cells {
		x := i % Width
		y := i / Width
		cells[i] = &Cell{
			Alive: false,
			X:     x,
			XS:    x * CellSize,
			Y:     y,
			YS:    y * CellSize,
		}
	}

	return Grid{
		Width: Width,
		Cells: cells,
	}
}

func (g *Grid) Update() {
	nextGrid := InitGrid(g.Width)
	for _, c := range g.Cells {
		nextGrid.GetCell(c.X, c.Y).Alive = g.Next(*c)
	}
	g.Cells = nextGrid.Cells
}
