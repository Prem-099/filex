package ui

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

type Terminal struct{
	oldState *term.State
}

func (t *Terminal) Init() error {
	fd := int(os.Stdin.Fd())

	state,err := term.MakeRaw(fd)
	if err!=nil{
		return err
	}
	t.oldState = state
	HideCursor()
	return nil
}

func (t *Terminal) Close(){
	if t.oldState != nil{
		term.Restore(int(os.Stdin.Fd()),t.oldState)
		fmt.Print("\033[0m\n")
		ShowCursor()
	}
}

func ResetToTop() {
	fmt.Print("\033[H")
}

func ClearScreen(){
	fmt.Print("\033[H\033[2J")
}

func MoveCursor(row, col int){
	fmt.Printf("\033[%d;%dH", row, col)
}

func HideCursor(){
	fmt.Print("\033[?25l")
}

func ShowCursor(){
	fmt.Print("\033[?25h")
}