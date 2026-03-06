package app

import (
	"os"
	"path/filepath"
	"github.com/Prem-099/filex/internal/ui"
)

type RenameMode struct{
	OriginalName string
	Input string
}

func NewRenameMode(orgName string,) *RenameMode{
	return &RenameMode{
		OriginalName: orgName,
		Input: orgName,
	}
}

func (m *RenameMode) Name() string {
	return "RENAME MODE"
}

func (m *RenameMode) HandleKey(key ui.Key,ch rune,a *App){
	switch key{
	case ui.KeyChar:
		m.Input += string(ch)
	case ui.KeyBack:
		if len(m.Input) > 0{
			m.Input = m.Input[:len(m.Input)-1]
		}
	case ui.KeyEnter:
		if m.Input != "" {
			oldPath := filepath.Join(a.Explorer.CurrentPath,m.OriginalName)
			newPath := filepath.Join(a.Explorer.CurrentPath,m.Input)
			os.Rename(oldPath,newPath)
		}
		a.RefreshFiles()
		a.SetMode(&NormalMode{})
	case ui.KeyEscape:
		a.SetMode(&NormalMode{})
	case ui.KeyQuit:
		a.Running = false
	}
}