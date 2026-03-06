package app

import (
	"os"
	"path/filepath"
	"github.com/Prem-099/filex/internal/ui"
)

type NormalMode struct {}

func (m *NormalMode) Name () string {
	return "NORMAL MODE"
}

func (m *NormalMode) HandleKey(key ui.Key,ch rune,a *App){
	switch key{
		case ui.KeyUp:
			a.Selected--
			if a.Selected < 0{
				a.Selected = len(a.Files)-1
			}
			if a.Selected < a.ViewportStart {
				a.ViewportStart = a.Selected
			}
		case ui.KeyDown:
			a.Selected++
			if a.Selected >= len(a.Files){
				a.Selected = 0
			}
			visibleHeight := 20
			if a.Selected >= a.ViewportStart+visibleHeight{
				a.ViewportStart++
			}
		case ui.KeyBack:
			a.Explorer.GoBack()
			a.Selected = 0
			a.ViewportStart = 0
			a.RefreshFiles()
		case ui.KeyEnter:
			item := a.Files[a.Selected]
			if item.IsDir(){
				a.Explorer.Enter(item.Name())
				a.Selected = 0
				a.ViewportStart = 0
				a.RefreshFiles()
			}else{
				path := filepath.Join(a.Explorer.CurrentPath,item.Name())
				abs,_ := filepath.Abs(path)
				OpenFile(abs)
			}
		case ui.KeyChar:
			switch ch{
			case 'd':
				if len(a.Files) == 0{
					return
				}
				a.SetMode(&ConfirmDeleteMode{})
			case 'r':
				if len(a.Files) == 0{
					return
				}
				a.SetMode(NewRenameMode(a.Files[a.Selected].Name()))
			case 'c':
				if len(a.Files) == 0{
					return
				}
				name := a.Files[a.Selected].Name()
				a.ClipboardPath = filepath.Join(a.Explorer.CurrentPath,name)
				a.ClipboardCut = false
			case 'x':
				if len(a.Files) == 0{
					return
				}
				name := a.Files[a.Selected].Name()
				a.ClipboardPath = filepath.Join(a.Explorer.CurrentPath,name)
				a.ClipboardCut = true
			case 'v':
				if a.ClipboardPath == ""{
					return
				}
				name := filepath.Base(a.ClipboardPath)
				dest := filepath.Join(a.Explorer.CurrentPath,name)
				if a.ClipboardCut{
					os.Rename(a.ClipboardPath,dest)
					}else{
						CopyPath(a.ClipboardPath,dest)
					}
				a.ClipboardPath = ""
				a.ClipboardCut = false
				a.RefreshFiles()
			case 'y':
				a.ClipboardPath = ""
				a.ClipboardCut = false
			case '/':
				a.SetMode(NewSearchMode())
			case 'n':
				a.Sort = SortName
				a.ApplySort()
			case 's':
				a.Sort = SortSize
				a.ApplySort()
			case 't':
				a.Sort = SortTime
				a.ApplySort()
			case 'q':
				a.Running = false
			case '?':
				a.SetMode(NewHelpMode())
			case '.':
				a.ShowHidden = !a.ShowHidden
				a.RefreshFiles()
				a.Selected = 0
				a.ViewportStart = 0
			}
		case ui.KeyQuit:
			ui.ClearScreen()
			a.Running = false 
	}
}