package internal // import "github.com/ani-prog-go/gosample/internal"

примеры с буфером структуры в буфер и обратно

работа с коммандной мтрокой

FUNCTIONS

func Bolt()
    Работа с BD boltDB

func BoltJson()
    запись чтение из базы boltDB

func CodeGetBuffer(buf bytes.Buffer)
    получаем структуру из буфера

func CodeGetBufferArr(buf bytes.Buffer)
    получаем массив структур из буфера

func CodeToBufer() (bytes.Buffer, error)
    записываем струтктуру в буфер

func CodeToBuferArr() (bytes.Buffer, error)
    делаем массив структур и загоняем в буфер

func Get(buf bytes.Buffer, value interface{}) error
    Функция Get получем структуры из буфера

func GetArr(buf bytes.Buffer, value interface{}) error
func ListDir()
func ListDirGorutin()
    используем два канала один для данных другой для завершения функции

func ListDirGorutin2()
    используем один канал, и отлавливаем закрытие канала

func ListDirGorutin3()
    используем один канал, и отлавливаем закрытие канала через range

func PrintColor(color Color, message string)
func Set(value interface{}) (bytes.Buffer, error)
    https://golang.org/pkg/encoding/gob/ примеры загоняем структуру в буфер

func SetArr(value interface{}) (bytes.Buffer, error)
    https://golang.org/pkg/encoding/gob/ примеры запись структуры в буфер


TYPES

type Cmdstr struct {
	Str string
}

func CommandString() Cmdstr

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)
type Geoa struct {
	Addr string
	Dom  float32
}

type GeoaArr struct {
	Addr string
	Dom  float32
}

type Us struct {
	Name string
	Geo  Geoa
	Fam  string

	Age int
}

type UsArr struct {
	Name string
	Geo  Geoa
	Fam  string

	Age int
}

type User struct {
	Id   int
	Name string
	Age  int
}

