package main

import (
	"object_tree_pattern/tree"
)

func main() {
	file1 := &tree.File{Name: "File1"}
	file2 := &tree.File{Name: "File2"}
	file3 := &tree.File{Name: "File3"}

	folder1 := &tree.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &tree.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
