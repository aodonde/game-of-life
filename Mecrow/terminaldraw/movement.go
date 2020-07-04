package terminaldraw

import "fmt"

// Direction - Does this comment count, linter?
type Direction rune

const (
	// Up - ANSI code to move Up includes A
	Up Direction = iota + 'A'
	// Down - ANSI code to move down includes B
	Down
	// Right - ANSI code to move right includes C
	Right
	// Left - ANSI code to move left includes D
	Left
	// Esc - The start of each escape sequence
	Esc = "\u001b["
)

// Cursor - interface for a cursor that can move?
type Cursor interface {
	Move(dir Direction, distance int) bool
	MoveTo(x, y int)
	MoveToStart()
}

// TerminalCursor - struct that uses the CursorMovement interface
type TerminalCursor struct {
	Cursor
}

// Move - current method for moving, though it is currently too concrete.
func (cur TerminalCursor) Move(dir Direction, distance int) bool {

	fmt.Printf(Esc+"%d%c", distance, dir)

	return true
}

// MoveToStart - just a carriage return without a newline
func (cur TerminalCursor) MoveToStart() bool {
	fmt.Print('\r')
	return true
}

// MoveTo - x is the column, y the line
func (cur TerminalCursor) MoveTo(x, y int) {
	fmt.Printf(Esc+"%d;%df", y, x)
}
