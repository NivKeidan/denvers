package world

const (
	Grass = '░'
	Path  = ' '
	Wall  = '#'
)

type Map struct {
	Tiles [][]rune
}

func (m *Map) CanWalk(pos Position) bool {
	if pos.y < 0 || pos.y >= len(m.Tiles) || pos.x < 0 || pos.x >= len(m.Tiles[0]) {
		return false // Out of bounds
	}
	return m.Tiles[pos.y][pos.x] == Path // Only walk on paths
}

func DefaultMap() (*Map, *Position) {
	height := 7
	width := 50

	// Generate empty map
	tiles := make([][]rune, height)
	for i := range tiles {
		tiles[i] = make([]rune, width)
	}

	// Populate
	world := []rune{Wall, Grass, Grass, Path, Grass, Grass, Wall}
	for i, row := range tiles {
		for j := range row {
			row[j] = world[i]
		}
	}
	return &Map{Tiles: tiles}, &Position{x: 0, y: 3}
}
