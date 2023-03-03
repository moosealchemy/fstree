package fstree

// Entry represents a file tree node.
type Entry struct {
	Name     string
	IsDir    bool
	Children []*Entry
	Locked   bool
}

// Copy deep copies the current node.
func (n *Entry) Copy() *Entry {
	cp := &Entry{
		Name:     n.Name,
		IsDir:    n.IsDir,
		Children: make([]*Entry, len(n.Children)),
	}
	for i, c := range n.Children {
		cp.Children[i] = c.Copy()
	}
	return cp
}
