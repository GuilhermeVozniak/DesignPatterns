package prototype

import "fmt"

// File is an abstraction of a file in a OS
type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() Inode {
	return &File{Name: f.Name + "_clone"}
}
