package app

import (
	"io"
	"os"
	"path/filepath"
)

func CopyPath(src, dst string) error{
	info,err := os.Stat(src)
	if err!=nil{
		return err
	}
	if info.IsDir(){
		return copyDir(src, dst)
	}
	return copyFile(src , dst)
}

func copyFile(src, dst string) error {
	in,err := os.Open(src)
	if err!=nil{
		return err
	}
	defer in.Close()
	out,err := os.Create(dst)
	if err!=nil{
		return err
	}
	defer out.Close()
	_,err = io.Copy(out,in)
	return err 
}

func copyDir(src , dst string) error {
	err := os.MkdirAll(dst,0755)
	if err!=nil{
		return err
	}
	entries,err := os.ReadDir(src)
	if err!=nil{
		return err
	}
	for _,entry := range entries{
		srcPath := filepath.Join(src,entry.Name())
		dstPath := filepath.Join(dst,entry.Name())
		err = CopyPath(srcPath,dstPath)
		if err!=nil{
		return err
		}
	}
	return nil
}