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
	t, err := template.ParseFiles("template/default.gohtml")
	if err != nil {
		panic(err)
	}

	type FrontData struct {
		Email         string
		Title         string
		Description   string
		Date          string
		PublishedDate string
		Data          []FlatPost
	}

	frontData := FrontData{
		Email:         "oke@gmail.com",
		Title:         "Iklan Baris",
		Description:   "Aplikasi web untuk membuat iklan baris online secara gratis. Simple dan nyaman dibaca.",
		Date:          time.Now().Local().Format("2 January 2006"),
		PublishedDate: time.Now().Local().Format("2006-01-0215:04:05"),
		Data:          td,
	}

	var bfr bytes.Buffer
	err = t.Execute(&bfr, frontData)
	if err != nil {
		panic(err)
	}

	internal.WriteFile(bfr.String(), "public/index.html")
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

type FlatPost struct {
	Header string
	internal.Post
}

func groupByTopic(d []internal.Post) []FlatPost {
	posts := map[string][]internal.Post{}
	for _, v := range d {
		filtered := Filter(d, func(p internal.Post) bool {
			return p.Topic == v.Topic
		})
		posts[v.Topic] = filtered
	}

	fp := []FlatPost{}
	for topic, postMap := range posts {
		for idx, element := range postMap {
			header := ""
			post := element
			if idx == 0 {
				header = topic
			}

			if idx == 7 {
				break
			}
			// cut larger content
			if len(post.Content) > 95 {
				post.Content = post.Content[0:95]
			}
			// cut larger title
			if len(post.Title) > 29 {
				post.Title = post.Title[0:29]
			}
			data := FlatPost{
				Header: header,
				Post:   post,
			}
			fp = append(fp, data)
		}
	}

	return fp
}

func createData() []FlatPost {
	body := createJsonPost()
	data := []FlatPost{}
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
