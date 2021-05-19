package main

import (
	"fmt"
	_ "fmt"
	"runtime"

	"github.com/ani-prog-go/gosample/internal"
	"github.com/ani-prog-go/gosample/sample"
)

func main() {
	fmt.Printf("Версия Go: %s\n", runtime.Version())
	cmds := internal.CommandString()
	//cmds.str
	println("Ком строка: ", cmds.Str)
	scores := []int{1, 2, 3, 4, 5}
	sample.RemoveAtIndexSlice(scores, 2) // удаляем элемент с индексом
	internal.ListDirGorutin3()
	internal.PrintColor(internal.ColorGreen, "Привет")
	internal.Bolt()
	internal.BoltJson()
	buf, _ := internal.CodeToBufer()
	internal.CodeGetBuffer(buf)
	buf2, _ := internal.CodeToBuferArr()
	internal.CodeGetBufferArr(buf2)
	internal.ListDirContents(".", []string{})
	internal.LoadJsonFile()
	internal.LoadJsonHtml()
	internal.CriptoAes256()
}
