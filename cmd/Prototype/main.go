package main

import (
	"fmt"

	"github.com/GuilhermeVozniak/DesignPatterns/pkg/prototype"
)

// Prototype is a creational design pattern that lets you copy existing objects
// without making your code dependent on their class in go case the interface.
func main() {
	file1 := &prototype.File{Name: "File 1"}
	file2 := &prototype.File{Name: "File 2"}
	file3 := &prototype.File{Name: "File 3"}

	folder1 := &prototype.Folder{
		Children: []prototype.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &prototype.Folder{
		Children: []prototype.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hirearchy for Folder2")
	folder2.Print("--")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hirearchy for clone Folder")

	cloneFolder.Print("--")
}
