package utils

func (g *Game) Update() {
	for x := range g.Grid.Cells {
		for y := range g.Grid.Cells[x] {
			g.NextGrid.Cells[x][y].Alive = g.Next(x, y)
		}
	}
	g.Grid.Cells, g.NextGrid.Cells = g.NextGrid.Cells, InitGrid(g.Width, g.Height, g.InitCellSize)
}
