package main

import (
	"bytes"
	"text/template"
	"time"

	"github.com/fastrodev/serverless/internal"
)

func main() {
	td := internal.ReadJson("internal/data/index.json")

	t, err := template.ParseFiles("template/default.html")
	if err != nil {
		panic(err)
	}

	type FrontData struct {
		Title string
		Date  string
		Data  []internal.Data
	}

	frontData := FrontData{
		Title: "phonic-altar",
		Date:  time.Now().Local().Format("Jan-02-06"),
		Data:  td,
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, frontData)
	if err != nil {
		panic(err)
	}

	internal.WriteFile(tpl.String(), "static/index.html")
}
