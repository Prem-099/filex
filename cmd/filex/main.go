package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Prem-099/filex/internal/app"
	"github.com/Prem-099/filex/internal/explorer"
	"github.com/Prem-099/filex/internal/fs"
	"github.com/Prem-099/filex/internal/ui"
)

const Version = "v1.0.0"

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "--version"{
			fmt.Println("filex", Version)
			return
		}
		if os.Args[1] == "--help"{
			fmt.Println("filex - terminal file explorer")
			fmt.Println("")
			fmt.Println("Usage:")
			fmt.Println(" filex [path] ")
			fmt.Println("")
			fmt.Println("options:")
			fmt.Println(" --help         show this message")
			fmt.Println(" --version      show version")
			fmt.Println("")
			fmt.Println("Examples")
			fmt.Println(" filex")
			fmt.Println(" filex /Downloads")
			fmt.Println(" filex /etc")
			return
		}
	}

	start := ""
	if len(os.Args) > 1 {
		start = os.Args[1]
	}else{
		var err error
		start,err = os.Getwd()
		if err!=nil{
			panic(err)
		}
	}
	if _,err := os.Stat(start); err!=nil{
		fmt.Println("Invalid path:", start)
		return
	}

	term := ui.Terminal{}
	term.Init()
	defer term.Close()

	ui.HideCursor()
	defer ui.ShowCursor()

	exp := explorer.New(fs.LocalFS{}, start)
	a := &app.App{
		Explorer: exp,
		Running:  true,
	}
	a.RefreshFiles()
	a.SetMode(&app.NormalMode{})

	for a.Running {
		renameInput := ""
		searchQuery := ""
		clipboard := ""

		if a.ClipboardPath != "" {
			name := filepath.Base(a.ClipboardPath)
			mode := "copy"
			if a.ClipboardCut {
				mode = "move"
			}
			clipboard = name + " (" + mode + ")"
		}

		if r, ok := a.Mode.(*app.RenameMode); ok {
			renameInput = r.Input
		}

		if s, ok := a.Mode.(*app.SearchMode); ok {
			searchQuery = s.Query
		}

		preview := []string{}
		if len(a.Files) > 0 && a.Selected < len(a.Files) {
			item := a.Files[a.Selected]
			if !item.IsDir() {
				path := filepath.Join(a.Explorer.CurrentPath, item.Name())
				abs, _ := filepath.Abs(path)
				lines, err := app.ReadPreview(abs)
				if err == nil {
					preview = lines
				}
			}
		}
		ui.Render(
			a.Explorer.CurrentPath,
			a.Files,
			a.Selected,
			a.ViewportStart,
			a.Mode.Name(),
			renameInput,
			searchQuery,
			a.SortName(),
			clipboard,
			preview,
		)
		key, ch := ui.ReadKey()
		a.Mode.HandleKey(key, ch, a)
	}
}
