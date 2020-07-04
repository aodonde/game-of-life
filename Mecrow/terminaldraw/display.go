package terminaldraw

import "fmt"

// Display - interface for a display
type Display interface {
	ClearScreen()
	UpdateDisplay()
}

const (
	// Box is the ascii for a black box that should fill the cell
	Box = "█"
)

// TerminalDisplay the terminal implementation of a display
type TerminalDisplay struct {
	Display
}

// Print - Print a string
func (cons Console) Print(output string) (int, error) {
	return fmt.Print(output)
}

// PrintInPlace - Print a string but then move back to the old cursor location
func (cons Console) PrintInPlace(output string) {
	amountPrinted, _ := fmt.Print(output)
	cons.Move(Left, amountPrinted)
}

// Console - struct that uses all of the interfaces?
type Console struct {
	TerminalDisplay
	TerminalCursor
}

// ClearScreen - clear the whole screen
func (disp TerminalDisplay) ClearScreen() {
	fmt.Print(Esc + "2J")
}
