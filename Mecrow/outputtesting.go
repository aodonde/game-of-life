package main

import (
	"fmt"

	"time"

	term "github.com/aodonde/game-of-life/Mecrow/terminaldraw"

	life "github.com/aodonde/game-of-life/Mecrow/life"
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
	var cons term.Console

	cons.PrintInPlace("1st String")
	time.Sleep(2 * time.Second)
	cons.PrintInPlace("2nd String")
	time.Sleep(2 * time.Second)
	cons.PrintInPlace("3rd String")
	time.Sleep(2 * time.Second)
	cons.PrintInPlace("4th")
	time.Sleep(2 * time.Second)
	cons.PrintInPlace("5")
	time.Sleep(2 * time.Second)
	cons.PrintInPlace("Final Stirng")
	time.Sleep(2 * time.Second)
	cons.Move(term.Right, 8)
	cons.PrintInPlace("rign")
	time.Sleep(2 * time.Second)
	cons.PrintInPlace("ring, dammit")
	time.Sleep(2 * time.Second)

}

func testingMoveTo() {

	var cons term.Console

	cons.ClearScreen()
	time.Sleep(2 * time.Second)

	cons.PrintInPlace("First Line")
	cons.MoveTo(1, 1)
	cons.PrintInPlace("This was printed from 1,1")
	cons.MoveTo(2, 2)
	cons.PrintInPlace("This was printed from 2,2")
	cons.MoveTo(6, 12)
	cons.PrintInPlace("And this from 6,12")
	time.Sleep(5 * time.Second)
	cons.ClearScreen()
	cons.PrintInPlace("This should be also at 6,12")
	time.Sleep(5 * time.Second)
}

func drawBigBox() {
	var cons term.Console

	cons.ClearScreen()
	for x := 1; x < 20; x++ {
		for y := 1; y < 20; y++ {
			cons.MoveTo(x, y)
			cons.Print(term.Box)
		}
	}
	time.Sleep(5 * time.Second)
}

func testingGameMoveTo() {

	var game life.TerminalGame

	game.ClearScreen()
	time.Sleep(2 * time.Second)

	game.PrintInPlace("First Line")
	game.MoveTo(1, 1)
	game.PrintInPlace("This was printed from 1,1")
	game.MoveTo(2, 2)
	game.PrintInPlace("This was printed from 2,2")
	game.MoveTo(6, 12)
	game.PrintInPlace("And this from 6,12")
	time.Sleep(5 * time.Second)
	game.ClearScreen()
	game.PrintInPlace("This should be also at 6,12")
	time.Sleep(5 * time.Second)
}

func main() {

	testingMoveTo()
	testingGameMoveTo()
}
