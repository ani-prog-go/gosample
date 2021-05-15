package main

import (
	_ "fmt"

	"github.com/ani-prog-go/gosample/sample"
)

func main() {
	scores := []int{1, 2, 3, 4, 5}
	sample.RemoveAtIndexSlice(scores, 2) // удаляем элемент с индексом
}
