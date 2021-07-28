package internal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func ReadPostByTopic(topic string) []Post {
	data := []Post{}
	ctx := context.Background()
	db := createDatabase(ctx)

	for _, v := range db.getPostByTopic(ctx, topic) {
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

func ReadPostByUsername(username string) []Post {
	data := []Post{}
	ctx := context.Background()
	db := createDatabase(ctx)

	for _, v := range db.getPostByUsername(ctx, username) {
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

func createResponsePage(res fastrex.Response, title string, msg string, url string) {
	u := strings.ToLower(url)
	resp := struct {
		Date     string
		Response string
		Title    string
		URL      string
		Domain   string
	}{time.Now().Format("2 January 2006"), msg, title, u, DOMAIN}
	res.Render("response", resp)
}

func groupByTopic(d []Post) []FlatPost {
	posts := map[string][]Post{}
	for _, v := range d {
		filtered := Filter(d, func(p Post) bool {
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

func Filter(vs []Post, f func(Post) bool) []Post {
	filtered := []Post{}
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (p *pageService) getUserFromSession(req fastrex.Request, res fastrex.Response) (*User, error) {
	c, _ := req.Cookie("__session")
	sessionByte, err := base64.StdEncoding.DecodeString(c.GetValue())
	if err != nil {
		return nil, err
	}
	if len(sessionByte) == 0 {
		return nil, errors.New("getUserFromSession:sessionByte empty")
	}
	userAgent := req.UserAgent()
	sessionID := string(sessionByte)
	userID, err := p.db.getUserIDWithSession(req.Context(), string(sessionID), userAgent)
	if err != nil {
		return nil, err
	}

	user, err := p.db.getUserDetailByID(req.Context(), userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
