package fs

import "os"

type FileSystem interface{
	List(path string)([]os.DirEntry,error)
	Read(path string)([]byte,error)
	Mkdir(path string, perm os.FileMode) error
	Delete(path string) error
}


// Anything that satisfies interface rules or provides all prescribed methods 
// will be automatically considered as interface by GO  