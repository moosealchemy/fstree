package fstree

import (
	"bytes"
	"strings"
)

func SplitPath(rootPath, separator, path string) []string {
	path = CleanPath(path, separator)
	path = strings.TrimPrefix(path, rootPath)
	p := strings.Split(path, separator)
	for i := 0; i < len(p); i++ {
		if p[i] == "" {
			p = append(p[:i], p[i+1:]...)
		}
	}
	return p
}

func CleanPath(path, separator string) string {
	if len(path) == 0 {
		return "."
	}

	startRoot := string(path[0]) == separator
	isDir := string(path[len(path)-1]) == separator

	doubleSlash := []byte{separator[0], separator[0]}
	b := []byte(path)
	for first, i := true, 0; i != -1; i = bytes.Index(b, doubleSlash) {
		if first {
			first = false
			continue
		} else {
			b = append(b[:i], b[i+1:]...)
		}
	}

	parts := bytes.Split(b, []byte(separator))
	for i := 0; i < len(parts); i++ {
		switch {
		case i > 0 && string(parts[i]) == "..":
			parts = append(parts[:i-1], parts[i:]...)
			i--
			fallthrough
		case len(parts[i]) == 0:
			fallthrough
		case i == 0 && string(parts[i]) == "..":
			fallthrough
		case len(parts[i]) == 1 && parts[i][0] == byte('.'):
			parts = append(parts[:i], parts[i+1:]...)
			i--
		}
	}

	path = string(bytes.Join(parts, []byte(separator)))
	if startRoot {
		path = separator + path
	}
	if isDir {
		path = path + separator
	}
	return path
}
