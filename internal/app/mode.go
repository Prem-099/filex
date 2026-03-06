package app

import "github.com/Prem-099/filex/internal/ui"

type Mode interface {
	HandleKey(key ui.Key,ch rune, a *App)
	Name() string
}

// Any struct having these methods can be considered as mode 
