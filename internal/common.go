package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func WriteFile(data string, output string) {
	f, errCreate := os.Create(output)
	if errCreate != nil {
		log.Fatal(errCreate)
	}

	_, errWrite := f.WriteString(data)
	if errWrite != nil {
		log.Fatal(errWrite)
	}

	defer f.Close()
}

func ReadJson(file string) []Data {
	body, errReadFile := ioutil.ReadFile(file)
	if errReadFile != nil {
		log.Fatal("ReadJson" + errReadFile.Error())
	}
	data := []Data{}
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		log.Fatal("ReadJson" + errUnmarshal.Error())
	}
	return data
}

func ReadData(file string) []Post {
	body, errReadFile := ioutil.ReadFile(file)
	if errReadFile != nil {
		log.Fatal("ReadData" + errReadFile.Error())
	}
	data := []Post{}
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		log.Fatal("ReadData" + errUnmarshal.Error())
	}
	return data
}
