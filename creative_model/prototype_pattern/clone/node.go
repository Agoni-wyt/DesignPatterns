package clone

type Inode interface {
	Print(string)
	Clone() Inode
}
