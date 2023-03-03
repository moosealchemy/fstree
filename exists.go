package fstree

// Exists returns true if the path exists, or false if it does not.
func (t *Tree) Exists(path string) (isDir bool, exists bool) {
	return treeExists(t.root, SplitPath(t.rootPath, t.separator, path))
}

func treeExists(root *Entry, path []string) (bool, bool) {
	if len(path) == 0 {
		return root.IsDir, true
	}
	return checkChildrenExists(root, path)
}

func checkChildrenExists(root *Entry, path []string) (bool, bool) {
	for _, c := range root.Children {
		if c.Name == path[0] {
			return treeExists(c, path[1:])
		}
	}
	return false, false
}
