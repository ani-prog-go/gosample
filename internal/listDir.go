package internal

import (
	"io/fs"
	"path/filepath"
)

func walk(s string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if d.IsDir() && s[0:1] != "." {
		println(s)
	}
	return nil
}

func ListDir() {
	filepath.WalkDir(".", walk)
}
