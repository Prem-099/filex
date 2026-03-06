package app

import (
	"os"
	"strings"
	"github.com/Prem-099/filex/internal/ui"
)

type SearchMode struct{
	Query string
}

func NewSearchMode() *SearchMode{
	return &SearchMode{}
}

func (s *SearchMode) Name() string {
	return "SEARCH MODE"
}

func (s *SearchMode) HandleKey(key ui.Key,ch rune,a *App){
	switch key{
	case ui.KeyChar:
		s.Query += string(ch)
		filterFiles(s.Query,a)
	case ui.KeyBack:
		if len(s.Query) > 0{
			s.Query = s.Query[:len(s.Query)-1]
			filterFiles(s.Query,a)
		}
	case ui.KeyEscape:
		a.Files = a.AllFiles
		a.SetMode(&NormalMode{})
	case ui.KeyEnter:
		a.SetMode(&NormalMode{})
	case ui.KeyQuit:
		a.Running = false
	}
}

func filterFiles(query string,a *App){
	query = strings.ToLower(query)
	var result []os.DirEntry
	for _,f := range a.AllFiles{
		name := strings.ToLower(f.Name())
		if strings.Contains(name,query){
			result = append(result, f)
		}
	}
	a.Files = result
	a.Selected = 0
	a.ViewportStart = 0
}