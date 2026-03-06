package app

import (
	"bufio"
	"os"
)

const previewLimit = 200 * 1024

func ReadPreview(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err!=nil{
		return nil,err
	}
	if info.Size() > previewLimit {
		return []string{"file too large to preview"},nil
	}
	file,err := os.Open(path)
	if err!=nil{
		return nil,err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
		count++
		if count > 200 {
			break
		}
	}
	return lines,nil
}
