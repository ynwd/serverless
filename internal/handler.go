package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fastrodev/fastrex"
)

func init() {
	fmt.Println("init is executed")
}

func readJson() Data {
	body, errReadFile := ioutil.ReadFile("static/iklan.json")
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

type Handler struct {
	data Data
}

func (h *Handler) htmlHandler(req fastrex.Request, res fastrex.Response) {
	res.Render(h.data.Data)
}
