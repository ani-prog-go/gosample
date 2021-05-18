package internal

import (
	"fmt"
	"io/ioutil"
	"time"
)

func processed(fileName string, processedDirectories []string) bool {
	for i := 0; i < len(processedDirectories); i++ {
		if processedDirectories[i] != fileName {
			continue
		}
		return true
	}
	return false
}

// обход каталогов рекурсивно
func ListDirContents(path string, dirs []string) {
	start := time.Now()
	files, _ := ioutil.ReadDir(path)

	for idx, f := range files {
		if f.Name()[0:1] == "." {
			continue // пропустим с точкой?
		}
		var newPath string
		if path != "/" {
			newPath = fmt.Sprintf("%s/%s", path, f.Name())
		} else {
			newPath = fmt.Sprintf("%s%s", path, f.Name())
		}

		if f.IsDir() {
			if !processed(newPath, dirs) {
				dirs = append(dirs, newPath)
				ListDirContents(newPath, dirs)
			}
		} else {
			fmt.Println(idx, newPath)
		}
	}
	elapsedTime := time.Since(start)
	println("Время поиска", elapsedTime.String())
}
