package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type UserFromFile struct {
	Name       string
	Occupation string
	Born       string
}

func LoadJsonFile() {
	filename, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()
	// d := time .Now()
	// println(d)
	data, err := ioutil.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	var result []UserFromFile

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("%+v\n", result)
}
