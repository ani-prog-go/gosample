package internal

import (
	"io/fs"
	"path/filepath"
	_ "runtime"
	"time"
)

var c11 chan string

func init() {
	c11 = make(chan string)
}
func walkG2(s string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if d.IsDir() && s[0:1] != "." {
		c11 <- "D " + s
	} else if s[0:1] != "." {
		c11 <- s
	}

	return nil
}
func findFiles2(startDir string) {
	err := filepath.WalkDir(startDir, walkG2)
	if err != nil {
		println("Error: " + err.Error())
	}

}

// используем один канал, и отлавливаем закрытие канала
func ListDirGorutin2() {
	//runtime.GOMAXPROCS(4)
	start := time.Now()
	go func() {
		findFiles2(".")
		close(c11)
	}()
	for {
		val, ok := <-c11
		if !ok { // если канал закрыт
			println("close>>>" + val)
			break
		}
		println(val)
	}

	elapsedTime := time.Since(start)
	println(elapsedTime.String())
}
