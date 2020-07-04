package life

import (
	"testing"
)

func TestAdvanceCell(t *testing.T) {
	t.Run("BasicBoard Tests", func(t *testing.T) {

		board := BasicBoard{10, 10, make([][]basicCell, 10)}
		for i := range board.cells {
			board.cells[i] = make([]basicCell, 10)
		}
		board.cells[3][2].nextState = true
		board.cells[5][1].state = true
		board.AdvanceCells()
		var expectedResult bool
		for x, column := range board.cells {
			for y, currentCell := range column {
				if x == 3 && y == 2 {
					expectedResult = true
				} else {
					expectedResult = false
				}
				if currentCell.state != expectedResult {
					t.Errorf("Cell %d,%d was %t, expected %t", x, y, currentCell.state, expectedResult)
				}
			}
		}
	})
}

func TestNumberOfLiveNeighbours(t *testing.T) {
	t.Run("BasicBoard Tests", func(t *testing.T) {

		// board := BasicBoard{10, 10, make([][]basicCell, 10)}
		// for i := range board.cells {
		// 	board.cells[i] = make([]basicCell, 10)
		// }

		// Currently because the indexing is [x][y], these nested initialisations are flipped, with x going down and y going across.
		board := BasicBoard{6, 6, [][]basicCell{
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{true, false}, basicCell{true, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{true, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{true, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
		}}

		expectedResult := [][]int{
			{0, 1, 2, 2, 1, 0},
			{1, 2, 2, 1, 1, 0},
			{1, 1, 4, 3, 2, 0},
			{1, 1, 2, 0, 1, 0},
			{0, 0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0, 0},
		}

		for x := 0; x < board.xLimit; x++ {
			for y := 0; y < board.yLimit; y++ {
				result := board.NumberOfLiveNeighbours(x, y)
				if result != expectedResult[x][y] {
					t.Errorf("Cell %d,%d returned %d, expected %d", x, y, result, expectedResult[x][y])
				}
			}
		}
	})
}

func TestSetNextState(t *testing.T) {

	t.Run("BasicBoard Tests", func(t *testing.T) {

		board := BasicBoard{6, 6, [][]basicCell{
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{true, false}, basicCell{true, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{true, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{true, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
			{basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}, basicCell{false, false}},
		}}

		/*
			Number of neighbours is as follows, and * are currently alive
				{0,  1,  2,  2,  1,  0},
				{1,  2,  2*, 1*, 1,  0},
				{1,  1*, 4,  3,  2,  0},
				{1,  1,  2,  0*, 1,  0},
				{0,  0,  1,  1,  1,  0},
				{0,  0,  0,  0,  0,  0},
		*/

		expectedResult := [][]bool{
			{false, false, false, false, false, false},
			{false, false, true, false, false, false},
			{false, false, false, true, false, false},
			{false, false, false, false, false, false},
			{false, false, false, false, false, false},
			{false, false, false, false, false, false},
		}

		board.setNextState()

		for x := 0; x < board.xLimit; x++ {
			for y := 0; y < board.yLimit; y++ {
				if board.cells[x][y].nextState != expectedResult[x][y] {
					t.Errorf("Cell %d,%d has next state of %t, expected %t", x, y, board.cells[x][y].nextState, expectedResult[x][y])
				}
			}
		}

	})
}
