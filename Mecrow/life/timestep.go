package life

// This stuff should probably be in life.go or something
/*
const (
	// Unknown - Cell hasn't been updated yet
	Unknown = iota
	// Alive - Cell is alive
	Alive
	// Dead - Cell is dead...
	Dead
)
type cell struct {
	state     int
	nextState int
}
*/

type basicCell struct {
	state     bool
	nextState bool
}

// Board is a general interface for different implementations of the game "board"
type Board interface {
	AdvanceCells()
	NumberOfLiveNeighbours(x, y int) int
}

// BasicBoard - Just do everything on every cell each timestep
type BasicBoard struct {
	xLimit, yLimit int
	cells          [][]basicCell
}

// AdvanceCells assigns the nextState to state, and Unknown to nextState
func (b BasicBoard) AdvanceCells() {
	for x := 0; x < b.xLimit; x++ {
		for y := 0; y < b.yLimit; y++ {
			b.cells[x][y].state, b.cells[x][y].nextState = b.cells[x][y].nextState, false
		}
	}
}

// NumberOfLiveNeighbours returns the number of live cells around the x, y cell coordinate given.
func (b BasicBoard) NumberOfLiveNeighbours(x, y int) int {
	neighbourCount := 0
	validX := make([]int, 0, 3)
	validY := make([]int, 0, 3)
	if x > 0 {
		validX = append(validX, x-1)
	}
	validX = append(validX, x)
	if x < (b.xLimit - 1) {
		validX = append(validX, x+1)
	}

	if y > 0 {
		validY = append(validY, y-1)
	}
	validY = append(validY, y)
	if y < (b.yLimit - 1) {
		validY = append(validY, y+1)
	}

	for _, neighbourX := range validX {
		for _, neighbourY := range validY {
			if (neighbourX != x) || (neighbourY != y) {
				if b.cells[neighbourX][neighbourY].state {
					neighbourCount++
				}
			}
		}
	}
	return neighbourCount

}

func (b BasicBoard) setNextState() {

	for x := 0; x < b.xLimit; x++ {
		for y := 0; y < b.yLimit; y++ {
			n := b.NumberOfLiveNeighbours(x, y)
			if b.cells[x][y].state {
				if n == 2 || n == 3 {
					b.cells[x][y].nextState = true
				} else {
					b.cells[x][y].nextState = false
				}
			} else {
				if n == 3 {
					b.cells[x][y].nextState = true
				} else {
					b.cells[x][y].nextState = false
				}
			}
		}
	}
}
