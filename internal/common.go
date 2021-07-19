package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fastrodev/fastrex"
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

func ReadPost() []Post {
	data := []Post{}
	ctx := context.Background()
	db := createDatabase(ctx)

	for _, v := range db.getPost(ctx) {
		var p map[string]interface{} = v.(map[string]interface{})
		post := Post{}
		post.ID = p["id"].(string)
		post.Title = p["title"].(string)
		post.Topic = p["topic"].(string)
		post.Type = p["type"].(string)
		post.Created = p["created"].(time.Time)
		post.Content = p["content"].(string)
		data = append(data, post)
	}

	return data
}

func createResponsePage(title string, msg string, url string, res fastrex.Response) {
	resp := struct {
		Date     string
		Response string
		Title    string
		URL      string
	}{time.Now().Format("2 January 2006"), msg, title, url}
	res.Render("response", resp)
}
