package app

import (
	"os"
	"strings"
	"github.com/Prem-099/filex/internal/explorer"
)

type App struct {
	Explorer      *explorer.Explorer
	ViewportStart int
	Selected      int
	AllFiles      []os.DirEntry
	Files         []os.DirEntry
	ClipboardPath string
	ClipboardCut  bool
	Sort 		  SortType
	ShowHidden    bool
	Mode          Mode
	Running 	  bool
}

type SortType int

const(
	SortName SortType = iota
	SortSize
	SortTime
)

func (a *App) SetMode (m Mode){
	a.Mode = m
}

func (a *App) RefreshFiles(){
	files,_ := a.Explorer.List()
	a.AllFiles = files
	var visible []os.DirEntry
	for _,f := range files{
		name := f.Name()
		if !a.ShowHidden && strings.HasPrefix(name,"."){
			continue
		}
		visible = append(visible, f)
	}
	a.Files = visible
	a.ApplySort()
}