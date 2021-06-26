package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/fastrodev/serverless/internal"
)

func writeFile(data string) {
	f, errCreate := os.Create("static/index.html")
	if errCreate != nil {
		log.Fatal(errCreate)
	}

	_, errWrite := f.WriteString(data)
	if errWrite != nil {
		log.Fatal(errWrite)
	}

	defer f.Close()
}

func main() {
	h := internal.Handler{}
	td := h.ReadJson("internal/iklan.json")

	t, err := template.ParseFiles("template/default.tmpl")
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, td.Data)
	if err != nil {
		panic(err)
	}

	writeFile(tpl.String())
}
