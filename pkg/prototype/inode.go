package prototype

// Inode is an representation of a file system node
type Inode interface {
	Print(string)
	Clone() Inode
}
