package explorer

import (
	"os"
	"path/filepath"
	"github.com/Prem-099/filex/internal/fs"
)

type Explorer struct{
	FS fs.FileSystem
	CurrentPath string
}

func New(fileSystem fs.FileSystem, startPath string) *Explorer {
	return &Explorer{
		FS: fileSystem,
		CurrentPath: startPath,
	}
} 

func (e *Explorer) List()([]os.DirEntry, error){
	return e.FS.List(e.CurrentPath)
}

func (e *Explorer) GoBack(){
	e.CurrentPath = filepath.Dir(e.CurrentPath)
}

func (e *Explorer) Enter(name string){
	e.CurrentPath = filepath.Join(e.CurrentPath,name)
}

func (e *Explorer) CreateFolder(name string) error {
	path := filepath.Join(e.CurrentPath,name)
	return e.FS.Mkdir(path, 0755)
}

func (e *Explorer) Delete(name string) error {
	path := filepath.Join(e.CurrentPath,name)
	return e.FS.Delete(path)
}

func (e *Explorer) ReadFile(name string)([]byte, error){
	path := filepath.Join(e.CurrentPath,name)
	return e.FS.Read(path)
}