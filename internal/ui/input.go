package ui

import (
	"os"
)

type Key int

const (
	KeyUnkown = iota
	KeyUp
	KeyDown
	KeyBack
	KeyEnter
	KeyQuit
	KeyDelete
	KeyRename
	KeyChar
	KeyEscape
	KeySearch
	KeyCopy
	KeyCut
	KeyPaste
	KeySortName 
	KeySortSize
	KeySortTime
	KeyHelp
	KeyToggleHidden
)

func ReadKey() (Key,rune){
	buf := make([]byte,1)
	_,err:=os.Stdin.Read(buf)
	if err!=nil{
		return KeyUnkown, 0
	}
	switch buf[0]{
	case 13:
		return KeyEnter, 0
	/*case 'q':
		return KeyQuit, 0
	case 'h':
		return KeyBack, 0
    case 'd':
		return KeyDelete, 0
	case 'r':
		return KeyRename, 0
	case 'y':
		return KeyCopy, 0
	case 'x':
		return KeyCut, 0
	case 'p':
		return KeyPaste, 0
	case 'n':
		return KeySortName, 0
	case 's':
		return KeySortSize, 0
	case 't':
		return KeySortTime, 0
	case '/':
		return KeySearch, 0
	case '.':
		return KeyHelp, 0*/
	case 8:
		return KeyBack, 0
	case 127:
		return KeyBack, 0
	case 27:
		seq := make([]byte,2)
		os.Stdin.Read(seq)
		if seq[0]==91{
			switch seq[1]{
			case 65:
				return KeyUp, 0
			case 66:
				return KeyDown, 0
			}
		}
		return KeyEscape, 0
	}
	if buf[0]>=32 && buf[0] <= 126{
		return KeyChar, rune(buf[0])
	}
	return KeyUnkown, 0
}