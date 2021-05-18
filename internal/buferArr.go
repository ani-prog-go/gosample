package internal

import (
	"bytes"
	"encoding/gob"
)

//https://golang.org/pkg/encoding/gob/  примеры
func SetArr(value interface{}) (bytes.Buffer, error) {
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

func GetArr(buf bytes.Buffer, value interface{}) error {

	//return gob.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(value)
	return gob.NewDecoder(&buf).Decode(value)
}

type GeoaArr struct {
	Addr string
	Dom  float32
}
type UsArr struct {
	Name string
	Geo  Geoa
	Fam  string

	Age int
}
type userM []UsArr

func CodeToBuferArr() (bytes.Buffer, error) {

	usera := UsArr{
		Name: "ark",
		Geo: Geoa{
			Addr: "khudognikov",
			Dom:  345.789,
		},
		Fam: "Ив",
		Age: 78,
	}

	var userb userM
	userb = append(userb, usera)
	userb[0].Geo.Dom = 80
	userb[0].Age = 10
	userb = append(userb, usera)
	userb[1].Geo.Dom = 81
	userb[1].Age = 11
	userb = append(userb, usera)
	userb[2].Geo.Dom = 82
	userb[2].Age = 12
	userb = append(userb, usera)
	userb[3].Geo.Dom = 83
	userb[3].Age = 13

	buf, err := Set(userb)
	if err != nil {
		println(err)
		return bytes.Buffer{}, err
	}
	println(buf.Bytes())
	return buf, nil
	//CodeGetBuffer(buf)

}
func CodeGetBufferArr(buf bytes.Buffer) {
	var usera userM
	Get(buf, &usera)
	for idx, element := range usera {
		println(idx, element.Name, element.Fam, element.Age, element.Geo.Addr, element.Geo.Dom)
	}
}
