package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Handler struct{}

func (h *Handler) ReadJson(file string) Data {
	// file := 
	body, errReadFile := ioutil.ReadFile(file)
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}
	data := Data{}
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	return data
}
