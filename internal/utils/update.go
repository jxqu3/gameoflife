package utils

func (g *Game) Update() {
	for i := range g.Grid {
		for j := range g.Grid[i] {
			c := g.Grid[i][j]
			if g.GetNumberAliveNeighbors(c) <= 3 {
				c.Alive = !c.Alive
			}
		}
	}
}
