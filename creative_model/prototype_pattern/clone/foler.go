package clone

import "fmt"

type Folder struct {
	Children []Inode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.Children {
		copyF := i.Clone()
		tempChildren = append(tempChildren, copyF)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
