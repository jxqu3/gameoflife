package utils

func (g *Game) Update() {
	g.NextGrid.Cells = InitGrid(g.Width, g.Height, 5)
	for x := range g.Grid.Cells {
		for y := range g.Grid.Cells[x] {
			g.NextGrid.Cells[x][y].Alive = g.Next(x, y)
		}
	}
	g.Grid, g.NextGrid = g.NextGrid, g.Grid
}
