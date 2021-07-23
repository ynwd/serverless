package internal

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/fastrodev/fastrex"
)

func (p *pageService) topicPage(req fastrex.Request, res fastrex.Response) {
	params := req.Params("topic")
	user, _ := p.getUserFromSession(req, res)
	email := ""
	if user != nil {
		email = user.Email
	}

	topic := strings.Title(params[0])

	td := createData(topic)
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date := now.In(loc)
	desc := "Hasil pencarian berdasarkan topic: " + topic
	data := struct {
		Email       string
		Title       string
		Data        []FlatPost
		Description string
		Date        string
		Domain      string
	}{email, topic, td, desc, date.Format("2 January 2006"), DOMAIN}
	res.Render("result", data)
}

type FlatPost struct {
	Header string
	Post
}

func createData(topic string) []FlatPost {
	t := strings.ToLower(topic)
	body := createJsonPost(t)
	data := []FlatPost{}
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		log.Fatal("ReadJson" + errUnmarshal.Error())
	}
	return data
}

func createJsonPost(topic string) []byte {
	d := ReadPostByTopic(topic)
	output, errMarshal := json.Marshal(groupByTopic(d))
	if errMarshal != nil {
		panic(errMarshal)
	}

	return output
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
