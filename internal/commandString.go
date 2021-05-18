// работа с коммандной мтрокой
package internal

import (
	"flag"
)

type Cmdstr struct {
	Str string
}

var cmdS Cmdstr

func CommandString() Cmdstr {
	// ссылка на переменную, имя параметра, значение по умолчанию, описание для хелпа
	flag.StringVar(&cmdS.Str, "scmd", "не задано", "Укажите строку для теста")
	flag.Parse() // фиксируем все
	// if scmd != "" {
	//println("Команда " + cmdS.Str)
	// }
	return cmdS
}
