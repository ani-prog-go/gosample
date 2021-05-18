package internal

import (
	"io/fs"
	"path/filepath"
	_ "runtime"
	"time"
)

var c111 chan string

func init() {
	c111 = make(chan string)
}
func walkG3(s string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if d.IsDir() && s[0:1] != "." {
		c111 <- "D " + s
	} else if s[0:1] != "." {
		c111 <- s
	}

	return nil
}
func findFiles3(startDir string) {
	err := filepath.WalkDir(startDir, walkG3)
	if err != nil {
		println("Error: " + err.Error())
	}

}

// используем один канал, и отлавливаем закрытие канала через range
func ListDirGorutin3() {
	//runtime.GOMAXPROCS(4)
	start := time.Now()
	go func() {
		findFiles3(".")
		close(c111)
	}()
	for val := range c111 {
		println(val)
	}

	elapsedTime := time.Since(start)
	println(elapsedTime.String())
}
