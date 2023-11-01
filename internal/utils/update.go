package utils

func (g *Game) Update() {
	for x := range g.Grid {
		for y := range g.Grid[x] {
			c := g.GetCell(x, y)
			if g.GetNumberAliveNeighbors(c) <= 3 {
				c.Alive = !c.Alive
			}
		}
	}
}
