package main

import (
	"encoding/json"

	"github.com/fastrodev/serverless/internal"
)

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

func createFile() {
	d := internal.ReadData("internal/data/post.json")
	output, errMarshal := json.Marshal(groupByTopic(d))
	if errMarshal != nil {
		panic(errMarshal)
	}
	internal.WriteFile(string(output), "internal/data/index.json")
}

func main() {
	createFile()
}
