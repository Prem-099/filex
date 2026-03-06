package fs

import "os"

// Local File System

type LocalFS struct{}

// Here LocalFS struct has all methods of FileSystem interface
// So this is also an interface 
// Syntax : func (receiver) method_name (Parameters) return type {}
// (l localFS) : This says that l belongs to LocalFS struct and the methods defined will belong to localFS

func (l LocalFS) List(path string) ([]os.DirEntry,error){
	return os.ReadDir(path)
}

func (l LocalFS) Read(path string) ([]byte,error){
	return os.ReadFile(path)
}

func (l LocalFS) Mkdir(path string, perm os.FileMode) error {
	return os.Mkdir(path,perm)
}

func (l LocalFS) Delete(path string) error {
	return os.RemoveAll(path)
}