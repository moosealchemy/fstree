package fstree

import (
	"errors"
	"fmt"
	"io/fs"
)

var (
	ErrDeleteRoot = errors.New("cannot delete root")
	ErrLocked     = errors.New("path is locked")
	ErrUnlocked   = errors.New("path is unlocked")
	ErrExist      = fs.ErrExist
	ErrNotExist   = fs.ErrNotExist
)

type AddMismatchedIsDir struct {
	Expected bool
	Got      bool
}

func (e *AddMismatchedIsDir) Error() string {
	return fmt.Sprintf("expected isDir=%t got isDir=%t when adding to tree", e.Expected, e.Got)
}
