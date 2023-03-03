package fstree

// Add adds the path to the tree.
// isDir specifies if the leaf is a directory.
// Any directories in the path will be created if they do not exist.
// Returns ErrExist if attempting to add a leaf that already exists.
// Returns AddMismatchedIsDir if mismatched isDir values are encountered.
func (t *Tree) Add(path string, isDir bool) error {
	parts := SplitPath(t.rootPath, t.separator, path)
	if len(path) == 0 {
		return ErrExist
	}
	return add(t.root, isDir, parts)
}

func add(root *Entry, isDir bool, path []string) error {
	if len(path) == 0 {
		return ErrExist
	}

	name := path[0]
	path = path[1:]
	thisIsDir := len(path) > 0 || isDir

	for _, c := range root.Children {
		if c.Name == name {
			if len(path) == 0 {
				if thisIsDir != c.IsDir {
					return &AddMismatchedIsDir{
						Expected: thisIsDir,
						Got:      c.IsDir,
					}
				}
				return ErrExist
			} else if !c.IsDir {
				return &AddMismatchedIsDir{
					Expected: true,
				}
			}
			return add(c, isDir, path)
		}
	}

	n := &Entry{
		Name:  name,
		IsDir: thisIsDir,
	}

	root.Children = append(root.Children, n)
	if len(path) > 0 {
		return add(n, isDir, path)
	}
	return nil
}
