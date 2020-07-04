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
)

// CursorMovement - interface for a cursor that can move?
type CursorMovement interface {
	Move(dir Direction, distance int) bool

	MoveToStart()
}

// TerminalCursor - struct that uses the CursorMovement interface
type TerminalCursor struct {

}

// PrintInPlace - Print a string but then move back to the old cursor location
func (cur TerminalCursor) PrintInPlace(output string) {
	amountPrinted, _ := fmt.Print(output)
	cur.Move(Left, amountPrinted)
}

// Move - current method for moving, though it is currently too concrete.
func (cur TerminalCursor) Move(dir Direction, distance int) bool {

	fmt.Printf("\u001b[%d%c", distance, dir)

	return true
}
