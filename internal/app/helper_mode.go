package app

import "github.com/Prem-099/filex/internal/ui"

type Helper struct{}

func NewHelpMode() *Helper{
	return &Helper{}
}

func (h *Helper) Name() string {
	return "HELP MODE"
}

func (h *Helper) HandleKey(key ui.Key, ch rune,a *App){
	a.SetMode(&NormalMode{})
}