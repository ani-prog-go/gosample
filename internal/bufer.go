package internal

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

//https://golang.org/pkg/encoding/gob/  примеры
// загоняем структуру в буфер
func Set(value interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	var encoder = gob.NewEncoder(&buf)

	//id := "sample-key"
	err := encoder.Encode(value)
	if err != nil {
		return bytes.Buffer{}, err
	}
	// return redisdb.HSet(id, field, buf.Bytes()).Err()
	return buf, nil
}

//Функция Get получем структуры из буфера
func Get(buf bytes.Buffer, value interface{}) error {

	//return gob.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(value)
	return gob.NewDecoder(&buf).Decode(value)
}

type Geoa struct {
	Addr string
	Dom  float32
}
type Us struct {
	Name string
	Geo  Geoa
	Fam  string
	Age  int
}

// записываем струтктуру в буфер
func CodeToBufer() (bytes.Buffer, error) {

	usera := Us{
		Name: "ark",
		Geo: Geoa{
			Addr: "khudognikov",
			Dom:  345.789,
		},
		Fam: "Ив",
		Age: 78,
	}

	buf, err := Set(usera)
	if err != nil {
		println(err)
		return bytes.Buffer{}, err
	}
	println(buf.Bytes())
	return buf, nil
	//CodeGetBuffer(buf)

}

// получаем структуру из буфера
// , и вывод JSON в форматированном виде с отступами
func CodeGetBuffer(buf bytes.Buffer) {
	var usera Us
	Get(buf, &usera)

	data, err := json.MarshalIndent(&usera, "", "   ") // отступы

	if err != nil {
		println(err)
	}
	fmt.Println(string(data))
	//	println(usera.Name, usera.Fam, usera.Age, usera.Geo.Addr, usera.Geo.Dom)
}
