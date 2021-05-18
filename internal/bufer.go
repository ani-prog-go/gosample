package internal

import (
	"bytes"
	"encoding/gob"
)

//https://golang.org/pkg/encoding/gob/  примеры
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

//Функция Get является обратной к функции set:
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

	Age int
}

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
func CodeGetBuffer(buf bytes.Buffer) {
	var usera Us
	Get(buf, &usera)
	println(usera.Name, usera.Fam, usera.Age, usera.Geo.Addr, usera.Geo.Dom)
}
