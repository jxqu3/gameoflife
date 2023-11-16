package game

type Grid struct {
	Width int
	Cells []*Cell
}

func (g *Grid) GetCell(x, y int) *Cell {
	x = (g.Width + x) % g.Width
	y = (g.Width + y) % g.Width
	return g.Cells[x+y*g.Width]
}

func InitGrid(Width int) Grid {
	cells := make([]*Cell, Width*Width)
	for i := range cells {
		x := (i % Width)
		y := (i / Width)
		cells[i] = &Cell{
			Alive:    false,
			Position: NewVec2(x, y),
		}
	}

	return Grid{
		Width: Width,
		Cells: cells,
	}
}
