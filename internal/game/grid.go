package game

type Grid struct {
	Width  int
	Height int
	Cells  []*Cell
}

func (g *Grid) GetCell(x, y int) *Cell {
	x = (g.Width + x) % g.Width
	y = (g.Height + y) % g.Height
	return g.Cells[x+y*g.Width]
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
