package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/fastrodev/serverless/internal"
)

func writeFile(data string) {
	f, err := os.Create("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := f.WriteString(data)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer f.Close()
	fmt.Println("done")
}

func main() {
	h := internal.Handler{}
	td := h.ReadJson()

	t, err := template.ParseFiles("template/default.gohtml")
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
