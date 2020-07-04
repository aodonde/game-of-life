package main

import (
	"fmt"

	"time"

	term "github.com/aodonde/game-of-life/Mecrow/terminaldraw"
)

func boxOfLetters() {

	var cursor term.TerminalCursor

	for index := 0; index < 10; index++ {
		cursor.Move(term.Right, 1)
		fmt.Printf("R")
		cursor.Move(term.Left, 1)
		time.Sleep(100 * time.Millisecond)
	}
	for index := 0; index < 10; index++ {
		cursor.Move(term.Up, 1)
		fmt.Printf("U")
		cursor.Move(term.Left, 1)
		time.Sleep(100 * time.Millisecond)
	}
	for index := 0; index < 10; index++ {
		cursor.Move(term.Left, 1)
		fmt.Printf("L")
		cursor.Move(term.Left, 1)
		time.Sleep(100 * time.Millisecond)
	}
	for index := 0; index < 10; index++ {
		cursor.Move(term.Down, 1)
		fmt.Printf("D")
		cursor.Move(term.Left, 1)
		time.Sleep(100 * time.Millisecond)
	}
}

func changingString() {
	var cursor term.TerminalCursor

	cursor.PrintInPlace("1st String")
	time.Sleep(2 * time.Second)
	cursor.PrintInPlace("2nd String")
	time.Sleep(2 * time.Second)
	cursor.PrintInPlace("3rd String")
	time.Sleep(2 * time.Second)
	cursor.PrintInPlace("4th")
	time.Sleep(2 * time.Second)
	cursor.PrintInPlace("5")
	time.Sleep(2 * time.Second)
	cursor.PrintInPlace("Final Stirng")
	time.Sleep(2 * time.Second)
	cursor.Move(term.Right, 8)
	cursor.PrintInPlace("rign")
	time.Sleep(2 * time.Second)
	cursor.PrintInPlace("ring, dammit")
	time.Sleep(2 * time.Second)

}

func main() {

	changingString()
}
