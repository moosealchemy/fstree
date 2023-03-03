package fstree

// Remove removes the path from the tree.
// Returns ErrNotExist if the path is not found.
// Returns ErrDeleteRoot if the path is empty.
func (t *Tree) Remove(path string) error {
	parts := SplitPath(t.rootPath, t.separator, path)
	if len(path) == 0 {
		return ErrDeleteRoot
	}
	return remove(t.root, parts)
}

func remove(root *Entry, path []string) error {
	name := path[0]
	path = path[1:]

	for i, c := range root.Children {
		if c.Name == name {
			if len(path) == 0 {
				c.Children = append(c.Children[:i], c.Children[i+1:]...)
				return nil
			}
			return remove(c, path)
		}
	}
	return ErrNotExist
}