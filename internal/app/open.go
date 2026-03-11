package app

import (
	"os/exec"
	"runtime"
)

func OpenFile(path string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open",path)
	case "darwin":
		cmd = exec.Command("open", path)
	case "windows":
		cmd = exec.Command("cmd","/c","start","",path)
	default:
		return nil
	}
	return cmd.Start()
}