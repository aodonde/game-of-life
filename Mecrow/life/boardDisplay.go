package life

import term "github.com/aodonde/game-of-life/Mecrow/terminaldraw"

const (
	xOffset = 1
	yOffset = 1
)

// TerminalGame is a type that needs a description
type TerminalGame struct {
	term.Console
	BasicBoard
}

// Setup clears the screen etc
func (game TerminalGame) Setup() {

	game.ClearScreen()

}

// MoveTo should hopefully "override" the console MoveTo with including the offset etc
func (game TerminalGame) MoveTo(x, y int) {
	game.Console.MoveTo(x+xOffset, y+yOffset)
}
