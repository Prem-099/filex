package ui

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

const (
	ColorReset    = "\033[0m"
	ColorSelected = "\033[42m\033[30m"
	ClearLine = "\033[K"
)

func Render(path string, files []os.DirEntry, selected int, viewportStart int,
	mode string, renameInput string, searchQuery string,
	sortType string, clipboard string, preview []string) {

	ClearScreen()
	HideCursor()

	w := os.Stdout
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24
	}
	MoveCursor(1, 1)
	fmt.Fprintf(w, "Path: %s\n", path)
	fmt.Fprintf(w, "\033[K")
	MoveCursor(2, 1)
	fmt.Fprintln(w, "────────────────────────────")

	startRow := 3
	visibleHeight := height - 6

	end := viewportStart + visibleHeight
	if end > len(files) {
		end = len(files)
	}

	if mode == "HELP MODE" {

		ClearScreen()

		MoveCursor(2, 2)
		fmt.Println("FILEX HELP")
		MoveCursor(3, 2)
		fmt.Println("────────────────────────")

		fmt.Println("\nNavigation")
		fmt.Println("  ↑ ↓        move")
		fmt.Println("  Enter      open file/folder")
		fmt.Println("  Backspace  parent folder")

		fmt.Println("\nFile Operations")
		fmt.Println("  c           copy")
		fmt.Println("  x           cut")
		fmt.Println("  v           paste")
		fmt.Println("  r           rename")
		fmt.Println("  d           delete")

		fmt.Println("\nSearch")
		fmt.Println("  /           search")
		fmt.Println("  ESC         exit search")

		fmt.Println("\nSorting")
		fmt.Println("  n           sort name")
		fmt.Println("  s           sort size")
		fmt.Println("  t           sort time")

		fmt.Println("\nOther")
		fmt.Println("  .           toggle hidden files")
		fmt.Println("  q           quit")
		fmt.Println("\nPress any key to return")
		ShowCursor()
		return
	}

	for i := viewportStart; i < end; i++ {
		screenRow := startRow + (i - viewportStart)
		MoveCursor(screenRow, 1)
		cursor := "  "
		if i == selected {
			cursor = ">"
			fmt.Print(ColorSelected)
		}

		icon := "[F]"
		if files[i].IsDir() {
			icon = "[D]"
		}

		name := files[i].Name()
		if len(name) > 30 {
			name = name[:27] + "..."
		}
		info, _ := files[i].Info()
		sizeText := ""
		if !files[i].IsDir() {
			sizeText = formatSize(info.Size())
		}
		fmt.Fprintf(w, "%2s %s %-30s %8s", cursor, icon, name, sizeText)
		fmt.Print(ColorReset)
		fmt.Fprint(w, "\033[K")
	}

	for r := 3; r < height-2; r++ {
		MoveCursor(r, width/2)
		fmt.Print("│")
	}

	previewStart := width/2 + 2
	for j, line := range preview {
		row := 3 + j
		if row >= height-2 {
			break
		}
		MoveCursor(row, previewStart)
		if len(line) > width/2-4 {
			line = line[:width/2-7] + "..."
		}
		fmt.Print(line)
	}

	if mode == "CONFIRM DELETE MODE" {
		midRow := height / 2
		midCol := width / 2
		MoveCursor(midRow-1, midCol-15)
		fmt.Print("┌──────────────────────────┐")
		MoveCursor(midRow, midCol-15)
		fmt.Print("│   Delete selected item?  │")
		MoveCursor(midRow+1, midCol-15)
		fmt.Print("│  Enter = Yes  Back = No  │")
		MoveCursor(midRow+2, midCol-15)
		fmt.Print("└──────────────────────────┘")
	}

	if mode == "RENAME MODE" {
		boxRow := height / 2
		boxCol := width / 2
		MoveCursor(boxRow-1, boxCol-15)
		fmt.Print("┌──────────────────────────┐")
		MoveCursor(boxRow, boxCol-15)
		fmt.Print("│ Rename to:               │")
		MoveCursor(boxRow+1, boxCol-15)
		fmt.Printf("│ %-24s │", renameInput)
		MoveCursor(boxRow+2, boxCol-15)
		fmt.Print("└──────────────────────────┘")
	}

	if mode == "SEARCH MODE" {
		boxRow := height / 2
		boxCol := width / 2
		MoveCursor(boxRow-1, boxCol-15)
		fmt.Print("┌──────────────────────────┐")
		MoveCursor(boxRow, boxCol-15)
		fmt.Print("│ Search:                  │")
		MoveCursor(boxRow+1, boxCol-15)
		fmt.Printf("│ %-24s │", searchQuery)
		MoveCursor(boxRow+2, boxCol-15)
		fmt.Print("└──────────────────────────┘")
	}

	MoveCursor(height-1, 1)
	fmt.Fprintln(w, "────────────────────────────────────────")
	MoveCursor(height, 1)
	status := fmt.Sprintf("[%s] sort:%s", mode, sortType)
	if clipboard != "" {
		status += " | clipboard:" + clipboard
	}
	fmt.Fprintf(w, "%d items | %s | ? help | q quit", len(files), status)
	fmt.Fprint(w, "\033[K")
	ShowCursor()
}

func formatSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	if size < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(size)/1024)
	}
	if size < 1024*1024*1024 {
		return fmt.Sprintf("%.1f MB", float64(size)/(1024*1024))
	}
	if size < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.1f GB", float64(size)/(1024*1024*1024))
	}
	return ""
}