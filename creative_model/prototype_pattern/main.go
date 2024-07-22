package main

import (
	"fmt"
	"prototype_pattern/clone"
)

//客户端

func main() {
	// 创建文件
	file1 := &clone.File{Name: "File1"}
	file2 := &clone.File{Name: "File2"}
	file3 := &clone.File{Name: "File3"}
	// 由文件组成的文件夹
	folder1 := &clone.Folder{
		Children: []clone.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &clone.Folder{
		Children: []clone.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	// 列出文件目录
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")
	// 克隆文件夹
	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
