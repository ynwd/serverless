package internal

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(data.Title)
	err := res.Render("result", data)
	if err != nil {
		log.Println(err)
	}
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
		log.Printf("createData %v:", errUnmarshal.Error())
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
