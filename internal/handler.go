package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fastrodev/fastrex"
)

type Handler struct {
	data       Data
	serverless bool
}

func (h *Handler) htmlHandler(req fastrex.Request, res fastrex.Response) {
	res.Render(h.data.Data)
}

func (h *Handler) ReadJson() Data {
	file := "internal/iklan.json"
	if h.serverless {
		file = "serverless_function_source_code/internal/iklan.json"
	}
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
