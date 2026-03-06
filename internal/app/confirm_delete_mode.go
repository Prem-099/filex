package app

import (
	"github.com/Prem-099/filex/internal/ui"
)

type ConfirmDeleteMode struct {}

func (m *ConfirmDeleteMode) Name() string{
		return "CONFIRM DELETE MODE"
}

func (m *ConfirmDeleteMode) HandleKey(key ui.Key,ch rune,a *App){
	switch key{
	case ui.KeyEnter:
		item := a.Files[a.Selected]
		a.Explorer.Delete(item.Name())
		a.RefreshFiles()
		a.SetMode(&NormalMode{})
	case ui.KeyBack:
		a.SetMode(&NormalMode{})
	case ui.KeyQuit:
		a.SetMode(&NormalMode{})
		a.Running = false
	}
} 