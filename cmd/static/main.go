package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"time"

	"github.com/fastrodev/serverless/internal"
)

func main() {
	td := createData()
	t, err := template.ParseFiles("template/default.html")
	if err != nil {
		panic(err)
	}

	type FrontData struct {
		Email string
		Title string
		Date  string
		Data  []internal.Data
	}

	frontData := FrontData{
		Email: "oke@gmail.com",
		Title: "Iklan Baris",
		Date:  time.Now().Local().Format("2 January 2006"),
		Data:  td,
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, frontData)
	if err != nil {
		panic(err)
	}

	internal.WriteFile(tpl.String(), "public/index.html")
}

func Filter(vs []internal.Post, f func(internal.Post) bool) []internal.Post {
	filtered := []internal.Post{}
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func groupByTopic(d []internal.Post) []internal.Data {
	posts := map[string][]internal.Post{}
	for _, v := range d {
		filtered := Filter(d, func(p internal.Post) bool {
			return p.Topic == v.Topic
		})
		posts[v.Topic] = filtered
	}

	items := []internal.Data{}
	for key, element := range posts {
		data := internal.Data{
			Topic: key,
			Posts: element,
		}
		items = append(items, data)
	}
	return items
}

func createData() []internal.Data {
	body := createJsonPost()
	data := []internal.Data{}
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		log.Fatal("ReadJson" + errUnmarshal.Error())
	}
	return data
}

func createJsonPost() []byte {
	d := internal.ReadPost()
	output, errMarshal := json.Marshal(groupByTopic(d))
	if errMarshal != nil {
		panic(errMarshal)
	}

	return output
}
