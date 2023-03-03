package fstree

import (
	"errors"
)

// Lock locks the given path. Returns true if the operation was successful.
// The error will be ErrNotExist if the path does not exist.
func (t *Tree) Lock(path string) error {
	return setLock(t.root, true, SplitPath(t.rootPath, t.separator, path))
}

// Unlock unlocks the given path. Returns true if the operation was successful.
// The error will be ErrNotExist if the path does not exist.
func (t *Tree) Unlock(path string) error {
	return setLock(t.root, false, SplitPath(t.rootPath, t.separator, path))
}

func (t *Tree) setLock(path string, lock bool) error {
	return setLock(t.root, lock, SplitPath(t.rootPath, t.separator, path))
}

func setLock(root *Entry, lock bool, path []string) error {
	if len(path) == 0 {
		return setLockAll(root, lock)
	}
	return findAndLock(root, lock, path)
}

func findAndLock(root *Entry, lock bool, path []string) error {
	for _, c := range root.Children {
		if c.Name == path[0] {
			return setLock(c, lock, path[1:])
		}
	}
	return ErrNotExist
}

func setLockAll(root *Entry, lock bool) error {
	updated := root.Locked != lock
	for _, c := range root.Children {
		updated = updated || setLockAll(c, lock) == nil
	}

	if updated {
		return nil
	} else if lock {
		return ErrLocked
	}
	return ErrUnlocked
}

// IsLocked takes a path and returns one of three errors depending on the status of the path:
// - ErrLocked if the path is locked.
// - ErrUnlocked if the path is unlocked.
// - ErrNotExist if the path does not exist.
func (t *Tree) IsLocked(path string) error {
	return isLocked(t.root, SplitPath(t.rootPath, t.separator, path))
}

func isLocked(root *Entry, path []string) error {
	if len(path) == 0 {
		return foundIsLocked(root)
	}

	for _, c := range root.Children {
		if c.Name == path[0] {
			return isLocked(c, path[1:])
		}
	}
	return ErrNotExist
}

func foundIsLocked(root *Entry) error {
	if root.IsDir {
		for _, c := range root.Children {
			if errors.Is(foundIsLocked(c), ErrLocked) {
				return ErrLocked
			}
		}
		return ErrUnlocked
	} else {
		if root.Locked {
			return ErrLocked
		} else {
			return ErrUnlocked
		}
	}
}
