package internal

import (
	"io/fs"
	"path/filepath"
	"runtime"
	"time"
)

var c1 chan string
var c2 chan string

func init() {
	c1 = make(chan string)
	c2 = make(chan string)
}
func walkG(s string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if d.IsDir() && s[0:1] != "." {
		//		i++
		//		println(i)
		c1 <- "D " + s
	} else if s[0:1] != "." {
		//		i++
		//		println(i)
		c1 <- s
	}

	return nil
}
func findFiles(startDir string) {
	err := filepath.WalkDir(startDir, walkG)
	if err != nil {
		println("Error: " + err.Error())
	}

}

// используем два канала один для данных другой для завершения функции
func ListDirGorutin() {
	runtime.GOMAXPROCS(4)
	start := time.Now()
	go func() {
		findFiles(".")
		c2 <- "__EOF__"
	}()
	go func() {
		for {
			println(<-c1)
		}
	}()
	<-c2
	//time.Sleep(time.Second * 1)
	elapsedTime := time.Since(start)
	println(elapsedTime.String())

}
