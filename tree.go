package fstree

type Tree struct {
	rootPath  string
	separator string

	root *Entry
}

func New(rootPath, separator string) *Tree {
	return &Tree{
		root: &Entry{
			IsDir: true,
		},
		rootPath:  CleanPath(rootPath, separator),
		separator: separator,
	}
}

func (t *Tree) Copy() []*Entry {
	return t.root.Copy().Children
}
